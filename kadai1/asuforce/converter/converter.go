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
	path string
	name string
	ext  string
}

// Convert image functon
func Convert(path string) error {
	var c converter

	c.path = path
	err := filepath.Walk(c.path, c.crawlFile)
	if err != nil {
		return err
	}

	for _, v := range c.files {
		file, err := os.Open(v.path)
		if err != nil {
			return err
		}
		defer file.Close()

		img, err := jpeg.Decode(file)
		if err != nil {
			return err
		}

		outputFile, err := os.Create(v.getFileName())
		if err != nil {
			return err
		}
		defer outputFile.Close()

		err = png.Encode(outputFile, img)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *converter) crawlFile(path string, info os.FileInfo, err error) error {

	if filepath.Ext(path) == ".jpg" {
		if !info.IsDir() {
			var i = newImage(path)
			c.files = append(c.files, i)
		}
	}
	return nil
}

func newImage(path string) image {
	ext := filepath.Ext(path)
	rep := regexp.MustCompile(ext + "$")
	name := filepath.Base(rep.ReplaceAllString(path, ""))

	return image{
		path: path,
		name: name,
		ext:  ext,
	}
}

func (i *image) getFileName() string {
	return i.name + ".png"
}
