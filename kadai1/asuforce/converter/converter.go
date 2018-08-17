package converter

import (
	"fmt"
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

func Convert(path string) {
	var c converter
	c.path = path
	c.ext = filepath.Ext(path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("[Error]", err)
		return
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("[Error]", err)
		return
	}

	outputFile, err := os.Create(c.getFileName())
	if err != nil {
		fmt.Println("[Error]", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	if err != nil {
		fmt.Println("[Error]", err)
	}
}

func (c *converter) getFileName() string {
	rep := regexp.MustCompile(c.ext + "$")
	dest := filepath.Base(rep.ReplaceAllString(c.path, ""))
	return dest + ".png"
}
