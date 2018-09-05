/*
ConvertImage
*/
package convert

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"sync"
)

var (
	// デバックフラグ true:デバックコメント表示
	DEBUG = false

	// Goルーチンフラグ true:並列処理する
	GOROUTINE = false

	// ベースパス
	basePath = ""

	// ベースのコンバート拡張子
	baseExtension = "jpg"

	// フォルダから抽出する正規表現パターン
	// (最後の/除去はGlobの方で実装済みだった)
	basePattern = "/*."
)

func Run() {
	// use flag
	var (
		src          = flag.String("s", "", "変換するフォルダパスを指定")
		fromExt      = flag.String("i", "jpg", "変換元の画像拡張子を指定（default: jpg）")
		toExt        = flag.String("o", "png", "変換後の画像拡張子を指定（default: png）")
		debugFlg     = flag.Bool("debug", false, "デバック用: 指定でログが標準出力される")
		goroutineFlg = flag.Bool("multi", false, "Goルーチン用: 指定すると並列処理する")
	)
	flag.Parse()
	opts := map[string]bool{
		"debugFlg":     *debugFlg,
		"goroutineFlg": *goroutineFlg,
	}

	SetOpts(opts)
	SetPath(src)
	SetExtension(*fromExt)
	ConvertAll(*toExt)
}

func ConvertAll(toExt string) {
	debug("- Convert From: " + getExtension() + " To: " + toExt)

	list := getFileList()

	var waitGroup sync.WaitGroup
	for _, n := range list {
		if GOROUTINE {
			// goルーチン
			waitGroup.Add(1)
			go func() {
				defer waitGroup.Done()
				err := Convert(n, toExt)
				logErr(err)
			}()
		} else {
			err := Convert(n, toExt)
			logErr(err)
		}
	}

	// チャネルを待つ
	if GOROUTINE {
		waitGroup.Wait()
	}
}

func Convert(name string, toExt string) (err error) {
	debug("   Doing... " + name + " → " + toExt)

	switch {
	// jpeg→png
	case getExtension() == "jpg" && toExt == "png":
		convertJpegToPng(name, toExt)
	}

	return err
}

func getFileList() []string {
	debug("- Find Pattern: ", getPath()+getPattern())

	list, err := filepath.Glob(getPath() + getPattern())
	logErr(err)

	return list
}

func SetOpts(opts map[string]bool) {
	for k, flg := range opts {
		switch {
		case k == "debugFlg" && flg:
			DEBUG = true
		case k == "goroutineFlg" && flg:
			GOROUTINE = true
		}
		debug("Option:", k, strconv.FormatBool(flg))
	}
}

func SetPath(d *string) {
	basePath = *d
}

func getPath() string {
	return basePath
}

func SetExtension(fromExt string) {
	baseExtension = fromExt
}

func getExtension() string {
	return baseExtension
}

func SetPattern(pattern string) {
	basePattern = pattern
}

func getPattern() string {
	return basePattern + getExtension()
}

// Errorチェック
func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func debug(m ...string) {
	if DEBUG {
		fmt.Printf("%+v\n", m)
	}
}
