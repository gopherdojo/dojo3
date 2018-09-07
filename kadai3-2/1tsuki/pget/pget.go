package pget

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
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
	d.printf("Downloading %s in %d parallel with timeout %s\n", url.String(), parallel, timeout)

	// get target file size
	byteLength, err := contentSize(url)
	if err != nil {
		return err
	}
	d.printf("target file size: %d\n", byteLength)

	// start parallel download with Goroutines
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < parallel; i++ {
		r := NewRangeDownload(url, byteLength, parallel, i)
		eg.Go(func() error {
			return r.Download(ctx)
		})
	}

	// wait for all Goroutines
	if err := eg.Wait(); err != nil {
		return err
	}
	d.printf("download complete\n")

	// join partials into one file
	if err := joinPartials(url, parallel); err != nil {
		return err
	}

	// remove all partials
	for _, path := range partialFilePaths(url, parallel) {
		err := os.Remove(path)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to remove partial file: %s", path))
		}
	}
	return nil
}

func (d *Downloader) printf(format string, a ... interface{}) {
	fmt.Fprintf(d.writer, format, a...)
}

type RangeDownload struct {
	url              *url.URL
	min, max         int
	parallel, worker int
}

func NewRangeDownload(url *url.URL, byteLength int, parallel, i int) *RangeDownload {
	var (
		min, max int
	)

	lenSub := byteLength / parallel
	diff := byteLength % parallel

	filePath := partialFilePath(url, parallel, i)
	stat, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		min = lenSub * i
	} else {
		// FIXME stat.Size() returns int64
		// TODO resume detection
		min = int(stat.Size())
	}

	min = lenSub * i
	max = lenSub * (i + 1)
	if i == parallel-1 {
		max += diff
	}

	return &RangeDownload{
		url:      url,
		min:      min,
		max:      max,
		parallel: parallel,
		worker:   i,
	}
}

func (r *RangeDownload) Download(ctx context.Context) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", r.url.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", r.min, r.max-1))
	req.WithContext(ctx)

	errCh := make(chan error, 1)
	go func() {
		// download
		resp, err := client.Do(req)
		if err != nil {
			errCh <- errors.Wrap(err, fmt.Sprintf("failed to download partial file: %s", r.url))
		}
		defer resp.Body.Close()

		// write into
		out, err := os.OpenFile(r.savePath(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			errCh <- errors.Wrap(err, fmt.Sprintf("failed to write partial file: %s", r.savePath()))
		}
		defer out.Close()

		io.Copy(out, resp.Body)
		errCh <- nil
	}()

	// wait for complete or abort on ctx.Done()
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return errors.Wrap(ctx.Err(), "context has closed during download")
	}

	return nil
}

func (r *RangeDownload) savePath() string {
	return partialFilePath(r.url, r.parallel, r.worker)
}

func joinPartials(url *url.URL, parallel int) error {
	dst := filepath.Base(url.EscapedPath())

	out, err := os.Create(dst)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to create out file: %s", dst))
	}
	defer out.Close()

	for i := 0; i < parallel; i++ {
		if err := joinPartial(partialFilePath(url, parallel, i), out); err != nil {
			return err
		}
	}
	return nil
}

func joinPartial(src string, out *os.File) error {
	in, err := os.Open(src)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to open partial file: %s", src))
	}
	defer in.Close()

	if _, err := io.Copy(out, in); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to join file: %s", src))
	}

	return nil
}

func partialFilePaths(url *url.URL, parallel int) []string {
	names := make([]string, parallel)
	for i := 0; i < parallel; i++ {
		names[i] = partialFilePath(url, parallel, i)
	}
	return names
}

func partialFilePath(url *url.URL, parallel, i int) string {
	return fmt.Sprintf("%s_%dof%d", filepath.Base(url.EscapedPath()), i+1, parallel)
}

func contentSize(url *url.URL) (int, error) {
	res, err := http.Head(url.String())
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("failed to request head: %s", url))
	}
	maps := res.Header
	length, err := strconv.Atoi(maps["Content-Length"][0])
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("failed to parse Content-Length: %v", maps))
	}
	return length, nil
}
