package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
)

func main() {
	var (
		dir = flag.String("o", ".", "output directory")
		num = flag.Int("n", 3, "num of worker")
	)
	flag.Parse()

	url := flag.Arg(0)
	filename := path.Join(*dir, path.Base(url))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %+v\n", err)
	}
	defer file.Close()

	err = Exec(url, file, *num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %+v\n", err)
		os.Remove(filename)
	}
}

func Exec(url string, w io.Writer, num int) error {
	length, ok := acceptsRangeRequest(url)
	if num > 1 && ok {
		b := length/num + 1
		eg, ctx := errgroup.WithContext(context.Background())

		// Range access
		tmpFiles := make([]io.Reader, num)
		defer func() {
			for _, tmpFile := range tmpFiles {
				if f, ok := tmpFile.(io.ReadCloser); ok {
					f.Close()
				}
			}
		}()

		for i := 0; i < num; i++ {
			i := i
			eg.Go(func() error {
				tmpFile, err := ioutil.TempFile("", path.Base(url))
				if err != nil {
					return errors.WithStack(err)
				}
				tmpFiles[i] = tmpFile
				header := map[string]string{"Range": fmt.Sprintf("bytes=%v-%v", i*b, (i+1)*b-1)}
				return download(ctx, url, header, tmpFile)
			})
		}
		if err := eg.Wait(); err != nil {
			return errors.WithStack(err)
		}

		// Concatenate partial files
		return concat(w, tmpFiles)
	} else {
		ctx, _ := context.WithCancel(context.Background())
		return download(ctx, url, nil, w)
	}
}

func acceptsRangeRequest(url string) (int, bool) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, false
	}
	defer resp.Body.Close()
	unit := resp.Header.Get("Accept-Ranges")
	length, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return 0, false
	}
	return length, unit == "bytes"
}

func download(ctx context.Context, url string, h map[string]string, w io.Writer) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	for k, v := range h {
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func concat(dst io.Writer, srcs []io.Reader) error {
	for _, src := range srcs {
		err := func() error {
			_, err := io.Copy(dst, src)
			if err != nil {
				return errors.WithStack(err)
			}
			return nil
		}()
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
