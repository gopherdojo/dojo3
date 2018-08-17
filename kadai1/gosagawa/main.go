package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	var (
		inputImageType  = flag.String("i", "jpeg", "変換対象の画像形式(jpeg|gif|png)")
		outputImageType = flag.String("o", "png", "変換語の画像形式(jpeg|gif|png)")
	)
	flag.Parse()
	args := flag.Args()

	fmt.Println(*inputImageType)
	fmt.Println(*outputImageType)
	fmt.Println(args[0])

	//TODO check parameter

	dir := args[0]
	convertList := getConvertList(*inputImageType, dir)
	fmt.Println(convertList)
}

func getConvertList(imageType string, dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, getConvertList(imageType, filepath.Join(dir, file.Name()))...)
			continue
		}
		path := filepath.Join(dir, file.Name())
		if imageType == getImageType(path) {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths
}

func getImageType(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, format, err := image.DecodeConfig(f)
	if err != nil {
		return ""
	}
	return format
}
