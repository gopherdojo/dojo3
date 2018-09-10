package download

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
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
	tempDir   string
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

	resMap, err := d.rangeDownload(byteRanges)
	if err != nil {
		return err
	}

	if err := d.merge(resMap); err != nil {
		return err
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

func (d *Downloader) rangeDownload(byteRanges []string) (map[int]string, error) {
	tempDir, err := ioutil.TempDir(d.dirName, "temp")
	if err != nil {
		return nil, errors.Wrap(err, "create temp directory.")
	}
	d.tempDir = tempDir

	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)
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

			fileName, err := d.writeTempFile(res, tempDir)
			if err != nil {
				return err
			}

			resMap[i] = fileName
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return resMap, nil
}

func (d *Downloader) writeTempFile(res *http.Response, tempDir string) (string, error) {
	temp, err := ioutil.TempFile(tempDir, "temp")
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

func (d *Downloader) merge(resMap map[int]string) error {
	_, fileName := path.Split(d.url.String())
	file, err := os.OpenFile(d.tempDir+fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrap(err, "failed create file.")
	}
	defer func() {
		file.Close()
		os.RemoveAll(d.tempDir)
	}()

	for i := 0; i < len(resMap); i++ {
		if res, ok := resMap[i]; ok {
			if err := d.mergeFile(file, res); err != nil {
				return err
			}
		} else {
			return errors.New("not found range download response.")
		}
	}

	return nil
}

func (*Downloader) mergeFile(file *os.File, tempFile string) error {
	temp, err := os.Open(tempFile)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed open temp file. file: %s", tempFile))
	}
	defer temp.Close()
	if _, err := io.Copy(file, temp); err != nil {
		return errors.Wrap(err,
			fmt.Sprintf("failed merge file. file: %s, temp_file: %s", file.Name(), tempFile))
	}

	return nil
}
