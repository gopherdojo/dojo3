package encoder

import (
	"image"
	"image/jpeg"
	"io"
)

type Jpg struct{}

func (j *Jpg) Run(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}

func (j *Jpg) Ext() string {
	return "jpg"
}
