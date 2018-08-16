package imgconv

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

type supportFormat string

const (
	gifType supportFormat = "gif"
	jpgType supportFormat = "jpg"
	pngType supportFormat = "png"
)

func (sf supportFormat) decode(r io.Reader) (image.Image, error) {
	switch sf {
	case gifType:
		return gif.Decode(r)
	case jpgType:
		return jpeg.Decode(r)
	case pngType:
		return png.Decode(r)
	default:
		return nil, fmt.Errorf("%v is not supported.", sf)
	}
}

func (sf supportFormat) encode(w io.Writer, img image.Image) error {
	switch sf {
	case gifType:
		return gif.Encode(w, img, nil)
	case jpgType:
		return jpeg.Encode(w, img, nil)
	case pngType:
		return png.Encode(w, img)
	default:
		return fmt.Errorf("%v is not supported.", sf)
	}
}

type ImageConverter struct {
	data *image.Image
}

// decode reads a JPG/GIF/PNG image from r and returns the first embedded
// image as an image.Image.
func (ic *ImageConverter) decode(r io.Reader, fromExt supportFormat) error {

	if ic == nil {
		return errors.New("ImageConverter is nil receiver.")
	}

	img, err := fromExt.decode(r)
	if err != nil {
		return err
	}

	ic.data = &img

	return nil
}

// Encode writes the Image m to w in JPG/GIF/PNG format.
func (ic *ImageConverter) encode(w io.Writer, toExt supportFormat) error {

	if ic == nil {
		return errors.New("ImageConverter is nil receiver.")
	}

	if ic.data == nil {
		return errors.New("Image data is nil.")
	}

	if err := toExt.encode(w, *ic.data); err != nil {
		return err
	}

	return nil
}

// Convert the image to the specified format.
func (ic *ImageConverter) Convert(r io.Reader, w io.Writer, fromExt, toExt string) error {

	if ic == nil {
		return errors.New("ImageConverter is nil receiver.")
	}

	if err := ic.decode(r, supportFormat(fromExt)); err != nil {
		return err
	}

	if err := ic.encode(w, supportFormat(toExt)); err != nil {
		return err
	}

	return nil
}
