package imgconv

import (
	"fmt"
	"testing"

	"github.com/dojo3/kadai2/nKumaya/format"
)

var (
	formatJPEGtoPNG       *format.Format
	formatPNGtoJPEG       *format.Format
	formatJPEGtoGIF       *format.Format
	formatGIFtoJPEG       *format.Format
	formatGIFtoPNG        *format.Format
	formatPNGtoGIF        *format.Format
	ConverterJPEGtoPNG    *Converter
	ConverterPNGtoJPEG    *Converter
	ConverterJPEGtoGIF    *Converter
	ConverterGIFtoJPEG    *Converter
	ConverterPNGtoGIF     *Converter
	ConverterGIFtoPNG     *Converter
	ConverterNGNotFound   *Converter
	ConverterNGPermission *Converter
	ConverterNGDecode     *Converter
)

func TestConvert(t *testing.T) {
	t.Run("Jpeg to PNG", func(t *testing.T) {
		formatJPEGtoPNG, _ = format.NewFormat("jpeg", "png")
		ConverterJPEGtoPNG = NewConverter("../images/gopher_kun.jpg", "../images/gopher_kun.png", *formatJPEGtoPNG)
		err := ConverterJPEGtoPNG.Convert()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("PNG to Jpeg", func(t *testing.T) {
		formatPNGtoJPEG, _ = format.NewFormat("png", "jpeg")
		ConverterPNGtoJPEG = NewConverter("../images/gopher_kun.png", "../images/gopher_kun.jpeg", *formatPNGtoJPEG)
		err := ConverterPNGtoJPEG.Convert()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("JPEG to GIF", func(t *testing.T) {
		formatJPEGtoGIF, _ = format.NewFormat("jpeg", "gif")
		ConverterJPEGtoGIF = NewConverter("../images/gopher_kun.jpeg", "../images/gopher_kun.gif", *formatJPEGtoGIF)
		err := ConverterJPEGtoGIF.Convert()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("GIF to JPEG", func(t *testing.T) {
		formatGIFtoJPEG, _ = format.NewFormat("gif", "jpeg")
		ConverterGIFtoJPEG = NewConverter("../images/gopher_kun.gif", "../images/gopher_kun.jpeg", *formatGIFtoJPEG)
		err := ConverterGIFtoJPEG.Convert()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("PNG to GIF", func(t *testing.T) {
		formatPNGtoGIF, _ = format.NewFormat("png", "gif")
		ConverterPNGtoGIF = NewConverter("../images/gopher_kun.png", "../images/gopher_kun.gif", *formatPNGtoGIF)
		err := ConverterPNGtoGIF.Convert()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("GIF to PNG", func(t *testing.T) {
		formatGIFtoPNG, _ = format.NewFormat("gif", "jpeg")
		ConverterGIFtoPNG = NewConverter("../images/gopher_kun.gif", "../images/gopher_kun.png", *formatGIFtoPNG)
		err := ConverterGIFtoPNG.Convert()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("File not found", func(t *testing.T) {
		ConverterNGNotFound = NewConverter("../images/unexist.png", "../images/hogehoge.jpeg", *formatPNGtoJPEG)
		err := ConverterNGNotFound.Convert()
		expecedError := fmt.Errorf("open %s: no such file or directory", ConverterNGNotFound.baseFile)
		if err.Error() != expecedError.Error() {
			t.Error(err.Error())
			t.Error(expecedError.Error())
		}
	})
	t.Run("Permission error", func(t *testing.T) {
		ConverterNGPermission = NewConverter("../images/gopher_kun.gif", "/etc/hogehoge.jpg", *formatGIFtoJPEG)
		err := ConverterNGPermission.Convert()
		expecedError := fmt.Errorf("open %s: permission denied", ConverterNGPermission.newFile)
		if err.Error() != expecedError.Error() {
			t.Error(err.Error())
			t.Error(expecedError.Error())
		}
	})
	t.Run("Decode error", func(t *testing.T) {
		ConverterNGDecode = NewConverter("../images/gopher_kun.gif", "/etc/hogehoge.jpg", *formatJPEGtoGIF)
		err := ConverterNGDecode.Convert()
		expecedError := fmt.Errorf("invalid JPEG format: missing SOI marker")
		if err.Error() != expecedError.Error() {
			t.Error(err.Error())
			t.Error(expecedError.Error())
		}
	})
}
