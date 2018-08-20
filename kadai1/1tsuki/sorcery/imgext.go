package sorcery

import (
	"strings"
	"fmt"
)

// imgExt is list of image file extensions supported by sorcery
type imgExt int

const (
	// Jpeg image format
	Jpeg imgExt = iota
	// Png image format
	Png
	// Gif image format
	Gif
	end
)

// ImgExt creates instance of imgExt from string.
// Returns error when unsupported extension was specified.
func ImgExt(ext string) (imgExt, error) {
	switch trimExt(ext) {
	case "jpeg":
		return Jpeg, nil
	case "jpg":
		return Jpeg, nil
	case "png":
		return Png, nil
	case "gif":
		return Gif, nil
	default:
		return 0, fmt.Errorf("unsupported extension specified: %v", ext)
	}
}

// String method is implemented to enable Stringer interface
func (e imgExt) String() string {
	switch e {
	case Jpeg:
		return "jpg"
	case Png:
		return "png"
	case Gif:
		return "gif"
	default:
		return "Unknown"
	}
}

func (e imgExt) isValid() bool {
	return e < end
}

func trimExt(ext string) string {
	return strings.TrimLeft(strings.ToLower(ext), ".")
}
