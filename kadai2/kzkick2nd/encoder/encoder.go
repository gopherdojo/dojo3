package encoder

import (
	"image"
	"io"
)

type Encoder interface {
	Ext() string
	Run(io.Writer, image.Image) error
}
