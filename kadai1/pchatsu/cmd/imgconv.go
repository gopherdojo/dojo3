package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gopherdojo/dojo3/kadai1/pchatsu"
)

//次の仕様を満たすコマンドを作って下さい
//ディレクトリを指定する
//指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
//ディレクトリ以下は再帰的に処理する
//変換前と変換後の画像形式を指定できる（オプション）
//以下を満たすように開発してください
//mainパッケージと分離する
//自作パッケージと標準パッケージと準標準パッケージのみ使う
//準標準パッケージ：golang.org/x以下のパッケージ
//ユーザ定義型を作ってみる
//GoDocを生成してみる

var (
	validFormat = map[string]struct{}{"gif": {}, "jpeg": {}, "png": {}}
)

var (
	srcDir = flag.String("d", "./", "target directory")
	srcExt = flag.String("from", "jpg", "source extension")
	dstExt = flag.String("to", "png", "number lines")
	//strict = flag.Bool("strict", false, "fail the command immediately if an error occurs in converting process")
)

func main() {
	flag.Parse()
	Run(*srcDir, *srcExt, *dstExt)
}

func Run(path string, src string, dst string) {
	if err := validate(path, src, dst); err != nil {
		log.Fatal(err)
		return
	}
	imgconv.Convert(path, src, dst)
}

func validate(path string, srcExt string, dstExt string) error {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return fmt.Errorf("imgconv: %s is invalid dir path", path)
	}

	if !isValidFormatType(srcExt) || !isValidFormatType(dstExt) {
		return errors.New("imgconv: available formats are 'gif', 'jpeg', 'png'")
	}

	return nil
}

func isValidFormatType(f string) bool {
	_, ok := validFormat[f]
	return ok
}
