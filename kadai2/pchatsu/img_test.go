package imgconv_test

import (
	"bytes"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/pchatsu"
)

func TestNewImg(t *testing.T) {
	type args struct {
		path   string
		format string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"#1 normal input for gif", args{"testdata/img/gif.golden", "gif"}, false},
		{"#2 normal input for jpeg", args{"testdata/img/jpeg.golden", "jpeg"}, false},
		{"#3 normal input png", args{"testdata/img/png.golden", "png"}, false},
		{"#4 not exist err", args{"testdata/not_exist", "jpg"}, true},
		{"#5 invalid format error", args{"testdata/img/png.golden", "jpeg"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := imgconv.NewImg(tt.args.path, tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewImg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestImg_Encode(t *testing.T) {
	type args struct {
		dst string
	}
	gif, err := imgconv.NewImg("testdata/img/gif.golden", "gif")
	if err != nil {
		t.Errorf("fail to prepare test file")
		return
	}
	jpeg, err := imgconv.NewImg("testdata/img/jpeg.golden", "jpeg")
	if err != nil {
		t.Errorf("fail to prepare test file")
		return
	}
	png, err := imgconv.NewImg("testdata/img/png.golden", "png")
	if err != nil {
		t.Errorf("fail to prepare test file")
		return
	}

	tests := []struct {
		name    string
		img     *imgconv.Img
		args    args
		wantErr bool
	}{
		{"#1 normal case for gif", gif, args{"jpeg"}, false},
		{"#2 normal case for jpeg", jpeg, args{"png"}, false},
		{"#3 normal case for png", png, args{"gif"}, false},
		{"#4 invalid dst", gif, args{"mov"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := tt.img.Encode(w, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("Img.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
