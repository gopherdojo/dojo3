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

func (ic *ImageConverter) openImage(path string) (io.ReadCloser, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (ic *ImageConverter) prepareImage(path string) (io.WriteCloser, error) {
	output, err := os.Create(rep.ReplaceAllString(path, ic.to))
	if err != nil {
		return nil, err
	}
	return output, err
}

// ConvertImage converts image to another extension from provided extension
func (ic *ImageConverter) ConvertImage(path string) (err error) {
	var img image.Image
	var input io.ReadCloser
	var output io.WriteCloser

	ext := rep.FindString(path)
	if ext == "" || ext != ic.from {
		return
	}

	input, err = ic.openImage(path)
	if err != nil {
		return
	}
	output, err = ic.prepareImage(path)
	if err != nil {
		return
	}
	defer input.Close()
	defer output.Close()

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
