package imgconv

import (
	"io"
	"image/jpeg"
	"image/png"
	"image"
	"image/gif"
)

/*
Converter converts file format
*/
type Converter interface {
	// r is io.Reader for src file.
	// w is io.Writer for dest file.
	Convert(r io.Reader, w io.Writer) error
}

// imgConverter is the converter for image files
type imgConverter struct {
	SrcFormat  []string
	DestFormat string
}

// Convert image format
func (c *imgConverter) Convert(r io.Reader, w io.Writer) error {
	// Decode
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	// Encode by Output format
	switch c.DestFormat {
	case "jpg":
		err = jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
	case "png":
		err = png.Encode(w, img)
	case "gif":
		err = gif.Encode(w, img, nil)
	default:
		err = png.Encode(w, img)
	}
	return err
}

// ImgConverter creates Converter specialized for images from Option
func ImgConverter(opt *Option) Converter {
	return &imgConverter{SrcFormat: opt.Input, DestFormat: opt.Output}
}
