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

type ImageConverter struct {
	data *image.Image
}

// decode reads a JPG/GIF/PNG image from r and returns the first embedded
// image as an image.Image.
func (ic *ImageConverter) decode(r io.Reader, fromExt string) error {

	if ic == nil {
		return errors.New("ImageConverter is nil receiver.")
	}

	var img image.Image
	var err error

	switch fromExt {
	case "gif":
		img, err = gif.Decode(r)
	case "jpg":
		img, err = jpeg.Decode(r)
	case "png":
		img, err = png.Decode(r)
	default:
		return fmt.Errorf("%v is not supported.", fromExt)
	}

	if err != nil {
		return err
	}

	ic.data = &img

	return nil
}

// Encode writes the Image m to w in JPG/GIF/PNG format.
func (ic *ImageConverter) encode(w io.Writer, toExt string) error {

	if ic == nil {
		return errors.New("ImageConverter is nil receiver.")
	}

	if ic.data == nil {
		return errors.New("Image data is nil.")
	}

	var err error

	switch toExt {
	case "gif":
		err = gif.Encode(w, *ic.data, nil)
	case "jpg":
		err = jpeg.Encode(w, *ic.data, nil)
	case "png":
		err = png.Encode(w, *ic.data)
	}

	if err != nil {
		return err
	}

	return nil
}

// Convert the image to the specified format.
func (ic *ImageConverter) Convert(r io.Reader, w io.Writer, fromExt, toExt string) error {

	if ic == nil {
		return errors.New("ImageConverter is nil receiver.")
	}

	if err := ic.decode(r, fromExt); err != nil {
		return err
	}

	if err := ic.encode(w, toExt); err != nil {
		return err
	}

	return nil
}
