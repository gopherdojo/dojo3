package convertImage

import (
	"testing"
	"image"
		"bytes"
	"net/http"
	"image/color"
	"image/jpeg"
)

func TestConvert(t *testing.T) {
	var input bytes.Buffer
	img := image.NewRGBA(image.Rect(0, 0, 400, 300))
	for i := img.Rect.Min.Y; i<img.Rect.Max.Y;i++ {
		for j := img.Rect.Min.X; j<img.Rect.Max.X; j++ {
			img.Set(j, i, color.RGBA{255, 255, 0, 0})
		}
	}
	jpeg.Encode(&input, img, nil)

	png := "png"
	out := Convert(&input, &png)

	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		panic(err)
	}

	contentType := http.DetectContentType(buffer)
	if contentType != "image/png" {
		t.Errorf("input: jpeg, but output format is not " + png + "output is " + contentType)
	}
}
