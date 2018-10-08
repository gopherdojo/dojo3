/*
Packege encoder provides encoder interface and type with formats.
*/
package encoder

import (
	"image"
	"io"
)

// Encoder represents image encoder.
type Encoder interface {
	Ext() string
	Run(io.Writer, image.Image) error
}
