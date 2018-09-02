package converter_test

import (
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/shuntaka9576/converter"
	_ "github.com/gopherdojo/dojo3/kadai2/shuntaka9576/imagetypes/gif"
	_ "github.com/gopherdojo/dojo3/kadai2/shuntaka9576/imagetypes/jpg"
	_ "github.com/gopherdojo/dojo3/kadai2/shuntaka9576/imagetypes/png"
)

func TestGetConverter(t *testing.T) {
	type input struct {
		from, to string
	}
	tests := []struct {
		pattern string
		name    string
		input   input
	}{
		{"normal", "png to jpg", input{"png", "jpg"}},
		{"normal", "png to jpeg", input{"png", "jpeg"}},
		{"normal", "png to gif", input{"png", "gif"}},
		{"normal", "gif to jpg", input{"gif", "jpg"}},
		{"normal", "gif to jpeg", input{"gif", "jpeg"}},
		{"normal", "gif to png", input{"gif", "png"}},
		{"normal", "jpg to png", input{"jpg", "png"}},
		{"normal", "jpeg to png", input{"jpeg", "png"}},
		{"normal", "jpg to gif", input{"jpg", "gif"}},
		{"normal", "jpeg to gif", input{"jpeg", "gif"}},
		{"non-normal", "fail from argument", input{"jp", "gif"}},
		{"non-normal", "fail to argument", input{"jpg", "gi"}},
		{"non-normal", "fail all argument", input{"jp", "gi"}},
	}
	for _, tt := range tests {
		if tt.pattern == "normal" {
			t.Run(tt.name, func(t *testing.T) {
				_, err := converter.GetConverter(tt.input.from, tt.input.to)
				if err != nil {
					t.Errorf("but got %v", err.Error())
				}
			})
		}
		if tt.pattern == "non-normal" {
			t.Run(tt.name, func(t *testing.T) {
				_, err := converter.GetConverter(tt.input.from, tt.input.to)
				if err == nil {
					t.Error("Test Fail")
				}
			})
		}
	}
}

func TestConverter_Convert(t *testing.T) {
	t.Helper()
	type input struct {
		from, to, input, output string
	}
	tests := []struct {
		pattern string
		name    string
		input   input
		want    string
	}{
		{"normal", "convert to jpg from gif", input{"gif", "jpg", "../testdata/1/sample1.gif", "test.jpg"}, ""},
		{"normal", "input path fail", input{"gif", "jpg", "../testdata/1/", "./test.jpg"}, ""},
		{"non-normal", "not found ext file", input{"gif", "jpg", "../testdata/1/s.gif", "./test.jpg"}, ""},
		{"non-normal", "convert to jpg from gif", input{"gif", "jpg", "../testdata/1/sample1.gif", "/te"}, ""},
	}
	for _, tt := range tests {
		if tt.pattern == "normal" {
			t.Run(tt.name, func(t *testing.T) {
				con, _ := converter.GetConverter(tt.input.from, tt.input.to)
				_, err := con.Convert(tt.input.input, tt.input.output)
				if err != nil {
					t.Errorf("but got %v", err.Error())
				}
			})
		}
		if tt.pattern == "non-normal" {
			t.Run(tt.name, func(t *testing.T) {
				con, _ := converter.GetConverter(tt.input.from, tt.input.to)
				_, err := con.Convert(tt.input.input, tt.input.output)
				if err == nil {
					t.Error("Test Fail")
				}
			})
		}
	}
}
