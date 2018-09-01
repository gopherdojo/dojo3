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
	testFilePath["jpg"] = "testdata/gopher.jpg"
	testFilePath["png"] = "testdata/gopher.png"
	testFilePath["gif"] = "testdata/gopher.gif"
}

var convertTests = []struct {
	name string
	*option.Option
}{
	{
		"convert from jpg to png",
		&option.Option{DirName: "dummy", FromExtension: "jpg", ToExtension: "png"},
	},
	{
		"convert from jpg to gif",
		&option.Option{DirName: "dummy", FromExtension: "jpg", ToExtension: "gif"},
	},
	{
		"convert from png to jpg",
		&option.Option{DirName: "dummy", FromExtension: "png", ToExtension: "jpg"},
	},
	{
		"convert from png to gif",
		&option.Option{DirName: "dummy", FromExtension: "png", ToExtension: "gif"},
	},
	{
		"convert from gif to jpg",
		&option.Option{DirName: "dummy", FromExtension: "gif", ToExtension: "jpg"},
	},
	{
		"convert from gif to png",
		&option.Option{DirName: "dummy", FromExtension: "gif", ToExtension: "png"},
	},
}

func TestConvert_Convert(t *testing.T) {
	for _, conv := range convertTests {
		conv := conv
		t.Run(conv.name, func(t *testing.T) {
			sut := convert.NewConvert(conv.Option)
			path := testFilePath[conv.FromExtension]
			if err := sut.Convert(path); err != nil {
				t.Fatal("failed convert", err)
			}
		})
	}
}
