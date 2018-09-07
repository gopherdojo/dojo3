package pget

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
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

type Range struct {
	min, max int
}

func NewRange(length, parallel, i int) *Range {
	lenSub := length / parallel
	diff := length % parallel

	min := lenSub * i
	max := lenSub * (i + 1)
	if i == parallel-1 {
		max += diff
	}

	return &Range{
		min: min,
		max: max,
	}
}

func (d *Downloader) Download(url *url.URL, parallel int, timeout time.Duration) (error) {
	d.printf("Downloading file %s with %d parallel\n", url.String(), parallel)
	byteLength, err := contentSize(url)
	if err != nil {
		return err
	}

	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, timeout)
	defer cancel()

	errCh := make(chan error, parallel)
	for i := 0; i < parallel; i++ {
		r := NewRange(byteLength, parallel, i)
		go func(r *Range, i int) {
			errCh <- d.rangeDownload(ctx, url, r, partialFileName(url, i))
		}(r, i)
	}

	// wait for all Goroutines
	for c := 0; c < cap(errCh); c++ {
		err := <-errCh
		if err != nil {
			return err
		}
	}

	// concat files
	if err := d.concatenate(url, parallel); err != nil {
		return err
	}

	// remove temp files
	d.printf("deleting all temp files\n")
	for j := 0; j < parallel; j++ {
		err := os.Remove(partialFileName(url, j))
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Downloader) rangeDownload(ctx context.Context, url *url.URL, r *Range, path string) error {
	d.printf("Start downloading byte range %d...%d\n", r.min, r.max)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url.String(), nil)
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

		reader, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errCh <- err
		}
		body := string(reader)
		ioutil.WriteFile(path, []byte(string(body)), 0x777)
		d.printf("saved file %s\n", path)

		errCh <- nil
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		<-errCh
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
