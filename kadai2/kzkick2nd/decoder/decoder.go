package decoder

import (
	"image"
	"io"
)

type Decoder interface {
	Ext() string
	Run(io.Reader) (image.Image, error)
}
