package convert

import (
	"image"
	"io"
)

// Converter is interface that has Decode and Encode function.
type Converter interface {
	Decode(r io.Reader) (image.Image, error)
	Encode(w io.Writer, m image.Image) error
}

var converts = map[string]Converter{}

// Register sets command to commands
func Register(key string, convert Converter) {
	converts[key] = convert
}
