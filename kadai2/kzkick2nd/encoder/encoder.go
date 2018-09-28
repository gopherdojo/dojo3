package encoder

import (
	"image"
	"io"
)

type Encoder interface {
	Run(io.Writer, image.Image) error
	Ext() string
}
