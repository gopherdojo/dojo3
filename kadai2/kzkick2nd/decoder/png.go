package decoder

import (
	"image"
	"io"
)

type Png struct{}

func (j *Png) Ext() string {
	return "png"
}

func (j *Png) Run(r io.Reader) (image.Image, error) {
	m, _, err := image.Decode(r)
	return m, err
}
