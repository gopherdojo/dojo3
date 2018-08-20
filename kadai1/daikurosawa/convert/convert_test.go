package convert_test

import (
	"image"
	"image/jpeg"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert"
	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/option"
)

const TestImageFileName = "test_image.jpg"

type Mock struct {
}

// Decode mock
func (Mock) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

// Encode mock
func (Mock) Encode(w io.Writer, m image.Image) error {
	return nil
}

var testFixtures = []struct {
	fromExtension string
	toExtension   string
}{
	{
		"jpg",
		"png",
	},
	{
		"jpg",
		"gif",
	},
	{
		"jpg",
		"jpg",
	},
}

func init() {
	var (
		x      = 0
		y      = 0
		width  = 100
		height = 50
	)

	// make test image
	img := image.NewRGBA(image.Rect(x, y, width, height))
	file, _ := os.Create(TestImageFileName)
	defer file.Close()
	if err := jpeg.Encode(file, img, nil); err != nil {
		panic(err)
	}

	// set mock
	convert.Register("jpg", Mock{})
	convert.Register("gif", Mock{})
	convert.Register("png", Mock{})
}

func TestConvert_Convert(t *testing.T) {
	defer func() {
		if exist := isExist(TestImageFileName); exist {
			if err := os.Remove(TestImageFileName); err != nil {
				panic(err)
			}
		}
	}()

	for _, fixture := range testFixtures {
		option := &option.Option{FromExtension: fixture.fromExtension, ToExtension: fixture.toExtension}
		convert := convert.NewConvert(option)
		if err := convert.Convert(TestImageFileName); err != nil {
			t.Fatalf("failed convert test: %#v", err)
		}

		// test exists output file
		outputFile := strings.TrimSuffix(TestImageFileName, fixture.fromExtension) + fixture.toExtension
		if exist := isExist(outputFile); exist == false {
			t.Fatal("failed: not found output file")
		}

		// delete output file
		if err := os.Remove(outputFile); err != nil {
			panic(err)
		}
	}
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
