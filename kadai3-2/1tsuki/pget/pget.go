package pget

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Downloader struct {
	writer io.Writer
}

func NewDownloader(writer io.Writer) (*Downloader) {
	return &Downloader{
		writer: writer,
	}
}

func (d *Downloader) Download(url *url.URL, parallel int, timeout time.Duration) (error) {
	d.printf("Download %s in %d parallel with timeout %s \n",
		url.String(),
		parallel,
		timeout)

	// get target filesize
	byteLength, err := contentSize(url)
	if err != nil {
		return err
	}
	d.printf("target file size: %d", byteLength)

	// start parallel download with Goroutines
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, timeout)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < parallel; i++ {
		r := NewRangeDownload(url, byteLength, parallel, i)
		eg.Go(func() error {
			return r.Run(ctx)
		})
	}

	// wait for all Goroutines
	if err := eg.Wait(); err != nil {
		return err
	}

	// concatenate split files
	if err := d.concatenate(url, parallel); err != nil {
		return err
	}

	// remove temp files on success
	d.printf("deleting all temp files\n")
	for j := 0; j < parallel; j++ {
		err := os.Remove(partialFileName(url, j))
		if err != nil {
			return err
		}
	}
	return nil
}

type RangeDownload struct {
	url  *url.URL
	min  int
	max  int
	path string
}

func NewRangeDownload(url *url.URL, byteLength, parallel, i int) *RangeDownload {
	lenSub := byteLength / parallel
	diff := byteLength % parallel

	min := lenSub * i
	max := lenSub * (i + 1)
	if i == parallel-1 {
		max += diff
	}

	return &RangeDownload{
		url:  url,
		min:  min,
		max:  max,
		path: partialFileName(url, i),
	}
}

func (r *RangeDownload) Run(ctx context.Context) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", r.url.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", r.min, r.max-1))
	req.WithContext(ctx)

	errCh := make(chan error, 1)
	go func() {
		resp, err := client.Do(req)
		if err != nil {
			errCh <- err
		}
		defer resp.Body.Close()

		// copy into file
		out, err := os.OpenFile(r.path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			errCh <- err
		}
		defer out.Close()

		io.Copy(out, resp.Body)
		errCh <- nil
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}

func (d *Downloader) printf(format string, a ... interface{}) {
	fmt.Fprintf(d.writer, format, a...)
}

func (d *Downloader) concatenate(url *url.URL, parallel int) error {
	dst := filepath.Base(url.EscapedPath())
	d.printf("concatenating files to %s\n", dst)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	for i := 0; i < parallel; i++ {
		if err := joinPartialFile(partialFileName(url, i), out); err != nil {
			return err
		}
	}
	return nil
}

func partialFileName(url *url.URL, i int) string {
	return filepath.Base(url.EscapedPath()) + strconv.Itoa(i)
}

func joinPartialFile(src string, out *os.File) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	return nil
}

func contentSize(url *url.URL) (int, error) {
	res, err := http.Head(url.String())
	if err != nil {
		return 0, err
	}
	maps := res.Header
	length, err := strconv.Atoi(maps["Content-Length"][0])
	if err != nil {
		return 0, err
	}
	return length, nil
}
