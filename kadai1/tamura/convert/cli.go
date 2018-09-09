package convert

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	format string
)

func init() {
	flag.StringVar(&format, "f", "", "ディレクトリ指定の場合に、変換する画像ファイルのフォーマット（`png|jpeg|jpg`）")
	flag.Parse()
}

type file string

func (f file) Ext() string {
	return filepath.Ext(string(f))
}

func (f file) Dir() string {
	return filepath.Dir(string(f))
}

func (f file) Name() string {
	return strings.Replace(filepath.Base(string(f)), f.Ext(), "", -1)
}

func Run() error {
	args := flag.Args()

	if len(args) < 1 {
		return fmt.Errorf("画像ファイルを指定してください。")
	}

	info, err := os.Stat(args[0])
	if os.IsNotExist(err) {
		return fmt.Errorf("画像ファイルが存在しません。%s", args[0])
	}

	if info.IsDir() {
		t, err := template.New("dst").Parse(args[1])
		if err != nil {
			return err
		}

		return filepath.Walk(args[0], func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			ext := strings.ToLower(filepath.Ext(p))

			if format != "" {
				if ext != "."+format {
					return nil
				}
			} else if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
				return nil
			}

			var buf bytes.Buffer
			t.Execute(&buf, file(p))

			imgConvert := Convert{
				dst: buf.String(),
				src: p,
			}

			return imgConvert.Convert()
		})
	}

	imgConvert := Convert{
		dst: os.Args[2],
		src: os.Args[3],
	}

	return imgConvert.Convert()
}
