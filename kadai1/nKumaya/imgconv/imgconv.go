package imgconv

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type Converter struct {
	baseFile, newFile string
}

func NewConverter(baseFile, newFile string) *Converter {
	c := Converter{}
	c.baseFile = baseFile
	c.newFile = newFile
	return &c
}

func (c *Converter) Convert() error {
	baseFile, err := os.Open(c.baseFile)
	defer baseFile.Close()
	img, _, err := image.Decode(baseFile)
	newFile, err := os.Create(c.newFile)
	switch filepath.Ext(c.newFile) {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(newFile, img, nil)
	case ".png":
		err = png.Encode(newFile, img)
	case ".gif":
		err = gif.Encode(newFile, img, nil)
	}

	if err != nil {
		return err
	}
	return nil
}
