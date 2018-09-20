/*
ConvertImage
*/
package convert

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var (
	toExt string
)

type Converter struct {
	src    string
	toExt  string
	reader io.Reader
}

type Converters []Converter

var (
	// デバックフラグ true:デバックコメント表示
	DEBUG = false

	// Goルーチンフラグ true:並列処理する
	GOROUTINE = false

	// ベースパス
	basePath = ""

	// デフォルト拡張子
	baseExtension = ""

	// フォルダから抽出する正規表現パターン
	// (最後の/除去はGlobの方で実装済みだった)
	basePattern = "/*."
)

func ConvertAll(toExt string) {
	debug("- Convert From: " + getExtension() + " To: " + toExt)

	var waitGroup sync.WaitGroup
	for _, c := range getFileList() {
		if GOROUTINE {
			// goルーチン
			waitGroup.Add(1)
			go func() {
				defer waitGroup.Done()
				err := c.Convert(toExt)
				logErr(err)
			}()
		} else {
			err := c.Convert(toExt)
			logErr(err)
		}
	}

	// チャネルを待つ
	if GOROUTINE {
		waitGroup.Wait()
	}
}

func (c Converter) Convert(toExt string) (err error) {
	debug("-- Doing... " + filepath.Ext(c.src) + " → " + toExt)

	switch {
	case getExtension() == "jpg" && toExt == "png":
		c.ConvertJpegToPng(toExt)
	}

	debug("--- Output: " + c.getFilename(toExt))

	return err
}

func (c Converter) getFilename(ext string) string {
	d, f := filepath.Split(c.src)
	f = filepath.Base(f[:len(f)-len(filepath.Ext(f))])
	path := filepath.Join(d, f+"."+ext)

	return path
}

func Open(src string) (io.Reader, error) {
	r, err := os.Open(src)
	logErr(err)

	return r, err
}

func getFileList() Converters {
	var files Converters

	debug("- Find Pattern: ", getPath()+getPattern())

	list, err := filepath.Glob(getPath() + getPattern())
	logErr(err)

	for _, src := range list {
		r, err := Open(src)
		logErr(err)
		files = append(files, Converter{
			src,
			toExt,
			r,
		})
	}

	return files
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
