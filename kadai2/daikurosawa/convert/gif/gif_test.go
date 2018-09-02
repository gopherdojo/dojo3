package gif_test

import (
	"image/jpeg"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/gif"
	"github.com/gopherdojo/dojo3/kadai2/daikurosawa/di"
)

func TestGif_Init(t *testing.T) {
	_, ok := di.Converts["gif"]
	if ok == false {
		t.Fatal("failed not register to di.Converts")
	}
}

func TestGif_Decode(t *testing.T) {
	gifFile, err := os.Open("./../../testdata/gopher.gif")
	if err != nil {
		t.Fatal("failed open gif image", err)
	}
	defer gifFile.Close()
	var gif gif.Gif
	if _, err := gif.Decode(gifFile); err != nil {
		t.Fatal("failed decode gif image", err)
	}
}

func TestGif_Encode(t *testing.T) {
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

	var gif gif.Gif
	if err := gif.Encode(tempFile, img); err != nil {
		t.Fatal("failed encode gif image", err)
	}
}
