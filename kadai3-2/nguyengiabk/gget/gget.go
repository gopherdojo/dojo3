// Package gget provides functions to download files
package gget

import (
	"context"
	"errors"
	"net/http"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai3-2/nguyengiabk/opt"
)

// GGet defines downloader
type GGet struct {
	url        string
	procNum    int
	fileName   string
	fileLength int64
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

// Process start downloading file
func (g *GGet) Process() error {
	ctx := context.Background()
	return g.Do(ctx)
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
	return nil
}
