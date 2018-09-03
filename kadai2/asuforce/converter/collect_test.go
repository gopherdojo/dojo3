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
		{name: "Select jpg file", fields: fields{Paths: []string{}, FromExt: "jpg"}, args: args{path: "testdata/"}, wantErr: false},
		{name: "Select png file", fields: fields{Paths: []string{}, FromExt: "png"}, args: args{path: "testdata/"}, wantErr: false},
		{name: "Select gif file", fields: fields{Paths: []string{}, FromExt: "gif"}, args: args{path: "testdata/"}, wantErr: false},
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

func TestCollect_appendFiles(t *testing.T) {
	type fields struct {
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
		{name: "Input jpeg type", fields: fields{FromExt: "jpeg"}, args: args{path: ".jpeg"}, wantErr: false},
		{name: "Input jpg type", fields: fields{FromExt: "jpg"}, args: args{path: ".jpg"}, wantErr: false},
		{name: "Input png type", fields: fields{FromExt: "png"}, args: args{path: ".png"}, wantErr: false},
		{name: "Input gif type", fields: fields{FromExt: "gif"}, args: args{path: ".gif"}, wantErr: false},
		{name: "Enpty path", fields: fields{FromExt: ""}, args: args{path: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Collect{
				FromExt: tt.fields.FromExt,
			}
			if err := c.appendFiles(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Collect.appendFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
