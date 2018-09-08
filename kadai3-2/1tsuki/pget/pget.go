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

// Downloader provides split download feature via http Range header
type Downloader struct {
	writer io.Writer
}


// NewDownloader returns new instance of Downloader
func NewDownloader(writer io.Writer) (*Downloader) {
	return &Downloader{
		writer: writer,
	}
}

// Download starts downloading a file
func (d *Downloader) Download(url *url.URL, parallel int, timeout time.Duration) (error) {
	// get target file size
	d.printf("Acquiring file size of %s ...", url.String())
	byteLength, err := contentSize(url)
	if err != nil {
		return err
	}
	d.printf("done: %d\n", byteLength)

	// start parallel download with Goroutines
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)
	d.printf("downloading in %d parallel with timeout %s\n", parallel, timeout)
	for i := 0; i < parallel; i++ {
		filePath := partialFilePath(url, parallel, i)
		start, end := calcRange(byteLength, parallel, i, filePath)
		w := newWorker(url, start, end, filePath, d.writer)
		eg.Go(func() error {
			return w.run(ctx)
		})
	}

	// wait for all Goroutines
	if err := eg.Wait(); err != nil {
		return err
	}
	d.printf("all goroutines has finished\n")

	// join partials into one file
	d.printf("joining partials...")
	if err := joinPartials(url, parallel); err != nil {
		return err
	}
	d.printf("done\n")

	// remove all partials
	d.printf("removing partials...")
	for _, path := range partialFilePaths(url, parallel) {
		err := os.Remove(path)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to remove partial file: %s", path))
		}
	}
	d.printf("done\n")
	return nil
}

func (d *Downloader) printf(format string, a ... interface{}) {
	fmt.Fprintf(d.writer, format, a...)
}


func calcRange(byteLength, parallel, i int, filePath string) (start, end int) {
	lenSub := byteLength / parallel
	diff := byteLength % parallel

	start = lenSub * i
	end = lenSub * (i + 1)
	if i == parallel-1 {
		start += diff
	}

	// resume download if partial file already exists
	if stat, err := os.Stat(filePath); err == nil {
		start += int(stat.Size())
	}

	return start, end
}

type worker struct {
	url          *url.URL
	start, end   int
	filePath 	 string
	writer       io.Writer
}

func newWorker(url *url.URL, start, end int, filePath string, writer io.Writer) *worker {
	return &worker{
		url:      url,
		start:    start,
		end:      end,
		filePath: filePath,
		writer:   writer,
	}
}

func (r *worker) run(ctx context.Context) error {
	r.printf("  started downloading partial: %s\n", r.filePath)
	client := &http.Client{}
	req, err := http.NewRequest("GET", r.url.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", r.start, r.end-1))
	req.WithContext(ctx)

	errCh := make(chan error, 1)
	go func() {
		// download
		resp, err := client.Do(req)
		if err != nil {
			errCh <- errors.Wrap(err, fmt.Sprintf("failed to download partial file: %s\n", r.url))
		}
		defer resp.Body.Close()

		// write into
		out, err := os.OpenFile(r.filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			errCh <- errors.Wrap(err, fmt.Sprintf("failed to write partial file: %s", r.filePath))
		}
		defer out.Close()

		io.Copy(out, resp.Body)

		r.printf("  finished downloading partial: %s\n", r.filePath)
		errCh <- nil
	}()

	select {
	case err := <-errCh: // wait for complete
		if err != nil {
			return err
		}
	case <-ctx.Done(): // abort on ctx.Done()
		return errors.Wrap(ctx.Err(), "context has closed during download")
	}

	return nil
}

func (r *worker) printf(format string, a... interface{}) {
	fmt.Fprintf(r.writer, format, a...)
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
