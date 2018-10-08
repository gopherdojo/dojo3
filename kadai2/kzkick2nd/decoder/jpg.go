package decoder

import (
	"image"
	"io"
)

type Jpg struct{}

func (j *Jpg) Ext() string {
	return "jpg"
}

func (j *Jpg) Run(r io.Reader) (image.Image, error) {
	m, _, err := image.Decode(r)
	return m, err
}
