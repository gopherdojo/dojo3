package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

const sample = "http://i.imgur.com/z4d4kWk.jpg"
const splitNum = 4

func request(r *http.Request, client *http.Client, errCh chan<- error) (*http.Response, []byte) {
	res, err := client.Do(r)
	if err != nil {
		errCh <- err
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, res.Body)
	if err != nil {
		errCh <- err
	}

	// debug
	header := res.Header
	//	for h := range header {
	fmt.Printf("Content-Range: %v\n", header["Content-Range"])
	//	}

	/*
	err = ioutil.WriteFile("test.jpg", buf.Bytes(), 0777)
	if err != nil {
		errCh <- err
	}*/

	return res, buf.Bytes()
}

func requestWithRange(r *http.Request, client *http.Client, errCh chan<- error, start, end int) (*http.Response, []byte) {
	r.Header.Add("Range", fmt.Sprintf("bytes=%v-%v", start, end))

	return request(r, client, errCh)
}

func fetchRange(r *http.Request, client *http.Client, errCh chan<- error) (int, error) {
	res, _ := request(r, client, errCh)
	cl := res.Header.Get("Content-Length")
	l, err := strconv.Atoi(cl)
	return l, err
}

func makeRequest(method, url string, ctx context.Context) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	return req, nil
}

func Run(ctx context.Context) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	errCh := make(chan error, 1)
	done := make(chan struct{}, 1)

	req, err := makeRequest("HEAD", sample, ctx)
	if err != nil {
		return err
	}
	l, err := fetchRange(req, client, errCh)
	if err != nil {
		return err
	}
	chunk := l / splitNum
	result := make([]byte, l, l)

	var wg sync.WaitGroup

	for i := 0; i < splitNum; i++ {
		i := i
		wg.Add(1)
		go func() {
			start := chunk * i
			end := chunk * (i + 1)
			if i == splitNum-1 {
				end = l
			}
			req, err := makeRequest("GET", sample, ctx)
			if err != nil {
				errCh <- err
			}
			_, body := requestWithRange(req, client, errCh, start, end)
			copy(result[start:end], body)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			println("errCh")
			return err
		}
	case <-ctx.Done():
		println("done")
		<-errCh
		return ctx.Err()
	case <-done:
		return writeFile(result)
	}

	return err
}

func writeFile(result []byte) error {
	return ioutil.WriteFile("test.jpg", result, 0777)
}
