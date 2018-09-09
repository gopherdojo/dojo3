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
	parallel  int
	timeout   time.Duration
}

// Create new Downloader struct
func NewDownloader(out io.Writer, url *url.URL, parallel int, timeout time.Duration) Download {
	return &Downloader{outStream: out, url: url, parallel: parallel, timeout: timeout}
}

// Parallel download
func (d *Downloader) Download() error {
	len, err := d.getContentLength()
	if err != nil {
		return err
	}

	fmt.Fprintf(d.outStream, "> Content-Length: %v\n", len)

	return nil
}

func (d *Downloader) getContentLength() (int64, error) {
	res, err := http.Head(d.url.String())
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("failed head request. url: %s", d.url))
	}

	len := res.ContentLength
	if len < 1 {
		errors.New(fmt.Sprintf("Content-Length is zero. length: %v", len))
	}
	return len, nil
}
