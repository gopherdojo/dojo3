package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {

	url := os.Args[1]
	fmt.Println(url)

	//TODO check args

	//TODO check RangeAccess available

	d, err := NewDownloader(url)
	if err != nil {
		fmt.Printf("\ndownload initialize error. %v", err)
		os.Exit(1)
	}

	err = d.Download()
	if err != nil {
		fmt.Printf("\ndownload error. %v", err)
		os.Exit(1)
	}
	err = d.Merge()
	if err != nil {
		fmt.Printf("\ndownload merge error. %v", err)
		os.Exit(1)
	}
}

type Downloader struct {
	FileName   string
	MaxProcess uint
	Workers    []*worker
}

type worker struct {
	processId            uint
	bytesToStartReading  uint
	bytesToFinishReading uint
	resourceUrl          string
	partFilePath         string
}

func NewDownloader(url string) (*Downloader, error) {
	d := new(Downloader)
	d.FileName = getFileName(url)
	d.MaxProcess = uint(runtime.NumCPU())

	res, err := http.Head(url)
	if err != nil {
		return nil, err
	}

	size := uint(res.ContentLength)

	split := size / d.MaxProcess
	for i := uint(0); i < d.MaxProcess; i++ {
		w, err := NewWorker(d, size, i, split, url)
		if err != nil {
			return nil, errors.Wrap(err, "initialize worker error")
		}
		d.Workers = append(d.Workers, w)
	}
	fmt.Fprintf(os.Stdout, "Download start from %s\n", url)
	return d, nil
}

func NewWorker(d *Downloader, size uint, i uint, split uint, url string) (*worker, error) {
	bytesToStartReading := split * i
	bytesToFinishReading := bytesToStartReading + split - 1
	partFilePath := fmt.Sprintf("%s.%d", d.FileName, i)

	if i == d.MaxProcess-1 {
		bytesToFinishReading = size
	}
	w := &worker{
		processId:            i,
		bytesToStartReading:  bytesToStartReading,
		bytesToFinishReading: bytesToFinishReading,
		resourceUrl:          url,
		partFilePath:         partFilePath,
	}
	return w, nil
}

func (d *Downloader) Download() error {
	eg := errgroup.Group{}
	for _, worker := range d.Workers {
		w := worker
		eg.Go(func() error {
			return w.Request()
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func (d *Downloader) Merge() error {
	outputFilePath := fmt.Sprintf("%s", d.FileName)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return errors.Wrap(err, "failed to create merge file")
	}
	defer outputFile.Close()
	for i := uint(0); i < d.MaxProcess; i++ {
		partFilePath := fmt.Sprintf("%s.%d", outputFilePath, i)
		partFile, err := os.Open(partFilePath)
		if err != nil {
			return errors.Wrap(err, "failed to open part file")
		}
		io.Copy(outputFile, partFile)
		partFile.Close()
		if err := os.Remove(partFilePath); err != nil {
			return errors.Wrap(err, "failed to remove a file")
		}
	}
	return nil
}

func (w *worker) Request() error {
	res, err := w.MakeResponse()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to split get requests: %d", w.processId))
	}
	defer res.Body.Close()
	output, err := os.Create(w.partFilePath)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to create file %s", w.partFilePath))
	}
	defer output.Close()
	io.Copy(output, res.Body)
	return nil
}

func (w *worker) MakeResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", w.resourceUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to split NewRequest for get: %d", w.processId))
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", w.bytesToStartReading, w.bytesToFinishReading))
	return http.DefaultClient.Do(req)
}

func getFileName(resourceUrl string) string {
	token := strings.Split(resourceUrl, "/")
	filename := token[len(token)-1]
	return filename
}
