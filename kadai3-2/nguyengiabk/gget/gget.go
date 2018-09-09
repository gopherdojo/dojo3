// Package gget provides functions to download files
package gget

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gopherdojo/dojo3/kadai3-2/nguyengiabk/opt"
	"golang.org/x/sync/errgroup"
)

// GGet defines downloader
type GGet struct {
	url        string
	procNum    int
	fileName   string
	fileLength int64
	partFiles  []string
}

const defaultFileMode = 0600
const writeBufferSize = 256

// NewGGet create new downloader
func NewGGet(params *opt.Parameter) (*GGet, error) {
	g := GGet{
		url:     params.URL,
		procNum: params.ProcNum,
	}
	if err := g.checkURLInfo(); err != nil {
		return nil, err
	}
	return &g, nil
}

func (g *GGet) checkURLInfo() error {
	g.fileName = filepath.Base(g.url)
	res, err := http.Head(g.url)
	if err != nil {
		return err
	}

	if res.ContentLength <= 0 {
		return errors.New("Content Length is invalid")
	}
	g.fileLength = res.ContentLength
	if res.Header.Get("Accept-Ranges") != "bytes" {
		g.procNum = 1
	}
	for i := 0; i < g.procNum; i++ {
		partFileName := fmt.Sprintf("%s.part.%d", g.fileName, i)
		g.partFiles = append(g.partFiles, partFileName)
	}
	return nil
}

func (g *GGet) makeRequest(rangeFrom int64, rangeTo int64) *http.Request {
	req, err := http.NewRequest("GET", g.url, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var ranges string
	ranges = fmt.Sprintf("bytes=%d-%d", rangeFrom, rangeTo)
	req.Header.Add("Range", ranges)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return req
}

func (g *GGet) processPart(ctx context.Context, partNum int) error {
	partLen := g.fileLength / int64(g.procNum)
	rangeFrom := int64(partNum) * partLen
	rangeTo := int64(partNum+1)*partLen - 1
	if partNum == g.procNum-1 {
		rangeTo = g.fileLength - 1
	}
	req := g.makeRequest(rangeFrom, rangeTo)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(g.partFiles[partNum], os.O_CREATE|os.O_WRONLY|os.O_APPEND, defaultFileMode)
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		select {
		case <-ctx.Done():
			f.Close()
			os.Remove(g.partFiles[partNum])
			return errors.New("process was cancelled")
		default:
			_, err := io.CopyN(f, resp.Body, writeBufferSize)
			if err != nil {
				if err != io.EOF {
					f.Close()
					os.Remove(g.partFiles[partNum])
					return err
				}
				return nil
			}
		}
	}
}

// Process downloads file
func (g *GGet) Process() error {
	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-signalChan
		cancel()
	}()

	for i := 0; i < g.procNum; i++ {
		partNum := i
		eg.Go(func() error {
			return g.processPart(ctx, partNum)
		})
	}
	if err := eg.Wait(); err != nil {
		cancel()
		return err
	}
	if err := g.Join(); err != nil {
		return err
	}
	return nil
}
