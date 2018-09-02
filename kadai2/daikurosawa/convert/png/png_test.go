package png_test

import (
	"image/jpeg"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/png"
	"github.com/gopherdojo/dojo3/kadai2/daikurosawa/di"
)

func TestPng_Init(t *testing.T) {
	_, ok := di.Converts["png"]
	if ok == false {
		t.Fatal("failed not register to di.Converts")
	}
}

func TestPng_Decode(t *testing.T) {
	pngFile, err := os.Open("./../../testdata/gopher.png")
	if err != nil {
		t.Fatal("failed open png image", err)
	}
	defer pngFile.Close()
	var png png.Png
	if _, err := png.Decode(pngFile); err != nil {
		t.Fatal("failed decode png image", err)
	}
}

func TestPng_Encode(t *testing.T) {
	jpgFile, err := os.Open("./../../testdata/gopher.jpg")
	if err != nil {
		t.Fatal("failed open jpg image", err)
	}
	defer jpgFile.Close()

	img, err := jpeg.Decode(jpgFile)
	if err != nil {
		t.Fatal("failed decode jpg image", err)
	}

	tempFile, err := ioutil.TempFile("./", "temp_file")
	if err != nil {
		t.Fatal("failed create temp file", err)
	}
	defer os.Remove(tempFile.Name())

	var png png.Png
	if err := png.Encode(tempFile, img); err != nil {
		t.Fatal("failed encode png image", err)
	}
}
