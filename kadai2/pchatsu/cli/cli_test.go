package cli_test

import (
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/pchatsu/cli"
)

func TestRun(t *testing.T) {
	type args struct {
		path string
		src  string
		dst  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"normal input jpeg", args{"testdata/test", "jpeg", "png"}, false},
		{"normal input png", args{"testdata/test", "png", "gif"}, false},
		{"normal input gif", args{"testdata/test", "gif", "png"}, false},
		{"not exist error", args{"testdata/path_to_not_exists", "jpeg", "png"}, true},
		{"invalid ext error #1", args{"testdata/test", "jpg", "png"}, true},
		{"invalid ext error #2", args{"testdata/test", "png", "jpg"}, true},
		{"the same ext error #1", args{"testdata/test", "jpeg", "jpeg"}, true},
		{"the same ext error #2", args{"testdata/test", "jpg", "jpg"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cli.Run(tt.args.path, tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
