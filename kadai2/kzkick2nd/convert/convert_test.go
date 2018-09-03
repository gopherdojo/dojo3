package convert

import (
	"testing"
)

func TestConvertFormat(t *testing.T) {
	t.Run("jpg to png", func(t *testing.T) {
		c := &Converter{
			Src:       "testfiles/1px.jpg",
			OutputExt: "png",
		}
		err := c.Convert()
		if err != nil {
			t.Errorf("jpg to png: %v", err)
			return
		}
	})

	t.Run("png to jpg", func(t *testing.T) {
		c := &Converter{
			Src:       "testfiles/1px.png",
			OutputExt: "jpg",
		}
		err := c.Convert()
		if err != nil {
			t.Errorf("png to jpg: %v", err)
			return
		}
	})
}

func TestNotFound(t *testing.T) {
	t.Run("file not found", func(t *testing.T) {
		c := &Converter{
			Src:       "testfiles/xxx.jpg",
			OutputExt: "png",
		}
		err := c.Convert()
		if err == nil {
			t.Errorf("file not found: %v", err)
			return
		}
	})
}

func TestUnsupported(t *testing.T) {
	t.Run("unsupported file", func(t *testing.T) {
		c := &Converter{
			Src:       "testfiles/dummy.txt",
			OutputExt: "png",
		}
		err := c.Convert()
		if err == nil {
			t.Errorf("unsupported file: %v", err)
			return
		}
	})
}
