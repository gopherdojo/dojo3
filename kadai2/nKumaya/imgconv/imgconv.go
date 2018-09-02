package imgconv

import (
	"os"

	"github.com/dojo3/kadai2/nKumaya/format"
)

type Converter struct {
	baseFile, newFile string
	Format            format.Format
}

func NewConverter(baseFile, newFile string, format format.Format) *Converter {
	c := Converter{}
	c.baseFile = baseFile
	c.newFile = newFile
	c.Format = format
	return &c
}

func (c *Converter) Convert() error {
	baseFile, err := os.Open(c.baseFile)
	if err != nil {
		return err
	}
	defer baseFile.Close()
	img, err := c.Format.Decoder.Decode(baseFile)
	if err != nil {
		return err
	}
	newFile, err := os.Create(c.newFile)
	if err != nil {
		return err
	}
	defer newFile.Close()
	err = c.Format.Encoder.Encode(newFile, img)
	return nil
}
