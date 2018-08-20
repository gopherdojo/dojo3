package main

import (
	"flag"
	"os"
	"log"
	"path/filepath"
	"change_image_format/convertImage"
)

func main() {
	var (
		targetDir = flag.String("targetDir", ".", "対象ディレクトリ")
		srcType = flag.String("srcType", "jpeg", "変換元ファイルタイプ")
		destType = flag.String("destType", "png", "変換先ファイルタイプ")
	)
	flag.Parse()

	walkFile(targetDir, srcType, destType)
}

func walkFile(targetDir *string, srcType *string, destType *string) {
	err := filepath.Walk(*targetDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir(){
			convertImage.Convert(path, srcType, destType)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

