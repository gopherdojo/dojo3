package convert_test

import (
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert"
	_ "github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/gif"
	_ "github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/jpg"
	_ "github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/png"
	"github.com/gopherdojo/dojo3/kadai2/daikurosawa/option"
)

var testFilePath = map[string]string{}

func init() {
	testFilePath["jpg"] = "./../testdata/gopher.jpg"
	testFilePath["png"] = "./../testdata/gopher.png"
	testFilePath["gif"] = "./../testdata/gopher.gif"
}

var convertTests = []struct {
	name string
	*option.Option
}{
	{
		"convert from jpg to png",
		&option.Option{FromExtension: "jpg", ToExtension: "png"},
	},
	{
		"convert from jpg to gif",
		&option.Option{FromExtension: "jpg", ToExtension: "gif"},
	},
	{
		"convert from png to jpg",
		&option.Option{FromExtension: "png", ToExtension: "jpg"},
	},
	{
		"convert from png to gif",
		&option.Option{FromExtension: "png", ToExtension: "gif"},
	},
	{
		"convert from gif to jpg",
		&option.Option{FromExtension: "gif", ToExtension: "jpg"},
	},
	{
		"convert from gif to png",
		&option.Option{FromExtension: "gif", ToExtension: "png"},
	},
}

func TestConvert_Convert(t *testing.T) {
	for _, tt := range convertTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			sut := convert.NewConvert(tt.Option)
			path := testFilePath[tt.FromExtension]
			if err := sut.Convert(path); err != nil {
				t.Fatal("failed convert", err)
			}
		})
	}
}
