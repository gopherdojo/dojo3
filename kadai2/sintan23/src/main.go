package main

import (
	"flag"

	"./convert"
)

var (
	src          = flag.String("s", "", "変換するフォルダパスを指定")
	fromExt      = flag.String("i", "jpg", "変換元の画像拡張子を指定（default: jpg）")
	toExt        = flag.String("o", "png", "変換後の画像拡張子を指定（default: png）")
	debugFlg     = flag.Bool("debug", false, "デバック用: 指定でログが標準出力される")
	goroutineFlg = flag.Bool("multi", false, "Goルーチン用: 指定すると並列処理する")
)

func main() {
	flag.Parse()
	opts := map[string]bool{
		"debugFlg":     *debugFlg,
		"goroutineFlg": *goroutineFlg,
	}

	convert.SetPath(src)
	convert.SetOpts(opts)
	convert.SetExtension(*fromExt)
	convert.ConvertAll(*toExt)
}
