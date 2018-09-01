package imageconverter

import (
	"github.com/pkg/errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"regexp"
)

var reg = regexp.MustCompile("(jpg|jpeg|png|gif)$")

// ImageConverter provides image conversion
type ImageConverter struct {
	from string
	to   string
}

// New returns an ImageConverter that has file extensions for conversion
func New(from, to string) (ImageConverter, error) {
	if !reg.MatchString(from) || !reg.MatchString(to) || from == to {
		return ImageConverter{}, errors.New("this extension is not allowed")
	}

	return ImageConverter{
		from,
		to,
	}, nil
}

// ConvertImage converts image to another extension from provided extension
func (ic *ImageConverter) ConvertImage(path string, output io.WriteCloser) (err error) {
	var img image.Image
	println(path)

	ext := reg.FindString(path)
	if ext == "" || ext != ic.from {
		return
	}

	input, err := os.Open(path)
	if err != nil {
		return
	}
	if output == nil {
		output, err = os.Create(reg.ReplaceAllString(path, ic.to))
		if err != nil {
			return
		}
	}
	defer input.Close()
	defer output.Close()

	switch {
	case ext == "jpg" || ext == "jpeg":
		img, err = jpeg.Decode(input)
	case ext == "png":
		img, err = png.Decode(input)
	case ext == "gif":
		img, err = gif.Decode(input)
	default:
		return
	}
	if err != nil {
		return
	}

	switch {
	case ic.to == "jpg" || ic.to == "jpeg":
		err = jpeg.Encode(output, img, nil)
	case ic.to == "png":
		err = png.Encode(output, img)
	case ic.to == "gif":
		err = gif.Encode(output, img, nil)
	}

	return
}
