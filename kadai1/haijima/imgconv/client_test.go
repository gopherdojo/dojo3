package imgconv

import (
	"os"
	"testing"
	"path"
	"bytes"
	"time"
	"errors"
	"io"
)

func TestCli_Run(t *testing.T) {
	type fields struct {
		Dir string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "#1 normal case", fields: fields{Dir: "testdata"}, want: SuccessExitCode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cli{
				Dir: tt.fields.Dir,
			}
			if got := c.Run(); got != tt.want {
				t.Errorf("Cli.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCli_execFile(t *testing.T) {
	type fields struct {
		Option Option
	}
	type args struct {
		path string
		info os.FileInfo
		err  error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		afterFn func()
	}{
		{name: "#1 normal case", fields: fields{Option: Option{Input: []string{"jpg"}, Output: "png"}}, args: args{path: "lena.jpg", info: fakeFileInfo("lena.jpg")}, wantErr: false, afterFn: func() {
			os.Remove(testdata("lena.png"))
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := &bytes.Buffer{}
			c := &Cli{
				Out:    out,
				Err:    err,
				Conv:   &DoNothingConverter{},
				Option: tt.fields.Option,
			}
			if err := c.execFile(testdata(tt.args.path), tt.args.info, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Cli.execFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			defer tt.afterFn()
		})
	}
}

type DoNothingConverter struct {
}

func (c *DoNothingConverter) Convert(r io.Reader, w io.Writer) error {
	return nil
}

func Test_check(t *testing.T) {
	type args struct {
		path   string
		info   os.FileInfo
		err    error
		option Option
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "#1 normal case", args: args{path: "file.jpg", info: fakeFileInfo("file.jpg"), option: Option{Input: []string{"jpg"}}}, want: true},
		{name: "#2 different extension", args: args{path: "file.png", info: fakeFileInfo("file.png"), option: Option{Input: []string{"jpg"}}}, want: false},
		{name: "#3 directory", args: args{path: "dir", info: fakeDirInfo("dir"), option: Option{Input: []string{"jpg"}}}, want: false},
		{name: "#4 some error", args: args{path: "file.jpg", info: fakeFileInfo("file.jpg"), err: errors.New("some error"), option: Option{Input: []string{"jpg"}}}, want: false},
		{name: "#5 file info is nil", args: args{path: "file.jpg", info: nil, option: Option{Input: []string{"jpg"}}}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.path, tt.args.info, tt.args.err, tt.args.option); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

type fakeFileInfo string

func (fi fakeFileInfo) Name() string    { return string(fi) }
func (fakeFileInfo) Sys() interface{}   { return nil }
func (fakeFileInfo) ModTime() time.Time { return time.Time{} }
func (fakeFileInfo) IsDir() bool        { return false }
func (fakeFileInfo) Size() int64        { return 0 }
func (fakeFileInfo) Mode() os.FileMode  { return 0644 }

type fakeDirInfo string

func (fd fakeDirInfo) Name() string    { return string(fd) }
func (fakeDirInfo) Sys() interface{}   { return nil }
func (fakeDirInfo) ModTime() time.Time { return time.Time{} }
func (fakeDirInfo) IsDir() bool        { return true }
func (fakeDirInfo) Size() int64        { return 0 }
func (fakeDirInfo) Mode() os.FileMode  { return 0755 }

func TestCli_print(t *testing.T) {
	type args struct {
		in  string
		out string
		opt Option
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1 normal case", args: args{in: "newFile.jpg", out: "newFile.png", opt: Option{Overwrite: false}}, want: "../testdata/newFile.jpg -> ../testdata/newFile.png\n"},
		{name: "#2 already exists", args: args{in: "lena.png", out: "lena.jpg", opt: Option{Overwrite: false}}, want: "../testdata/lena.png -> ../testdata/lena.jpg (skip: ../testdata/lena.jpg already exists)\n"},
		{name: "#3 Overwrite already exists", args: args{in: "lena.png", out: "lena.jpg", opt: Option{Overwrite: true}}, want: "../testdata/lena.png -> ../testdata/lena.jpg (Overwrite)\n"},
		{name: "#4 new file when Overwrite is true", args: args{in: "newFile.jpg", out: "newFile.png", opt: Option{Overwrite: true}}, want: "../testdata/newFile.jpg -> ../testdata/newFile.png\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := &bytes.Buffer{}
			c := &Cli{Out: out, Err: err}
			c.print(testdata(tt.args.in), testdata(tt.args.out), tt.args.opt)
			got := out.String()
			if got != tt.want {
				t.Errorf("print() %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_changeExt(t *testing.T) {
	type args struct {
		path string
		ext  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1 normal case", args: args{path: "file.jpg", ext: "png"}, want: "file.png"},
		{name: "#2 directory and file name", args: args{path: "/path/to/file.jpg", ext: "png"}, want: "/path/to/file.png"},
		{name: "#3 same extension", args: args{path: "file.jpg", ext: "jpg"}, want: "file.jpg"},
		{name: "#4 double extension", args: args{path: "file.jpg.bk", ext: "png"}, want: "file.jpg.png"},
		{name: "#5 no extension", args: args{path: "file", ext: "png"}, want: "file.png"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeExt(tt.args.path, tt.args.ext); got != tt.want {
				t.Errorf("changeExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "#1 normal case", args: args{path: "lena.jpg"}, want: true},
		{name: "#2 not exists", args: args{path: "not_exists.jpg"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exists(testdata(tt.args.path)); got != tt.want {
				t.Errorf("%s exists() = %v, want %v", tt.args.path, got, tt.want)
			}
		})
	}
}

func testdata(paths ...string) string {
	paths = append(paths[0:1], paths[0:]...)
	paths[0] = "../testdata"
	return path.Join(paths...)
}
