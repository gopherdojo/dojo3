package download

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/pkg/errors"
)

const green = "\x1b[32m%s\x1b[0m\n"

// Download interface
type Download interface {
	Download() error
}

// Downloader struct
type Downloader struct {
	outStream io.Writer
	url       *url.URL
	dirName   string
	parallel  int64
	timeout   time.Duration
	resCh     <-chan map[int]string
}

// Create new Downloader struct
func NewDownloader(out io.Writer, url *url.URL, dirName string,
	parallel int64, timeout time.Duration) Download {

	return &Downloader{
		outStream: out,
		url:       url,
		dirName:   dirName,
		parallel:  parallel,
		timeout:   timeout}
}

// Parallel download
func (d *Downloader) Download() error {
	len, err := d.getContentLength()
	if err != nil {
		return err
	}

	fmt.Fprintf(d.outStream, "> Content-Length: %v\n", len)

	byteRanges := d.calculationRange(len)
	timer := time.NewTimer(d.timeout)

	if err := d.rangeDownload(byteRanges); err != nil {
		return err
	}

	for {
		select {
		case resMap := <-d.resCh:
			// TODO
			println(resMap)
		case <-timer.C:
			return errors.New(fmt.Sprintf("timeout: %s", d.timeout))
		}
	}

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

	byteRanges := make([]string, d.parallel)

	start := int64(0)
	for i := 0; i < int(d.parallel); i++ {
		end := start + div
		if i == int(d.parallel-1) {
			end = length
		}
		byteRanges[i] = fmt.Sprintf("%v-%v", start, end)
		start = end + 1
	}

	// example
	// Content-Length: 74121, parallel: 5
	// [0-14824 14825-29649 29650-44474 44475-59299 59300-74120]
	return byteRanges
}

func (d *Downloader) sendHTTPRequest(byteRange string) (*http.Response, error) {
	client := http.Client{Timeout: d.timeout}

	req, err := http.NewRequest("GET", d.url.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err,
			fmt.Sprintf("failed create HTTP request. url: %s, byte_range: %s", d.url, byteRange))
	}
	req.Header.Set("Range", "bytes="+byteRange)

	return client.Do(req)
}

func (d *Downloader) rangeDownload(byteRanges []string) error {
	resCh := make(chan map[int]string)
	d.resCh = resCh

	var eg errgroup.Group
	resMap := make(map[int]string)

	for i, byteRange := range byteRanges {
		i := i
		byteRange := byteRange
		eg.Go(func() error {
			res, err := d.sendHTTPRequest(byteRange)
			if err != nil {
				return errors.Wrap(err,
					fmt.Sprintf("failed download. url: %s, byte_range: %s", d.url, byteRange))
			}

			fmt.Fprintf(d.outStream, "> %s "+green, byteRange, "download success")

			fileName, err := d.writeTempFile(res)
			if err != nil {
				return err
			}

			resMap[i] = fileName
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	resCh <- resMap
	return nil
}

func (d *Downloader) writeTempFile(res *http.Response) (string, error) {
	temp, err := ioutil.TempFile(d.dirName, "temp")
	if err != nil {
		return "", errors.Wrap(err, "failed create temp file.")
	}
	defer temp.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed read response body.")
	}

	if _, err := temp.Write(body); err != nil {
		return "", errors.Wrap(err,
			fmt.Sprintf("failed write response body. file: %s", temp.Name()))
	}

	return temp.Name(), nil
}
