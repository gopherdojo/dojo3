package convert

import (
	"testing"
)

func TestConvertFormat(t *testing.T) {
	patterns := []struct {
		name string
		src string
		outputExt string
		dir string
	}{
		{
      name: "jpg to png",
      src: "testdata/1px.jpg",
      outputExt: "png",
    },
		{
      name: "png to jpg",
      src: "testdata/1px.png",
      outputExt: "jpg",
    },
	}
  for _, p := range patterns {
		t.Run(p.name, func(t *testing.T) {
			c := &Converter{
				Src:       p.src,
				OutputExt: p.outputExt,
			}
			err := c.Convert()
			if err != nil {
				t.Errorf("p.name: %v", err)
			}
		})
	}
}

func TestConvertFail(t *testing.T) {
	patterns := []struct {
		name string
		src string
		outputExt string
		dir string
	}{
		{
      name: "file not found",
      src: "testdata/xxx.jpg",
      outputExt: "png",
    },
		{
      name: "unsupported file",
      src: "testdata/dummy.txt",
      outputExt: "png",
    },
	}
  for _, p := range patterns {
		t.Run(p.name, func(t *testing.T) {
			c := &Converter{
				Src:       p.src,
				OutputExt: p.outputExt,
			}
			err := c.Convert()
			if err == nil {
				t.Errorf("p.name: %v", err)
			}
		})
	}
}
