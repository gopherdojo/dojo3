package xget

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	netURL "net/url"
	"os"
	"path"
	"runtime"
)

var (
	MinChunkSize = int64(1024)
)

type Option struct {
	Procs int
}

type Client struct {
	url   string
	procs int
}

type plan struct {
	url, path string
	size      int64
	chunks    []chunk
}

type chunk struct {
	first int64
	last  int64
}

func NewClient(url string, opt Option) (*Client, error) {
	c := Client{}

	u, err := netURL.ParseRequestURI(url)
	if err != nil {
		return nil, errors.Wrap(err, "invalid url")
	}
	c.url = u.String()

	if opt.Procs > 0 {
		c.procs = opt.Procs
	} else {
		c.procs = runtime.NumCPU()
	}

	if c.procs <= 0 {
		return nil, errors.New("not positive Procs")
	}

	return &c, nil
}

func (c *Client) setMaxProcs() {
	if procs := os.Getenv("GOMAXPROCS"); procs == "" {
		runtime.GOMAXPROCS(c.procs)
	}
}

func (c *Client) Run(ctx context.Context) error {
	c.setMaxProcs()
	if err := c.download(ctx); err != nil {
		return errors.Wrap(err, "failed to download")
	}

	return nil
}

func (c *Client) download(ctx context.Context) error {
	plan, err := c.plan(ctx)

	if err != nil {
		return errors.Wrap(err, "failed to plan download")
	}

	eg, ctx := errgroup.WithContext(ctx)
	chunkPaths := make([]string, len(plan.chunks))

	for i, chunk := range plan.chunks {
		_chunk := chunk
		path := fmt.Sprintf("%s.chunk_%d", plan.path, i)
		chunkPaths[i] = path

		eg.Go(func() error {
			return chunkDownload(ctx, _chunk, path, plan.url)
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	// merge
	merge(chunkPaths, plan.path)
	return nil
}

func (c *Client) plan(ctx context.Context) (*plan, error) {
	res, err := ctxhttp.Head(ctx, http.DefaultClient, c.url)

	if err != nil {
		return nil, errors.Wrap(err, "failed to head request")
	}

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		return nil, errors.New("not supported range access")
	}

	if res.ContentLength <= 0 {
		return nil, errors.New("invalid content length")
	}

	current := int64(0)
	remainSize := res.ContentLength
	chunkSize := res.ContentLength/int64(c.procs) + 1

	if chunkSize < MinChunkSize {
		chunkSize = MinChunkSize
	}

	chunks := make([]chunk, c.procs)
	var lastIndex int
	for i := range chunks {
		if remainSize <= 0 {
			break
		}

		first := current
		var last int64
		if remainSize > chunkSize {
			last = current + chunkSize
		} else {
			last = current + remainSize
			// last chunk index
			lastIndex = i
		}
		remainSize -= last - first
		chunks[i].first = first
		chunks[i].last = last
		current = last + 1
	}

	if remainSize != 0 {
		return nil, errors.New("invalid chunk")
	}

	_, path := path.Split(res.Request.URL.Path)
	p := plan{
		res.Request.URL.String(), // for redirect
		path,
		res.ContentLength,
		chunks[:lastIndex+1],
	}

	return &p, nil
}

func chunkDownload(ctx context.Context, chunk chunk, path string, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.Wrap(err, "failed to make request")
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", chunk.first, chunk.last))

	res, err := ctxhttp.Do(ctx, http.DefaultClient, req)
	if err != nil {
		return errors.Wrap(err, "failed to request chunkDownload")
	}
	defer res.Body.Close()

	output, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to create %s", path))
	}
	defer output.Close()

	_, err = io.Copy(output, res.Body)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to copy %s", path))
	}

	return nil
}

func merge(fromPaths []string, toPath string) error {
	toFile, err := os.Create(toPath)
	if err != nil {
		return errors.Wrap(err, "failed to create output file")
	}
	defer toFile.Close()

	for _, path := range fromPaths {
		fromFile, err := os.Open(path)
		if err != nil {
			return errors.Wrap(err, "failed to open chunk file")
		}

		io.Copy(toFile, fromFile)
		fromFile.Close()

		if err := os.Remove(path); err != nil {
			return errors.Wrap(err, "failed to remove chunk file")
		}
	}

	return nil
}
