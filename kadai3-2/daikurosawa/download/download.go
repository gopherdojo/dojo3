package download

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// Download interface
type Download interface {
	Download() error
}

// Downloader struct
type Downloader struct {
	outStream io.Writer
	url       *url.URL
	parallel  int64
	timeout   time.Duration
}

// Create new Downloader struct
func NewDownloader(out io.Writer, url *url.URL, parallel int64, timeout time.Duration) Download {
	return &Downloader{outStream: out, url: url, parallel: parallel, timeout: timeout}
}

// Parallel download
func (d *Downloader) Download() error {
	len, err := d.getContentLength()
	if err != nil {
		return err
	}

	fmt.Fprintf(d.outStream, "> Content-Length: %v\n", len)

	byteRange := d.calculationRange(len)

	// TODO
	fmt.Println(byteRange)
	return nil
}

func (d *Downloader) getContentLength() (int64, error) {
	res, err := http.Head(d.url.String())
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("failed head request. url: %s", d.url))
	}

	len := res.ContentLength
	if len < 1 {
		errors.New(fmt.Sprintf("Content-Length is zero. url: %s, length: %v", d.url, len))
	}

	header := res.Header

	// confirm range download support
	if accept, ok := header["Accept-Ranges"]; ok {
		if accept[0] == "none" {
			errors.New(fmt.Sprintf("Accept-Ranges header is none, unsupported range download. url: %s", d.url))
		}
	} else {
		errors.New(fmt.Sprintf("no Accept-Ranges header. url: %s", d.url))
	}
	return len, nil
}

func (d *Downloader) calculationRange(length int64) []string {
	length--
	div := length / d.parallel

	byteRange := make([]string, d.parallel)

	start := int64(0)
	for i := 0; i < int(d.parallel); i++ {
		end := start + div
		if i == int(d.parallel-1) {
			end = length
		}
		byteRange[i] = fmt.Sprintf("%v-%v", start, end)
		start = end + 1
	}

	// example
	// Content-Length: 74121, parallel: 5
	// [0-14824 14825-29649 29650-44474 44475-59299 59300-74120]
	return byteRange
}
