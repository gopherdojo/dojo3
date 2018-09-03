// Package converter provides image convert between jpg and png.
package convert

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

// Converter has path of src and output extension (jpg|png).
type Converter struct {
	Src       string
	OutputExt string
}

// FIXME 処理計画を並べて実行的な事できない？

// Convert image to same dir/name with different extension.
func (c *Converter) Convert() error {
	file, err := os.Open(c.Src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	destPath := c.Src[:len(c.Src)-len(filepath.Ext(c.Src))] + "." + c.OutputExt
	fd, err := os.Create(destPath)
	if err != nil {
		log.Fatal(err)
	}

	switch c.OutputExt {
	case "jpg":
		err = jpeg.Encode(fd, img, nil)
	case "png":
		err = png.Encode(fd, img)
	}
	if err != nil {
		log.Fatal(err)
	}

	err = fd.Close()
	if err != nil {
		log.Fatal(err)
	}

	return err
}
