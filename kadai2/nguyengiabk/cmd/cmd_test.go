package cmd_test

import (
	"bytes"
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/nguyengiabk/cmd"
)

func Example() {
	params, err := cmd.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err = cmd.Run(*params, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

var testParseFixtures = []struct {
	name   string
	args   []string
	err    error
	output *cmd.Parameter
}{
	{
		"Test invalid input type",
		[]string{"-i", "doc", "testdata"},
		errors.New("Invalid input type"),
		nil,
	},
	{
		"Test invalid output type",
		[]string{"-o", "doc", "testdata"},
		errors.New("Invalid output type"),
		nil,
	},
	{
		"Test jpeg quality is smaller than minJpgQuality",
		[]string{"-o", "jpg", "-q", "0", "testdata"},
		errors.New("Invalid JPG quality"),
		nil,
	},
	{
		"Test jpeg quality is larger than maxJpgQuality",
		[]string{"-o", "jpg", "-q", "101", "testdata"},
		errors.New("Invalid JPG quality"),
		nil,
	},
	{
		"Test gif numOfColor is smaller than minGifNumColors",
		[]string{"-o", "gif", "-n", "0", "testdata"},
		errors.New("Invalid maximum number of colors for GIF encoding"),
		nil,
	},
	{
		"Test gif numOfColor is larger than maxGifNumColors",
		[]string{"-o", "gif", "-n", "257", "testdata"},
		errors.New("Invalid maximum number of colors for GIF encoding"),
		nil,
	},
	{
		"Test no input specified",
		[]string{},
		errors.New("No input specified"),
		nil,
	},
	{
		"Test default parameters",
		[]string{"testdata"},
		nil,
		&cmd.Parameter{
			InputType:    "jpg",
			OutputType:   "png",
			JpgQuality:   100,
			GifNumColors: 256,
			Path:         []string{"testdata"},
		},
	},
	{
		"Test two directories",
		[]string{"testdata1", "testdata2"},
		nil,
		&cmd.Parameter{
			InputType:    "jpg",
			OutputType:   "png",
			JpgQuality:   100,
			GifNumColors: 256,
			Path:         []string{"testdata1", "testdata2"},
		},
	},
	{
		"Test convert jpg to png",
		[]string{"-i", "jpg", "-o", "png", "testdata"},
		nil,
		&cmd.Parameter{
			InputType:    "jpg",
			OutputType:   "png",
			JpgQuality:   100,
			GifNumColors: 256,
			Path:         []string{"testdata"},
		},
	},
	{
		"Test convert jpg to gif",
		[]string{"-i", "jpg", "-o", "gif", "-n", "50", "testdata"},
		nil,
		&cmd.Parameter{
			InputType:    "jpg",
			OutputType:   "gif",
			JpgQuality:   100,
			GifNumColors: 50,
			Path:         []string{"testdata"},
		},
	},
	{
		"Test convert png to jpg",
		[]string{"-i", "png", "-o", "jpg", "-q", "50", "testdata"},
		nil,
		&cmd.Parameter{
			InputType:    "png",
			OutputType:   "jpg",
			JpgQuality:   50,
			GifNumColors: 256,
			Path:         []string{"testdata"},
		},
	},
	{
		"Test convert png to gif",
		[]string{"-i", "png", "-o", "gif", "-n", "50", "testdata"},
		nil,
		&cmd.Parameter{
			InputType:    "png",
			OutputType:   "gif",
			JpgQuality:   100,
			GifNumColors: 50,
			Path:         []string{"testdata"},
		},
	},
	{
		"Test convert gif to jpg",
		[]string{"-i", "gif", "-o", "jpg", "-q", "50", "testdata"},
		nil,
		&cmd.Parameter{
			InputType:    "gif",
			OutputType:   "jpg",
			JpgQuality:   50,
			GifNumColors: 256,
			Path:         []string{"testdata"},
		},
	},
	{
		"Test convert gif to png",
		[]string{"-i", "gif", "-o", "png", "testdata"},
		nil,
		&cmd.Parameter{
			InputType:    "gif",
			OutputType:   "png",
			JpgQuality:   100,
			GifNumColors: 256,
			Path:         []string{"testdata"},
		},
	},
}

func TestParse(t *testing.T) {
	for _, tc := range testParseFixtures {
		t.Run(tc.name, func(t *testing.T) {
			result, err := cmd.Parse(tc.args)
			if !reflect.DeepEqual(result, tc.output) {
				t.Errorf("Parse return wrong result, input = %v, actual = %v, expected = %v", tc.args, result, tc.output)
			}
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Parse return wrong error, actual error = %v, expected error = %v", err, tc.err)
			}
		})
	}
}

type testCase struct {
	name        string
	params      cmd.Parameter
	outputFiles []string
	outputLog   string
}

var testRunFixtures = []testCase{
	{
		"Test convert jpg to png",
		cmd.Parameter{
			InputType:  "jpg",
			OutputType: "png",
			Path:       []string{"testdata"},
		},
		[]string{"testdata/jpg/input1/gopher1.png", "testdata/jpg/input2/gopher2.png"},
		"",
	},
	{
		"Test convert jpg to gif",
		cmd.Parameter{
			InputType:    "jpg",
			OutputType:   "gif",
			GifNumColors: 100,
			Path:         []string{"testdata"},
		},
		[]string{"testdata/jpg/input1/gopher1.gif", "testdata/jpg/input2/gopher2.gif"},
		"",
	},
	{
		"Test convert png to gif",
		cmd.Parameter{
			InputType:    "png",
			OutputType:   "gif",
			GifNumColors: 100,
			Path:         []string{"testdata"},
		},
		[]string{"testdata/png/input1/gopher1.gif", "testdata/png/input2/gopher2.gif"},
		"Cannot decode file testdata/undecodable/invalid.png, continue processing\n",
	},
	{
		"Test convert png to jpg",
		cmd.Parameter{
			InputType:  "png",
			OutputType: "jpg",
			JpgQuality: 100,
			Path:       []string{"testdata"},
		},
		[]string{"testdata/png/input1/gopher1.jpg", "testdata/png/input2/gopher2.jpg"},
		"Cannot decode file testdata/undecodable/invalid.png, continue processing\n",
	},
	{
		"Test convert gif to jpg",
		cmd.Parameter{
			InputType:  "gif",
			OutputType: "jpg",
			JpgQuality: 100,
			Path:       []string{"testdata"},
		},
		[]string{"testdata/gif/input1/gopher1.jpg", "testdata/gif/input2/gopher2.jpg"},
		"",
	},
	{
		"Test convert gif to png",
		cmd.Parameter{
			InputType:  "gif",
			OutputType: "png",
			Path:       []string{"testdata"},
		},
		[]string{"testdata/gif/input1/gopher1.png", "testdata/gif/input2/gopher2.png"},
		"",
	},
}

func TestRun(t *testing.T) {
	for _, tc := range testRunFixtures {
		t.Run(tc.name, func(t *testing.T) {
			removeFiles(t, tc.outputFiles)
			buf := &bytes.Buffer{}
			if err := cmd.Run(tc.params, buf); err != nil {
				t.Errorf("Error: %s", err)
			}
			if !bytes.Equal(buf.Bytes(), []byte(tc.outputLog)) {
				t.Errorf("Run return unexpected error, actual = %s, expected = %s", buf.String(), tc.outputLog)
			}
			checkOutput(t, tc)
			removeFiles(t, tc.outputFiles)
		})
	}
}

func checkOutput(t *testing.T, tc testCase) {
	t.Helper()
	for _, path := range tc.outputFiles {
		// check output file exists
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Expected output file %s is not exist", path)
		}
		// check output file is in right format
		file, err := os.Open(path)
		if err != nil {
			t.Errorf("Cannot open file %s, error = %v", path, err)
		}
		defer file.Close()

		switch tc.params.OutputType {
		case "jpg":
			_, err = jpeg.Decode(file)
		case "png":
			_, err = png.Decode(file)
		case "gif":
			_, err = gif.Decode(file)
		}

		if err != nil {
			t.Errorf("Output file %s is in wrong format", path)
		}
	}
}

func removeFiles(t *testing.T, paths []string) {
	t.Helper()
	for _, Path := range paths {
		if err := os.Remove(Path); err != nil {
			if os.IsNotExist(err) {
				continue
			}
			t.Errorf("Cannot remove files: %v, error = %v", paths, err)
		}
	}
}
