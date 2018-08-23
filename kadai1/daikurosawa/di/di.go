// Package di is dependency injection to image convert cil tool.
package di

import (
	"image"
	"io"
)

// Converter is interface that has Decode and Encode function.
type Converter interface {
	Decode(r io.Reader) (image.Image, error)
	Encode(w io.Writer, m image.Image) error
}

var Converts = map[string]Converter{}

// Register sets command to converts.
func Register(key string, convert Converter) {
	Converts[key] = convert
}
