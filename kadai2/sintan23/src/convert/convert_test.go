/*
ConvertImage Test
*/
package convert

import (
	"os"
	"testing"
)

type Case = struct {
	name  string
	src   string
	toExt string
}

func TestConvertAll(t *testing.T) {
	var convertTests = []struct {
		caseName string
		toExt    string
	}{
		{caseName: "Test Convert Jpg to Png 1", toExt: "png"},
		{caseName: "Test Convert Jpg to Png 2", toExt: "png"},
	}
	for _, c := range convertTests {
		t.Run(
			c.caseName, func(t *testing.T) {
				ConvertAll("png")
			})
	}
}

func TestConvert(t *testing.T) {
	var convertTests = []Case{
		{name: "Test Convert Jpg to Png", src: "../../_data/jpg1.jpg", toExt: "png"},
	}

	for _, tt := range convertTests {
		testConvert(t, tt)
	}
}

func testConvert(t *testing.T, c Case) {
	t.Helper()

	r, err := os.Open(c.src)
	if err != nil {
		t.Fatalf("Open error = %v", err)
	}

	ct := Converter{
		c.src,
		c.toExt,
		r,
	}

	t.Run(c.name, func(t *testing.T) {
		err := ct.Convert(ct.toExt)
		if err != nil {
			t.Errorf("Convert error = %v", err)
		}
	})
}

func TestConvertJpegToPng(t *testing.T) {
	var convertTests = []struct {
		caseName string
		src      string
		toExt    string
	}{
		{caseName: "Test ConvertJpegToPng", src: "../../_data/jpg1.jpg", toExt: "png"},
	}

	for _, c := range convertTests {
		t.Run(
			c.caseName, func(t *testing.T) {
				r, err := Open(c.src)
				logErr(err)
				ct := Converter{
					c.src,
					c.toExt,
					r,
				}
				ct.ConvertJpegToPng(c.toExt)
			})
	}
}

func TestLogErr(t *testing.T) {
	var convertTests = []struct {
		caseName string
		src      string
	}{
		{caseName: "Test Open+LogErr", src: "../../_data/jpg1.jpg"},
	}

	for _, c := range convertTests {
		t.Run(
			c.caseName, func(t *testing.T) {
				_, err := Open(c.src)
				logErr(err)
			})
	}
}
