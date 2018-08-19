package imgconv

import (
	"image/png"
	"fmt"
	"image"
	"io"
	"errors"
	"image/jpeg"
)

type supportFormat string

type ImageConverter struct {
	data *image.Image
}

//ユーザー定義型
const (
	jpegType supportFormat = "jpg"
	pngType  supportFormat = "png"
)

//実際に変換する
func (ic *ImageConverter) Convert(r io.Reader, w io.Writer, fromExt, toExt string) error {
	if ic == nil {
		return errors.New("ImageConverter is nil")
	}
	if err := ic.decode(r, supportFormat(fromExt)); err != nil {
		return err
	}
	if err := ic.encode(w, supportFormat(toExt)); err != nil {
		return err
	}
	return nil
}

func (sf supportFormat) decode(r io.Reader) (image.Image, error) {
	switch sf {
	case jpegType:
		return jpeg.Decode(r)
	case pngType:
		return png.Decode(r)
	default:
		return nil, fmt.Errorf("%v is not supported", sf)
	}
}

func (sf supportFormat) encode(w io.Writer, img image.Image) error {
	switch sf {
	case jpegType:
		return jpeg.Encode(w, img, nil)
	case pngType:
		return png.Encode(w, img)
	default:
		return fmt.Errorf("%v is not supported", sf)
	}
}

func (ic *ImageConverter) decode(r io.Reader, fromExt supportFormat) error {
	if ic == nil {
		return errors.New("ImageConverter is nil")
	}
	img, err := fromExt.decode(r)
	if err != nil {
		return err
	}
	ic.data = &img
	return nil
}

func (ic *ImageConverter) encode(w io.Writer, toExt supportFormat) error {
	if ic == nil {
		return errors.New("ImageConverter is nil")
	}
	if ic.data == nil {
		return errors.New("data is nil")
	}
	if err := toExt.encode(w, *ic.data); err != nil {
		return err
	}
	return nil
}

