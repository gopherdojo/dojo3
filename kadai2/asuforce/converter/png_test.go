package converter

import (
	"bytes"
	"image"
	"io"
	"reflect"
	"testing"
)

func TestPng_Encode(t *testing.T) {
	type args struct {
		img image.Image
	}
	tests := []struct {
		name     string
		p        *Png
		args     args
		wantFile string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Png{}
			file := &bytes.Buffer{}
			if err := p.Encode(file, tt.args.img); (err != nil) != tt.wantErr {
				t.Errorf("Png.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFile := file.String(); gotFile != tt.wantFile {
				t.Errorf("Png.Encode() = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func TestPng_Decode(t *testing.T) {
	type args struct {
		file io.Reader
	}
	tests := []struct {
		name    string
		p       *Png
		args    args
		want    image.Image
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Png{}
			got, err := p.Decode(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Png.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Png.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPng_GetExt(t *testing.T) {
	tests := []struct {
		name string
		p    *Png
		want string
	}{
		{name: "success", want: "png"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Png{}
			if got := p.GetExt(); got != tt.want {
				t.Errorf("Png.GetExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
