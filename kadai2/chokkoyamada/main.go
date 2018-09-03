package main

import (
	"flag"
	"os"
	"log"
	"path/filepath"
	"./convertImage"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func main() {
	var (
		targetDir = flag.String("targetDir", "./", "対象ディレクトリ")
		srcType   = flag.String("srcType", "jpeg", "変換元ファイルタイプ")
		destType  = flag.String("destType", "png", "変換先ファイルタイプ")
	)
	flag.Parse()

	walkFile(targetDir, srcType, destType)
}

func walkFile(targetDir *string, srcType *string, destType *string) {
	err := filepath.Walk(*targetDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			inputFile, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer inputFile.Close()

			fileType, err := getFileContentType(inputFile)
			if err != nil {
				panic(err)
			}

			if fileType != "image/"+*srcType {
				fmt.Println(": image format is not " + *srcType)
				return errors.New("image format is not " + *srcType)
			}

			outputFile, err := os.Create(*targetDir + getFileNameWithoutExt(path) + "." + *destType)
			if err != nil {
				panic(err)
			}
			defer outputFile.Close()

			inputFile.Seek(0, 0)
			b := convertImage.Convert(inputFile, destType)
			outputFile.Write(b)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func getFileContentType(out io.Reader) (string, error) {
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
