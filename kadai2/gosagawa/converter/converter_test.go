package converter

import (
	"testing"
)

func TestNewConverter(t *testing.T) {

	c := NewConverter("jpg", "png", "../test", false)
	c.ConvertImage()
}
