package converter

import (
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
)

type converter struct {
	path string
	ext  string
}

func Convert(path string) error {
	var c converter
	c.path = path
	c.ext = filepath.Ext(path)

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(c.getFileName())
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

func (c *converter) getFileName() string {
	rep := regexp.MustCompile(c.ext + "$")
	dest := filepath.Base(rep.ReplaceAllString(c.path, ""))
	return dest + ".png"
}
