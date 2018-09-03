package imgconv

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"testing"
)

var (
	rect = image.Rect(0, 0, 1, 1)
	img  = image.NewGray(rect)
)

func TestConvertJpgToGif(t *testing.T) {
	output := new(bytes.Buffer)
	jpgBuf := new(bytes.Buffer)
	jpeg.Encode(jpgBuf, img, nil)
	gifBuf := new(bytes.Buffer)
	gif.Encode(gifBuf, img, nil)

	imgConverter := ImageConverter{}
	imgConverter.Convert(jpgBuf, output, "gif")

	if !bytes.Equal(output.Bytes(), gifBuf.Bytes()) {
		t.Errorf("Output=%v, want=%v", output.Bytes(), gifBuf.Bytes())
	}
}

func TestConvertJpgToPng(t *testing.T) {
	output := new(bytes.Buffer)
	jpgBuf := new(bytes.Buffer)
	jpeg.Encode(jpgBuf, img, nil)
	pngBuf := new(bytes.Buffer)
	png.Encode(pngBuf, img)

	imgConverter := ImageConverter{}
	imgConverter.Convert(jpgBuf, output, "png")

	if !bytes.Equal(output.Bytes(), pngBuf.Bytes()) {
		t.Errorf("Output=%v, want=%v", output.Bytes(), pngBuf.Bytes())
	}
}
