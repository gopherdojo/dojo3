package converter

import (
	"fmt"
	"os"
	"testing"
)

func TestConverter_Convert(t *testing.T) {
	type fields struct {
		Encoder Encoder
		Decoder Decoder
	}
	type args struct {
		path string
		dest string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Jpeg to Png", fields: fields{Encoder: &Png{}, Decoder: &Jpg{}}, args: args{path: "testdata/test1.jpeg", dest: "test1.png"}, wantErr: false},
		{name: "Jpeg to Gif", fields: fields{Encoder: &Gif{}, Decoder: &Jpg{}}, args: args{path: "testdata/test1.jpeg", dest: "test1.gif"}, wantErr: false},
		{name: "Png to Jpeg", fields: fields{Encoder: &Jpg{}, Decoder: &Png{}}, args: args{path: "testdata/test2.png", dest: "test2.jpg"}, wantErr: false},
		{name: "Png to Gif", fields: fields{Encoder: &Gif{}, Decoder: &Png{}}, args: args{path: "testdata/test2.png", dest: "test2.gif"}, wantErr: false},
		{name: "Gif to Jpeg", fields: fields{Encoder: &Jpg{}, Decoder: &Gif{}}, args: args{path: "testdata/test3.gif", dest: "test3.jpg"}, wantErr: false},
		{name: "Gif to Png", fields: fields{Encoder: &Png{}, Decoder: &Gif{}}, args: args{path: "testdata/test3.gif", dest: "test3.png"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Converter{
				Encoder: tt.fields.Encoder,
				Decoder: tt.fields.Decoder,
			}
			if err := c.Convert(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Converter.Convert() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := os.Remove(tt.args.dest); err != nil {
				fmt.Println(err)
			}
		})
	}
}

func TestConverter_getFileName(t *testing.T) {
	type fields struct {
		Encoder Encoder
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{name: "Jpeg to Png", fields: fields{Encoder: &Png{}}, args: args{path: "testdata/test.jpeg"}, want: "test.png", wantErr: false},
		{name: "Jpeg to Gif", fields: fields{Encoder: &Gif{}}, args: args{path: "testdata/test.jpeg"}, want: "test.gif", wantErr: false},
		{name: "Png to Jpeg", fields: fields{Encoder: &Jpg{}}, args: args{path: "testdata/test.png"}, want: "test.jpg", wantErr: false},
		{name: "Png to Gif", fields: fields{Encoder: &Gif{}}, args: args{path: "testdata/test.png"}, want: "test.gif", wantErr: false},
		{name: "Gif to Jpeg", fields: fields{Encoder: &Jpg{}}, args: args{path: "testdata/test.gif"}, want: "test.jpg", wantErr: false},
		{name: "Gif to Png", fields: fields{Encoder: &Png{}}, args: args{path: "testdata/test.gif"}, want: "test.png", wantErr: false},
		{name: "Empty path", fields: fields{Encoder: &Jpg{}}, args: args{path: ""}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Converter{
				Encoder: tt.fields.Encoder,
			}
			got, err := c.getFileName(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Converter.getFileName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Converter.getFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
