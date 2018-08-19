package gimupop

import (
	"io"
	"flag"
	"fmt"
	"path/filepath"
	"os"
	"strings"
	"github.com/gopherdojo/dojo3/kadai1/gimupop/imgconv"
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

const (
	success       = 0
	flagError     = 1
	traverseError = 1
)

//実行
func (c *CLI) Run(args []string) int {
	var from, to, targetDir string
	//取得
	flag.StringVar(&from, "from", "", "	Input file format. (Ex. jpg, png, gif)\n")
	flag.StringVar(&to, "to", "", "Output file format. (Ex. jpg, png, gif)\n")
	flag.StringVar(&targetDir, "target", "", "Target Directory \n")
	flag.Parse()
	//確認
	if from == "" || to == "" || targetDir == "" {
		fmt.Fprintf(c.ErrStream, "ParameterError: from: %v, to: %v. target:%v \n", from, to, targetDir)
		return flagError
	}
	//実行
	if err := DoRecursion(targetDir, from, to); err != nil {
		fmt.Fprintf(c.ErrStream, "DoRecursion error: %v \n", err)
		return traverseError
	}

	fmt.Fprintf(c.OutStream, "complete!! \n")
	return success
}

//ファイル指定(再起)
func DoRecursion(srcDir, fromExt, toExt string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		//チェック
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := strings.TrimPrefix(filepath.Ext(path), ".")
		if strings.ToLower(ext) != fromExt {
			return nil
		}
		//読み込み指定
		r, err := os.Open(path)
		if err != nil {
			return err
		}
		defer r.Close()
		//書き込み指定
		w, err := os.Create(strings.TrimSuffix(path, ext) + toExt)
		if err != nil {
			return err
		}
		defer w.Close()
		//実行
		converter := imgconv.ImageConverter{}
		if err := converter.Convert(r, w, fromExt, toExt); err != nil {
			return err
		}
		return nil
	})
}
