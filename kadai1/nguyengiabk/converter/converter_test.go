package converter_test

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"testing"

	"github.com/gopherdojo/dojo3/kadai1/nguyengiabk/converter"
)

func Example() {
	decoder := &converter.JPEG{}
	encoder := &converter.PNG{}
	converter := converter.Converter{Decoder: decoder, Encoder: encoder}
	converter.Run("../testdata/jpg")
}

var testFixtures = []struct {
	description   string
	converter     converter.Converter
	directoryName string
	inputFile     string
	inputType     string
	outputFile    string
}{
	{
		"Test convert jpg to png",
		converter.Converter{Decoder: &converter.JPEG{}, Encoder: &converter.PNG{}},
		"testdata",
		"image.jpg",
		"jpg",
		"image.png",
	},
	{
		"Test convert png to gif",
		converter.Converter{Decoder: &converter.PNG{}, Encoder: &converter.GIF{NumColors: 256}},
		"testdata",
		"image.png",
		"png",
		"image.gif",
	},
	{
		"Test convert gif to jpg",
		converter.Converter{Decoder: &converter.GIF{}, Encoder: &converter.JPEG{Quality: 100}},
		"testdata",
		"image.gif",
		"gif",
		"image.jpg",
	},
}

func TestConverter(t *testing.T) {
	const fileMode = 0777
	for _, tt := range testFixtures {
		if err := os.MkdirAll(tt.directoryName, fileMode); err != nil {
			t.Errorf("Cannot create directory, error: %s", err)
		}
		img := createImage()
		f, err := os.OpenFile(path.Join(tt.directoryName, tt.inputFile), os.O_WRONLY|os.O_CREATE, fileMode)
		if err != nil {
			t.Errorf("Cannot write to input file, error: %s", err)
		}
		switch tt.inputType {
		case "jpg":
			jpeg.Encode(f, img, nil)
		case "png":
			png.Encode(f, img)
		case "gif":
			gif.Encode(f, img, nil)
		}
		f.Close()
		if err := tt.converter.Run(tt.directoryName); err != nil {
			t.Errorf("Error when run converter, error: %s", err)
		}
		// check output file exists
		if _, err := os.Stat(path.Join(tt.directoryName, tt.outputFile)); os.IsNotExist(err) {
			t.Errorf("Expected output file %s is not exist", tt.outputFile)
		}
		os.RemoveAll(tt.directoryName)
	}
}

func createImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			img.Set(i, j, color.RGBA{255, 0, 0, 255})
		}
	}
	return img
}
