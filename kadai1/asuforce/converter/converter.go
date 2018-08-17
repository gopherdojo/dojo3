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
}

func Convert(path string) {
	var c converter
	c.path = path

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
	ext := filepath.Ext(c.path)
	rep := regexp.MustCompile(ext + "$")
	dest := filepath.Base(rep.ReplaceAllString(c.path, ""))
	return dest + ".png"
}
