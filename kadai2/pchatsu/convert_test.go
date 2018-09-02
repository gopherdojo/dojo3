package imgconv_test

import (
	"os"
	"reflect"
	"testing"
	"github.com/gopherdojo/dojo3/kadai2/pchatsu"
)

func TestConvert(t *testing.T) {
	type args struct {
		d   string
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"#1 mock run", args {"testdata/img","gif","jpeg"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := imgconv.Convert(tt.args.d, tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOutputPath(t *testing.T) {
	type args struct {
		path string
		dst  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
	}{
		{"#1 normal case", args{"testdata/img/foo.gif", "jpeg"}, "testdata/img/foo.jpg"},
		{"#2 normal case", args{"testdata/img/foo.jpg", "png"}, "testdata/img/foo.png"},
		{"#3 normal case", args{"testdata/img/foo.png", "gif"}, "testdata/img/foo.gif"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got  := imgconv.GetOutputPath(tt.args.path, tt.args.dst)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openOutputFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTargetFile(t *testing.T) {
	type args struct {
		info   os.FileInfo
		path   string
		format string
	}
	testDirPath := "testdata/img"
	dirInfo, err := os.Stat(testDirPath)
	if err != nil {
		t.Errorf("fail to get stat of %s", testDirPath)
		return
	}

	hiddenFilePath := "testdata/img/.hidden.golden"
	hiddenFileInfo, err := os.Stat(hiddenFilePath)
	if err != nil {
		t.Errorf("fail to get stat of %s", hiddenFilePath)
		return
	}

	normalFilePath := "testdata/img/gif.golden"
	normalFileInfo, err := os.Stat(normalFilePath)
	if err != nil {
		t.Errorf("failt to get stat of %s", normalFilePath)
		return
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"#1 a directory is not target file", args{dirInfo, testDirPath, ""}, false},
		{"#2 hidden file is not target file", args{hiddenFileInfo, hiddenFilePath, ""}, false},
		{"#3 hidden file is not target file", args{normalFileInfo, "test.png", "png"}, true},
		{"#4 hidden file is not target file", args{normalFileInfo, "test.jpg", "jpeg"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imgconv.IsTargetFile(tt.args.info, tt.args.path, tt.args.format); got != tt.want {
				t.Errorf("isTargetFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
