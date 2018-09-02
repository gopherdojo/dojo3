package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// Img represents a image file.
type Img struct {
	m      image.Image
	path   string
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

	i := Img{
		m,
		path,
		format,
	}
	return &i, nil
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
