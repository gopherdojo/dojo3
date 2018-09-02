package format

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"reflect"
	"testing"
)

var (
	formatOptions = []struct {
		src string
		dst string
	}{
		{"jpg", "jpeg"}, {"jpg", "png"},
		{"jpeg", "jpg"}, {"jpeg", "png"},
		{"png", "jpg"}, {"png", "jpeg"},
		{"jpg", "gif"}, {"gif", "png"},
		{"jpeg", "gif"}, {"gif", "jpeg"},
		{"png", "gif"}, {"gif", "jpg"},
	}
	formatOptionsNG = []struct {
		src string
		dst string
	}{
		{"hoge", "fuga"}, {"hoge", "jpg"},
		{"png", "hoge"}, {"gif", "hoge"},
		{"", ""},
	}
)

func TestNewFormatOK(t *testing.T) {
	var expectedDecoderType reflect.Type
	var expectedEncoderType reflect.Type
	for _, formatOption := range formatOptions {
		switch formatOption.src {
		case "jpg", "jpeg":
			expectedDecoderType = reflect.ValueOf(&JPEG{}).Type()
		case "png":
			expectedDecoderType = reflect.ValueOf(&PNG{}).Type()
		case "gif":
			expectedDecoderType = reflect.ValueOf(&GIF{}).Type()
		}
		switch formatOption.dst {
		case "jpg", "jpeg":
			expectedEncoderType = reflect.ValueOf(&JPEG{}).Type()
		case "png":
			expectedEncoderType = reflect.ValueOf(&PNG{}).Type()
		case "gif":
			expectedEncoderType = reflect.ValueOf(&GIF{}).Type()
		}
		format, err := NewFormat(formatOption.src, formatOption.dst)
		if err != nil {
			fmt.Println(err)
		}
		resultDecoderType := reflect.ValueOf(format.Decoder).Type()
		resultEncoderType := reflect.ValueOf(format.Encoder).Type()
		if expectedDecoderType != resultDecoderType {
			t.Errorf("expected : %s\n", expectedDecoderType)
			t.Errorf("result : %s\n", resultDecoderType)
		}
		if expectedEncoderType != resultEncoderType {
			t.Errorf("expected : %s\n", expectedEncoderType)
			t.Errorf("result : %s\n", resultEncoderType)
		}

	}
}

func TestNewFormatNG(t *testing.T) {
	for _, formatOption := range formatOptionsNG {
		format, err := NewFormat(formatOption.src, formatOption.dst)
		expecedError := fmt.Errorf("Supported formats are only jpeg, jpg or png")
		if format != nil {
			t.Error(format)
		}
		if err.Error() != expecedError.Error() {
			t.Error(expecedError)
			t.Error(err)
		}
	}
}

func TestJpegDecode(t *testing.T) {
	format, _ := NewFormat("jpeg", "png")
	t.Run("OK", func(t *testing.T) {
		jpegImg, _ := os.Open("../images/gopher_kun.jpg")
		_, err := format.Decoder.Decode(jpegImg)
		if err != nil {
			t.Error(err)
		}
		jpegImg.Close()
	})
	t.Run("NG", func(t *testing.T) {
		pngImg, _ := os.Open("../images/gopher_kun.png")
		_, err := format.Decoder.Decode(pngImg)
		if err == nil {
			t.Error("expected: JPEG Decoder can not treat PNG image")
		}
		pngImg.Close()
	})
}

func TestJpegEncode(t *testing.T) {
	t.Helper()
	format, _ := NewFormat("png", "jpg")
	testNewFile, _ := os.Create("../images/gopher_kun.jpeg")
	testBaseFile, _ := os.Open("../images/gopher_kun.png")
	testBaseImage, _ := png.Decode(testBaseFile)
	err := format.Encoder.Encode(testNewFile, testBaseImage)
	testNewFile.Close()
	if err != nil {
		t.Error(err)
	}
	testNewFile, _ = os.Open("../images/gopher_kun.jpeg")
	_, formatStr, _ := image.Decode(testNewFile)
	if formatStr != "jpeg" {
		t.Error(formatStr)
	}
}

func TestPngDecode(t *testing.T) {
	format, _ := NewFormat("png", "jpeg")
	t.Run("OK", func(t *testing.T) {
		pngImg, _ := os.Open("../images/gopher_kun.png")
		_, err := format.Decoder.Decode(pngImg)
		if err != nil {
			t.Error(err)
		}
		pngImg.Close()
	})
	t.Run("NG", func(t *testing.T) {
		jpegImg, _ := os.Open("../images/gopher_kun.jpeg")
		_, err := format.Decoder.Decode(jpegImg)
		if err == nil {
			t.Error("expected: PNG Decoder can not treat JPEG image")
		}
		jpegImg.Close()
	})
}
func TestPngEncode(t *testing.T) {
	t.Helper()
	format, _ := NewFormat("jpg", "png")
	testNewFile, _ := os.Create("../images/gopher_kun.png")
	testBaseFile, _ := os.Open("../images/gopher_kun.jpg")
	testBaseImage, _ := jpeg.Decode(testBaseFile)
	err := format.Encoder.Encode(testNewFile, testBaseImage)
	testNewFile.Close()
	if err != nil {
		t.Error(err)
	}
	testNewFile, _ = os.Open("../images/gopher_kun.png")
	_, formatStr, _ := image.Decode(testNewFile)
	if formatStr != "png" {
		t.Error(formatStr)
	}
}

func TestGIFDecode(t *testing.T) {
	format, _ := NewFormat("gif", "jpeg")
	t.Run("OK", func(t *testing.T) {
		gifImg, _ := os.Open("../images/gopher_kun.gif")
		_, err := format.Decoder.Decode(gifImg)
		if err != nil {
			t.Error(err)
		}
		gifImg.Close()
	})
	t.Run("NG", func(t *testing.T) {
		jpegImg, _ := os.Open("../images/gopher_kun.jpeg")
		_, err := format.Decoder.Decode(jpegImg)
		if err == nil {
			t.Error("expected: GIF Decoder can not treat JPEG image")
		}
		jpegImg.Close()
	})
}
func TestGIFEncode(t *testing.T) {
	t.Helper()
	format, _ := NewFormat("jpg", "gif")
	testNewFile, _ := os.Create("../images/gopher_kun.gif")
	testBaseFile, _ := os.Open("../images/gopher_kun.jpg")
	testBaseImage, _ := jpeg.Decode(testBaseFile)
	err := format.Encoder.Encode(testNewFile, testBaseImage)
	testNewFile.Close()
	if err != nil {
		t.Error(err)
	}
	testNewFile, _ = os.Open("../images/gopher_kun.gif")
	_, formatStr, _ := image.Decode(testNewFile)
	if formatStr != "gif" {
		t.Error(formatStr)
	}
}
