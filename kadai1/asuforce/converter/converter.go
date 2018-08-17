package converter

import (
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
)

type converter struct {
	path  string
	files []image
}

type image struct {
	name string
	ext  string
}

// Convert image functon
func Convert(path string) error {
	var c converter

	c.path = path
	var i = newImage(c.path)

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(i.getFileName())
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	if err != nil {
		return err
	}

	return nil
}

func newImage(path string) image {
	ext := filepath.Ext(path)
	rep := regexp.MustCompile(ext + "$")
	name := filepath.Base(rep.ReplaceAllString(path, ""))

	return image{
		name: name,
		ext:  ext,
	}
}

func (i *image) getFileName() string {
	return i.name + ".png"
}
