package imgconv

import (
	"reflect"
	"testing"
	"bytes"
	"os"
	"image"
)

func Test_imgConverter_Convert(t *testing.T) {
	t.Run("#1 normal case", func(t *testing.T) {
		assertConvert(t, "lena.jpg", "png")
		assertConvert(t, "lena.jpg", "gif")
		//assertConvert(t, "lena.png", "jpg")
		//assertConvert(t, "lena.png", "png")
		//assertConvert(t, "lena.gif", "png")
		//assertConvert(t, "lena.gif", "gif")
	})
}

func assertConvert(t *testing.T, filename string, destFormat string) {
	t.Helper()

	c := &imgConverter{
		DestFormat: destFormat,
	}

	w := &bytes.Buffer{}
	r, err := os.Open(testdata(filename))
	if err != nil {
		t.Error(err)
	}
	defer r.Close()

	if err := c.Convert(r, w); err != nil {
		t.Errorf("imgConverter.Convert() error = %v", err)
		return
	}

	_, format, e := image.Decode(w)
	if e != nil {
		t.Error(err)
	}
	if format != destFormat {
		t.Errorf("The format of the image produced by imgConverter.Convert() = %v, want %v", format, destFormat)
	}
}

func TestImgConverter(t *testing.T) {
	type args struct {
		opt *Option
	}
	tests := []struct {
		name string
		args args
		want *imgConverter
	}{
		{name: "#1 normal case", args: args{opt: &Option{Input: []string{"jpg"}, Output: "png"}}, want: &imgConverter{SrcFormat: []string{"jpg"}, DestFormat: "png"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ImgConverter(tt.args.opt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImgConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}
