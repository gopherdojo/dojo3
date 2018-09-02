package jpg_test

import (
	"image/png"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/jpg"
	"github.com/gopherdojo/dojo3/kadai2/daikurosawa/di"
)

func TestJpg_Init(t *testing.T) {
	_, ok := di.Converts["jpg"]
	if ok == false {
		t.Fatal("failed not register to di.Converts")
	}
}

func TestJpg_Decode(t *testing.T) {
	jpgFile, err := os.Open("./../../testdata/gopher.jpg")
	if err != nil {
		t.Fatal("failed open jpg image", err)
	}
	defer jpgFile.Close()
	var jpg jpg.Jpg
	if _, err := jpg.Decode(jpgFile); err != nil {
		t.Fatal("failed decode jpg image", err)
	}
}

func TestJpg_Encode(t *testing.T) {
	pngFile, err := os.Open("./../../testdata/gopher.png")
	if err != nil {
		t.Fatal("failed open png image", err)
	}
	defer pngFile.Close()

	img, err := png.Decode(pngFile)
	if err != nil {
		t.Fatal("failed decode png image", err)
	}

	tempFile, err := ioutil.TempFile("./", "temp_file")
	if err != nil {
		t.Fatal("failed create temp file", err)
	}
	defer os.Remove(tempFile.Name())

	var jpg jpg.Jpg
	if err := jpg.Encode(tempFile, img); err != nil {
		t.Fatal("failed encode jpg image", err)
	}
}
