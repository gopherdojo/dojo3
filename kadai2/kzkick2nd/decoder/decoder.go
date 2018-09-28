/*
Packege decoder provides decoder interface and type with formats.
*/

package decoder

import (
	"image"
	"io"
)

// Decoder represents image decoder.
type Decoder interface {
	Ext() string
	Run(io.Reader) (image.Image, error)
}
