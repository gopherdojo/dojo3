package convertImage

import (
	"testing"
	"image"
	"bytes"
	"net/http"
	"image/color"
	"image/jpeg"
	"image/png"
	"image/gif"
)

var flagTests = []struct {
	in  string
	out string
}{
	{"jpeg", "gif"},
	{"jpeg", "png"},
	{"png", "gif"},
	{"png", "jpeg"},
	{"gif", "png"},
	{"gif", "jpeg"},
}

func TestConvert(t *testing.T) {
	for _, c := range flagTests {
		var input bytes.Buffer
		img := image.NewRGBA(image.Rect(0, 0, 400, 300))
		for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
			for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
				img.Set(j, i, color.RGBA{255, 255, 0, 0})
			}
		}
		switch c.in {
		case "jpeg":
			jpeg.Encode(&input, img, nil)
		case "png":
			png.Encode(&input, img)
		case "gif":
			gif.Encode(&input, img, nil)
		}

		out := Convert(&input, &c.out)

		buffer := make([]byte, 512)
		_, err := out.Read(buffer)
		if err != nil {
			panic(err)
		}

		contentType := http.DetectContentType(buffer)
		if contentType != "image/"+c.out {
			t.Errorf("input: jpeg, but output format is not " + c.out + "output is " + contentType)
		}
	}
}
