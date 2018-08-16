package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
	"golang.org/x/sync/errgroup"
)

// Convert converts the format of image files.
// d is the target dir and it's children images are recursively converted except hidden files.
// When a converting process occurs error, the others don't stop.
func Convert(d string, src string, dst string) error {
	eg := errgroup.Group{}
	if err := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if !isTargetFile(info, path, src) {
			return nil
		}

		eg.Go(func() error {
			img, err := NewImg(path, src)
			if err != nil {
				return err
			}

			if err := img.Convert(dst); err != nil {
				return err
			}
			return nil
		})

		return nil
	}); err != nil {
		return err
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

// isTargetFile returns a boolean indicating whether the file should be converted.
func isTargetFile(info os.FileInfo, path string, format string) bool {
	// skip dir
	if info.IsDir() {
		return false
	}
	// skip hidden files
	if filepath.Base(path)[0:1] == "." {
		return false
	}

	// skip unmatched files with the target format
	ext := filepath.Ext(path)
	f := strings.ToLower(strings.TrimPrefix(ext, "."))
	if f != format {
		if f != "jpg" || format != "jpeg" {
			return false
		}
	}

	return true
}

// Img represents a image file.
type Img struct {
	m      image.Image
	path   string
	ext    string
	format string
}

// NewImg generates Img from given path and format.
func NewImg(path string, format string) (*Img, error) {
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
		return nil, fmt.Errorf("%s unmatched with the target format", path)
	}

	ext := filepath.Ext(path)

	i := Img{
		m,
		path,
		ext,
		format,
	}
	return &i, nil
}

// Convert converts the image to specified format.
// New file name is based on the name of source file and it overwrites if the same name file exists.
func (i *Img) Convert(dst string) error {
	var newExt string
	if dst == "jpeg" {
		newExt = ".jpg"
	} else {
		newExt = "." + dst
	}

	newPath := strings.TrimSuffix(i.path, i.ext) + newExt

	f, err := os.OpenFile(newPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := i.Encode(f, dst); err != nil {
		return err
	}

	return nil
}

// Encode encodes specified format and writes to given Writer.
func (i *Img) Encode(w io.Writer, dst string) error {
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
		return fmt.Errorf("%s unknown format", dst)
	}

	return nil
}
