package imgconv

import (
	"reflect"
	"testing"
	"bytes"
	"os"
)

func Test_imgConverter_Convert(t *testing.T) {
	type fields struct {
		DestFormat string
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "#1 normal case", fields: fields{DestFormat: "png"}, args: args{filename: "lena.jpg"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &imgConverter{
				DestFormat: tt.fields.DestFormat,
			}
			w := &bytes.Buffer{}
			r, err := os.Open(testdata(tt.args.filename))
			if err != nil {
				t.Error(err)
			}
			defer r.Close()
			if err := c.Convert(r, w); (err != nil) != tt.wantErr {
				t.Errorf("imgConverter.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
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
