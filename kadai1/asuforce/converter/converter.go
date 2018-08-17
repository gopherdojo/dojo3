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

func Convert(path string) error {
	var (
		c converter
		i image
	)

	c.path = path
	i.new(c.path)
	c.files = append(c.files, i)

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

func (i *image) new(path string) {
	i.ext = filepath.Ext(path)
	rep := regexp.MustCompile(i.ext + "$")
	i.name = filepath.Base(rep.ReplaceAllString(path, ""))
}

func (i *image) getFileName() string {
	return i.name + ".png"
}
