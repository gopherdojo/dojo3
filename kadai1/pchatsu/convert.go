package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	ErrUnMatchedFormat = fmt.Errorf("unmatched with the target format")
	ErrUnKnownFormat   = fmt.Errorf("unknown format")
)

//type dir struct {
//	path string
//}
//
//func NewDir(path string) *dir {
//	return &dir{
//		strings.TrimRight(path, "/"),
//	}
//}

//func (d *dir) Convert(src string, dst string) {
func Convert(d string, src string, dst string) {
	wg := &sync.WaitGroup{}
	filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func() {
			img, err := NewImg(path, src)
			if err != nil {
				fmt.Fprintln(os.Stderr, "imgconv:", err.Error(), path)
				wg.Done()
				return
			}

			if err := img.convert(dst); err != nil {
				fmt.Fprintln(os.Stderr, "imgconv:", err.Error(), path)
				wg.Done()
				return
			}

			wg.Done()
		}()

		return nil
	})
	wg.Wait()
}

type img struct {
	m      image.Image
	path   string
	ext    string
	format string
}

func NewImg(path string, format string) (*img, error) {
	ext := filepath.Ext(path)
	f := strings.ToLower(ext[1:])
	if f != format {
		if f != "jpg" || format != "jpeg" {
			return nil, ErrUnMatchedFormat
		}
	}

	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	m, magic, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	if magic != format {
		return nil, ErrUnMatchedFormat
	}

	i := img{
		m,
		path,
		ext,
		format,
	}
	return &i, nil
}

func (i *img) convert(dst string) error {
	var newExt string
	if dst == "jpeg" {
		newExt = ".jpg"
	} else {
		newExt = "." + dst
	}

	newPath := strings.TrimRight(i.path, i.ext) + newExt

	f, err := os.OpenFile(newPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := i.encode(f, dst); err != nil {
		return err
	}

	return nil
}

func (i *img) encode(w io.Writer, dst string) error {
	switch dst {
	case "gif":
		if err := gif.Encode(w, i.m, nil); err != nil {
			return err
		}
	case "jpeg":
		if err := jpeg.Encode(w, i.m, nil); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(w, i.m); err != nil {
			return err
		}

	default:
		return ErrUnKnownFormat
	}

	return nil
}
