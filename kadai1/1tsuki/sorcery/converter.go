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

// converter converts a single image file with specific extension into another.
// See imgExt for supported image file extensions.
type converter struct {
	Path string
}

func (c *converter) convert(to imgExt) (string, error) {
	img, err := c.decode()
	if err != nil {
		return "", err
	}

	return c.encode(img, to)
}

func (c *converter) decode() (image.Image, error) {
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

func (c *converter) encode(m image.Image, to imgExt) (string, error) {
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

func (c *converter) extString() string {
	return trimExt(filepath.Ext(c.Path))
}

func (c *converter) ext() (imgExt, error) {
	return ImgExt(c.extString())
}

func (c *converter) out(to imgExt) string {
	return c.Path[0:len(c.Path)-len(c.extString())] + to.String()
}
