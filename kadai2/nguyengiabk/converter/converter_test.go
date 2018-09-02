package converter_test

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/nguyengiabk/converter"
)

func Example() {
	decoder := &converter.JPEG{}
	encoder := &converter.PNG{}
	converter := converter.Converter{Decoder: decoder, Encoder: encoder}
	converter.Run("../testdata/jpg", os.Stdout)
}

type testCase struct {
	name          string
	converter     converter.Converter
	directoryName string
	inputFile     string
	inputType     string
	outputFile    string
	outputLog     string
}

var testFixtures = []testCase{
	{
		"Test convert jpg to png",
		converter.Converter{Decoder: &converter.JPEG{}, Encoder: &converter.PNG{}},
		"testdata",
		"image.jpg",
		"jpg",
		"image.png",
		"",
	},
	{
		"Test convert png to gif",
		converter.Converter{Decoder: &converter.PNG{}, Encoder: &converter.GIF{NumColors: 256}},
		"testdata",
		"image.png",
		"png",
		"image.gif",
		"",
	},
	{
		"Test convert gif to jpg",
		converter.Converter{Decoder: &converter.GIF{}, Encoder: &converter.JPEG{Quality: 100}},
		"testdata",
		"image.gif",
		"gif",
		"image.jpg",
		"",
	},
	{
		"Test undecodable file",
		converter.Converter{Decoder: &converter.JPEG{}, Encoder: &converter.PNG{}},
		"testdata",
		"image.jpg",
		"gif",
		"",
		"Cannot decode file testdata/image.jpg, continue processing\n",
	},
}

func TestConverter(t *testing.T) {
	for _, tc := range testFixtures {
		t.Run(tc.name, func(t *testing.T) {
			closer := createTestInput(t, tc)
			buf := &bytes.Buffer{}
			if err := tc.converter.Run(tc.directoryName, buf); err != nil {
				t.Errorf("Error when run converter, error: %s", err)
			}
			if !bytes.Equal(buf.Bytes(), []byte(tc.outputLog)) {
				t.Errorf("Run return unexpected error, actual = %s, expected = %s", buf.String(), tc.outputLog)
			}
			// check output file exists
			if _, err := os.Stat(path.Join(tc.directoryName, tc.outputFile)); os.IsNotExist(err) {
				t.Errorf("Expected output file %s is not exist", tc.outputFile)
			}
			closer()
		})
	}
}

func createTestInput(t *testing.T, tc testCase) func() {
	t.Helper()
	const fileMode = 0777
	if err := os.MkdirAll(tc.directoryName, fileMode); err != nil {
		t.Errorf("Cannot create directory, error: %s", err)
	}
	img := createImage(t)
	f, err := os.OpenFile(path.Join(tc.directoryName, tc.inputFile), os.O_WRONLY|os.O_CREATE, fileMode)
	if err != nil {
		t.Errorf("Cannot create input file, error: %s", err)
	}
	switch tc.inputType {
	case "jpg":
		err = jpeg.Encode(f, img, nil)
	case "png":
		err = png.Encode(f, img)
	case "gif":
		err = gif.Encode(f, img, nil)
	}
	if err != nil {
		t.Errorf("Cannot write to input file, error: %s", err)
	}
	f.Close()
	return func() {
		os.RemoveAll(tc.directoryName)
	}
}

func createImage(t *testing.T) image.Image {
	t.Helper()
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			img.Set(i, j, color.RGBA{255, 0, 0, 255})
		}
	}
	return img
}
