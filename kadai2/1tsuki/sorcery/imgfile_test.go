package sorcery

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

const (
	srcDir  string = "./testdata"
	srcJpeg string = "jpegfile.jpg"
	srcPng  string = "pngfile.png"
	srcGif  string = "giffile.gif"
	srcTiff string = "tifffile.tiff" // note that tiff is currently not supported
)

type TestFiles struct {
	Dir  string
	Jpeg string
	Png  string
	Gif  string
	Tiff string
}

func initTestDataSet(t *testing.T) *TestFiles {
	t.Helper()

	tempDir, err := ioutil.TempDir("", "sorcery-test")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	testFiles := &TestFiles{
		Dir:  tempDir,
		Jpeg: tempDir + "/" + srcJpeg,
		Png:  tempDir + "/" + srcPng,
		Gif:  tempDir + "/" + srcGif,
		Tiff: tempDir + "/" + srcTiff,
	}

	// FIXME : write method to copy items recursively
	if err := copyFile(srcDir+"/"+srcJpeg, testFiles.Jpeg, t); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := copyFile(srcDir+"/"+srcPng, testFiles.Png, t); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := copyFile(srcDir+"/"+srcGif, testFiles.Gif, t); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := copyFile(srcDir+"/"+srcTiff, testFiles.Tiff, t); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	return testFiles
}

func copyFile(src, dst string, t *testing.T) {
	t.Helper()

	from, err := os.Open(src)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer from.Close()

	to, err := os.Create(dst)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func Test_imgFile_convertTo(t *testing.T) {
	testFiles := initTestDataSet(t)
	defer os.RemoveAll(testFiles.Dir)

	type fields struct {
		Path string
	}
	type args struct {
		to imgExt
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// FIXME want値が雑になってしまった
		{"jpg to jpg", fields{testFiles.Jpeg}, args{Jpeg}, testFiles.Dir + "/" + "jpegfile.jpg", false},
		{"jpg to png", fields{testFiles.Jpeg}, args{Png}, testFiles.Dir + "/" + "jpegfile.png", false},
		{"jpg to gif", fields{testFiles.Jpeg}, args{Gif}, testFiles.Dir + "/" + "jpegfile.gif", false},
		{"jpg to invalid", fields{testFiles.Jpeg}, args{end}, "", true},
		{"png to jpg", fields{testFiles.Png}, args{Jpeg}, testFiles.Dir + "/" + "pngfile.jpg", false},
		{"png to png", fields{testFiles.Png}, args{Png}, testFiles.Dir + "/" + "pngfile.png", false},
		{"png to gif", fields{testFiles.Png}, args{Gif}, testFiles.Dir + "/" + "pngfile.gif", false},
		{"png to invalid", fields{testFiles.Png}, args{end}, "", true},
		{"gif to jpg", fields{testFiles.Gif}, args{Jpeg}, testFiles.Dir + "/" + "giffile.jpg", false},
		{"gif to png", fields{testFiles.Gif}, args{Png}, testFiles.Dir + "/" + "giffile.png", false},
		{"gif to gif", fields{testFiles.Gif}, args{Gif}, testFiles.Dir + "/" + "giffile.gif", false},
		{"gif to invalid", fields{testFiles.Gif}, args{end}, "", true},
		{"unsupported to jpg", fields{testFiles.Tiff}, args{Jpeg}, "", true},
		{"unsupported to png", fields{testFiles.Tiff}, args{Png}, "", true},
		{"unsupported to gif", fields{testFiles.Tiff}, args{Gif}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &imgFile{
				Path: tt.fields.Path,
			}
			got, err := c.convertTo(tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("imgFile.convertTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("imgFile.convertTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imgFile_extString(t *testing.T) {
	type fields struct {
		Path string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"normal", fields{"/path/with/extendsion.ext"}, "ext"},
		{"has multiple extensions", fields{"/path/with/extendsion.ext1.ext2"}, "ext2"},
		{"without extension", fields{"/path/with/extendsion"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &imgFile{
				Path: tt.fields.Path,
			}
			if got := c.extString(); got != tt.want {
				t.Errorf("imgFile.extString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imgFile_out(t *testing.T) {
	type fields struct {
		Path string
	}
	type args struct {
		to imgExt
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"normal", fields{"/some/jpeg/file.jpg"}, args{Gif}, "/some/jpeg/file.gif"},
		{"different ext length", fields{"/some/jpeg/file.jpeg"}, args{Png}, "/some/jpeg/file.png"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &imgFile{
				Path: tt.fields.Path,
			}
			if got := c.out(tt.args.to); got != tt.want {
				t.Errorf("imgFile.out() = %v, want %v", got, tt.want)
			}
		})
	}
}
