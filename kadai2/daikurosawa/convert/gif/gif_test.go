package gif_test

import (
	"image"
	"image/gif"
	"io/ioutil"
	"os"
	"testing"

	conv "github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/gif"
)

const tempFileName = "test_image"

func makeTestImage() *image.RGBA {
	var (
		x      = 0
		y      = 0
		width  = 50
		height = 50
	)

	return image.NewRGBA(image.Rect(x, y, width, height))
}

func TestGif_Decode(t *testing.T) {
	tempFile, err := ioutil.TempFile("./", tempFileName)
	if err != nil {
		t.Fatal("failed create temp file.", err)
	}
	defer func() {
		tempFile.Close()
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Fatal("failed remove temp file.", err)
		}
	}()

	img := makeTestImage()
	if err := gif.Encode(tempFile, img, nil); err != nil {
		t.Fatal("failed gif encode.", err)
	}

	_, err = conv.Gif{}.Decode(tempFile)
	if err != nil {
		t.Fatal("failed gif decode.", err)
	}
}

func TestGif_Encode(t *testing.T) {
	tempFile, err := ioutil.TempFile("./", tempFileName)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		tempFile.Close()
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Fatal("failed remove temp file", err)
		}
	}()

	img := makeTestImage()
	err = conv.Gif{}.Encode(tempFile, img)
	if err != nil {
		t.Fatal(err)
	}
}
