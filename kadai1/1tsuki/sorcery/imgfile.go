package sorcery

import (
	"fmt"
	"image"
	"os"
	"image/jpeg"
	"image/png"
	"image/gif"
	"path/filepath"
)

// imgFile converts a single image file with specific extension into another.
// See imgExt for supported image file extensions.
type imgFile struct {
	Path string
}

func (c *imgFile) convertTo(to imgExt) (string, error) {
	img, err := c.decode()
	if err != nil {
		return "", err
	}

	return c.encode(img, to)
}

func (c *imgFile) decode() (image.Image, error) {
	from, err := c.ext()
	if err != nil {
		return nil, err
	}

	reader, err := os.Open(c.Path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	switch from {
	case Jpeg:
		return jpeg.Decode(reader)
	case Png:
		return png.Decode(reader)
	case Gif:
		return gif.Decode(reader)
	default:
		return nil, fmt.Errorf("unsupported source extension")
	}
}

func (c *imgFile) encode(m image.Image, to imgExt) (string, error) {
	out := c.out(to)
	writer, err := os.Create(out)
	if err != nil {
		return "", err
	}
	defer writer.Close()

	switch to {
	case Jpeg:
		return out, jpeg.Encode(writer, m, nil)
	case Png:
		return out, png.Encode(writer, m)
	case Gif:
		return out, gif.Encode(writer, m, nil)
	default:
		return "", fmt.Errorf("unknown destionation extension")
	}
}

func (c *imgFile) extString() string {
	return trimExt(filepath.Ext(c.Path))
}

func (c *imgFile) ext() (imgExt, error) {
	return ImgExt(c.extString())
}

func (c *imgFile) out(to imgExt) string {
	return c.Path[0:len(c.Path)-len(c.extString())] + to.String()
}
