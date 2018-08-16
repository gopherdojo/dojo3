package main

import (
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"
)

var testFixtures = []struct {
	description string
	params      Parameter
	outputFiles []string
}{
	{
		"Test convert jpg to png",
		Parameter{
			inputType:  "jpg",
			outputType: "png",
			path:       []string{"testdata/jpg"},
		},
		[]string{"testdata/jpg/input1/gopher1.png", "testdata/jpg/input2/gopher2.png"},
	},
	{
		"Test convert jpg to gif",
		Parameter{
			inputType:    "jpg",
			outputType:   "gif",
			gifNumColors: 100,
			path:         []string{"testdata/jpg"},
		},
		[]string{"testdata/jpg/input1/gopher1.gif", "testdata/jpg/input2/gopher2.gif"},
	},
	{
		"Test convert png to gif",
		Parameter{
			inputType:    "png",
			outputType:   "gif",
			gifNumColors: 100,
			path:         []string{"testdata/png"},
		},
		[]string{"testdata/png/input1/gopher1.gif", "testdata/png/input2/gopher2.gif"},
	},
	{
		"Test convert png to jpg",
		Parameter{
			inputType:  "png",
			outputType: "jpg",
			jpgQuality: 100,
			path:       []string{"testdata/png"},
		},
		[]string{"testdata/png/input1/gopher1.jpg", "testdata/png/input2/gopher2.jpg"},
	},
	{
		"Test convert gif to jpg",
		Parameter{
			inputType:  "gif",
			outputType: "jpg",
			jpgQuality: 100,
			path:       []string{"testdata/gif"},
		},
		[]string{"testdata/gif/input1/gopher1.jpg", "testdata/gif/input2/gopher2.jpg"},
	},
	{
		"Test convert gif to png",
		Parameter{
			inputType:  "gif",
			outputType: "png",
			path:       []string{"testdata/gif"},
		},
		[]string{"testdata/gif/input1/gopher1.png", "testdata/gif/input2/gopher2.png"},
	},
}

func TestRun(t *testing.T) {
	for _, tt := range testFixtures {
		removeFiles(tt.outputFiles)
		if err := run(tt.params); err != nil {
			t.Errorf("Error: %s", err)
		}

		for _, file := range tt.outputFiles {
			// check output file exists
			if _, err := os.Stat(file); os.IsNotExist(err) {
				t.Errorf("Expected output file %s is not exist", file)
			}
			// check output file is in right format
			if !isDecodable(file, tt.params.outputType) {
				t.Errorf("Output file %s is in wrong format", file)
			}
		}

		removeFiles(tt.outputFiles)
	}
}

func removeFiles(paths []string) {
	for _, path := range paths {
		if err := os.Remove(path); err != nil {
			if os.IsNotExist(err) {
				continue
			}
			log.Fatal(err)
		}
	}
}

func isDecodable(path string, fileType string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	switch fileType {
	case "jpg":
		_, err = jpeg.Decode(file)
	case "png":
		_, err = png.Decode(file)
	case "gif":
		_, err = gif.Decode(file)
	}

	if err != nil {
		return false
	}
	return true
}
