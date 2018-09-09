package gget

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

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

func (g *GGet) processPart(partNum int, partLen int64) error {
	rangeFrom := int64(partNum) * partLen
	rangeTo := int64(partNum+1)*partLen - 1
	if partNum == g.procNum-1 {
		rangeTo = g.fileLength - 1
	}
	req := g.makeRequest(rangeFrom, rangeTo)
	//write to file
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	partFileName := fmt.Sprintf("%s.part.%d", g.fileName, partNum)
	f, err := os.OpenFile(partFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, defaultFileMode)
	if err != nil {
		return err
	}
	defer f.Close()

	current := int64(0)
	for {
		select {
		default:
			written, err := io.CopyN(f, resp.Body, writeBufferSize)
			current += written
			if err != nil {
				if err != io.EOF {
					return err
				}
				return nil
			}
		}
	}
}

// Do downloads file
func (g *GGet) Do(ctx context.Context) error {
	eg, _ := errgroup.WithContext(ctx)
	partLen := g.fileLength / int64(g.procNum)
	for i := 0; i < g.procNum; i++ {
		partNum := i
		eg.Go(func() error {
			return g.processPart(partNum, partLen)
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
	var partFiles []string
	for i := 0; i < g.procNum; i++ {
		partFileName := fmt.Sprintf("%s.part.%d", g.fileName, i)
		partFiles = append(partFiles, partFileName)
	}
	if err := Join(partFiles, g.fileName); err != nil {
		log.Fatal(err)
	}
	for _, name := range partFiles {
		os.Remove(name)
	}
	return nil
}
