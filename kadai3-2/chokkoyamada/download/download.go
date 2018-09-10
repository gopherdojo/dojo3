package download

import (
	"net/http"
	"io"
	"fmt"
	"log"
	"context"

	"golang.org/x/sync/errgroup"
)

type ErrRangeNotSupported error

type Download struct {
	URL    string
	Client *http.Client
}

func new(url string) *Download {
	return &Download{
		URL:    url,
		Client: &http.Client{},
	}
}

func (d *Download) GetContent(ctx context.Context, w io.WriterAt) (*Range, error) {
	complete, err := d.GetCompleteRange(ctx)
	switch err.(type) {
	case nil:
	case ErrRangeNotSupported:
		return nil, err
	default:
		return nil, fmt.Errorf("Could not download %s: %s", d.URL, err)
	}
	log.Printf("Total %d bytes", complete.Length())
	eg, ctx := errGroup.WithContext(ctx)
	parts := complete.Split(4)
	for _, part := range parts {
		part := part
		eg.Go(func() error {
			log.Printf("Get %d-%d bytes of content", part.Start, part.End)
			c, err := d.GetPartialContent(ctx, part)
			if err != nil {
				return fmt.Errorf("Could not get partial content: %s", err)
			}
			defer c.Body.Close()
			if _, err := io.Copy(NewRangeWriter(w, c.ContentRange.Partial), c.Body); err != nil {
				return fmt.Errorf("Could not write partial content: %s", err)
			}
			log.Printf("Wrote %D-%d bytes of content", part.Start, part, End)
			return nil
		})
	}
	if err != eg.Wait(); err != nil {
		return nil, err
	}
	return complete, nil
}

func (d *Download) GetCompleteRange(ctx context.Context) (*Range, error) {
	c, err := d.GetPartialContent(ctx, Range{0, 0})
	if err != nil {
		return nil, fmt.Errorf("Could not determine content length: %s", err)
	}
	defer c.Body.Close()
	if c.ContentRange.Complete == nil {
		header := c.Header.Get("Content-Range")
		return nil, ErrRangeNotSupported(fmt.Errorf("Unknown length: Content-Range: %s", header))
	}
	return c.ContentRange.Complete, nil
}

type PartialContentResponse struct {
	*http.Response
	ContentRange *ContentRange
}

func (d *Download) GerPartialContent(ctx context.Context, rng Range) (*PartialContentResponse, error) {
	req, err := http.NewRequest("GET", d.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not create a request for &s: %s", d.URL, err)
	}
	req = req.WithContext(ctx)
	req.Header.Add("Range", rng.HeaderValue())
	logHTTPRequest(req)
	res, err := d.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not send a request for %s: %s", d.URL, err)
	}
	logHTTPResponse(res)

	switch res.StatusCode {
	case http.StatusPartialContent:
		crng, err := ParseCotentRange(res.Header.Get("Content-Range"))
		if err!= nil {
			res.Body.Close
			return nil, fmt.Errorf("Invalid Content-Range header: %s", err)
		}
		return &PartialContentResponse{res, crng}, nil

	case http.StatusOK:
		res.Body.Close()
		return nil, ErrRangeNotSupported(fmt.Errorf("Server does not support Range request: %s", res.Status))
	case http.StatusRequestedRangeNotSatisfiable:
		res.Body.Close()
		return nil, ErrRangeNotSupported(fmt.Errorf("Server does not support Range request: %s", res.Status))
	default:
		res.Body.Close()
		return nil, fmt.Errorf("HTTP error: %s", res.Status)
	}
}
