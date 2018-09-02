package format

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

type Format struct {
	Decoder Decoder
	Encoder Encoder
}

// NewFormat - 画像を変換するためのDecoderとEncoderを選択してFormatを返す
func NewFormat(srcFormat, dstFormat string) (*Format, error) {
	format := Format{}
	if err := format.setDecoder(srcFormat); err != nil {
		return nil, err
	}
	if err := format.setEncoder(dstFormat); err != nil {
		return nil, err
	}
	return &format, nil
}

// Decoder - Decodeの実装が必要なインターフェース
type Decoder interface {
	Decode(io.Reader) (image.Image, error)
}

// Encoder - Encodeの実装が必要なインターフェース
type Encoder interface {
	Encode(io.Writer, image.Image) error
}

type JPEG struct{}

func (j *JPEG) Decode(srcFile io.Reader) (image.Image, error) {
	return jpeg.Decode(srcFile)
}

func (j *JPEG) Encode(dstFile io.Writer, img image.Image) error {
	err := jpeg.Encode(dstFile, img, nil)
	return err
}

type PNG struct{}

func (p *PNG) Decode(srcFile io.Reader) (image.Image, error) {
	return png.Decode(srcFile)
}

func (p *PNG) Encode(dstFile io.Writer, img image.Image) error {
	err := png.Encode(dstFile, img)
	return err
}

type GIF struct{}

func (p *GIF) Decode(srcFile io.Reader) (image.Image, error) {
	return gif.Decode(srcFile)
}

func (p *GIF) Encode(dstFile io.Writer, img image.Image) error {
	err := gif.Encode(dstFile, img, nil)
	return err
}

func (f *Format) setDecoder(srcFormat string) error {
	switch srcFormat {
	case "jpeg", "jpg":
		f.Decoder = &JPEG{}
	case "png":
		f.Decoder = &PNG{}
	case "gif":
		f.Decoder = &GIF{}
	default:
		return fmt.Errorf("Supported formats are only jpeg, jpg or png")
	}
	return nil
}

func (f *Format) setEncoder(dstFormat string) error {
	switch dstFormat {
	case "jpeg", "jpg":
		f.Encoder = &JPEG{}
	case "png":
		f.Encoder = &PNG{}
	case "gif":
		f.Encoder = &GIF{}
	default:
		return fmt.Errorf("Supported formats are only jpeg, jpg or png")
	}
	return nil
}
