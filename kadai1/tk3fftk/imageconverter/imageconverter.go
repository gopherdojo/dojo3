package imageconverter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"regexp"
)

var rep = regexp.MustCompile("\\.(jpg|jpeg|png|gif)$")

// ImageConverter provides image conversion
type ImageConverter struct {
	from string
	to   string
}

// New returns an ImageConverter that has file extensions for conversion
func New(from, to string) ImageConverter {
	return ImageConverter{
		from: "." + from,
		to:   "." + to,
	}
}

func (ic *ImageConverter) openImages(path string) (io.Reader, io.Writer, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	output, err := os.Create(rep.ReplaceAllString(path, ic.to))
	if err != nil {
		return nil, nil, err
	}

	return input, output, nil
}

// ConvertImage converts image to another extension from provided extension
func (ic *ImageConverter) ConvertImage(path string) (err error) {
	var img image.Image
	var input io.Reader
	var output io.Writer

	ext := rep.FindString(path)
	if ext == "" || ext != ic.from {
		return
	}

	input, output, err = ic.openImages(path)
	if err != nil {
		return
	}

	switch {
	case ext == ".jpg" || ext == ".jpeg":
		img, err = jpeg.Decode(input)
	case ext == ".png":
		img, err = png.Decode(input)
	case ext == ".gif":
		img, err = gif.Decode(input)
	default:
		return
	}
	if err != nil {
		return
	}

	switch {
	case ic.to == ".jpg" || ic.to == ".jpeg":
		err = jpeg.Encode(output, img, nil)
	case ic.to == ".png":
		err = png.Encode(output, img)
	case ic.to == ".gif":
		err = gif.Encode(output, img, nil)
	}

	return
}
