package converter

import (
	"testing"
)

func TestCollect_CollectPath(t *testing.T) {
	type fields struct {
		Paths   []string
		FromExt string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Select jpg file", fields: fields{Paths: []string{}, FromExt: "jpg"}, args: args{path: "../testdata/"}, wantErr: false},
		{name: "Select png file", fields: fields{Paths: []string{}, FromExt: "png"}, args: args{path: "../testdata/"}, wantErr: false},
		{name: "Select gif file", fields: fields{Paths: []string{}, FromExt: "gif"}, args: args{path: "../testdata/"}, wantErr: false},
		{name: "Input empty path", fields: fields{Paths: []string{}, FromExt: "jpg"}, args: args{path: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Collect{
				Paths:   tt.fields.Paths,
				FromExt: tt.fields.FromExt,
			}
			if err := c.CollectPath(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Collect.CollectPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkExtension(t *testing.T) {
	type args struct {
		ext string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Input jpeg type", args: args{ext: ".jpeg"}, want: ".jpg", wantErr: false},
		{name: "Input jpg type", args: args{ext: ".jpg"}, want: ".jpg", wantErr: false},
		{name: "Input png type", args: args{ext: ".png"}, want: ".png", wantErr: false},
		{name: "Input gif type", args: args{ext: ".gif"}, want: ".gif", wantErr: false},
		{name: "Enpty path", args: args{ext: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkExtension(tt.args.ext)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkExtension() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
