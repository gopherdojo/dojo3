package encoder

import (
	"image"
	"image/png"
	"io"
)

type Png struct{}

func (j *Png) Ext() string {
	return "png"
}

func (j *Png) Run(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}
