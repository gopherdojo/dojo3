package main

import (
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"

	"./converter"
)

func main() {

	var (
		inType  = flag.String("i", "jpeg", "変換対象の画像形式(jpeg|gif|png)")
		outType = flag.String("o", "png", "変換語の画像形式(jpeg|gif|png)")
		dispLog = flag.Bool("v", false, "詳細なログを表示")
	)
	flag.Parse()
	args := flag.Args()

	fmt.Println(*inType)
	fmt.Println(*outType)
	fmt.Println(args[0])

	//TODO check parameter

	dir := args[0]
	convertPaths := getConvertList(*inType, dir)
	for _, path := range convertPaths {
		dist := getConvertToPath(*outType, path)

		//TODO check file is already exist
		if *dispLog {
			fmt.Printf("%v => %v\n", path, dist)
		}
		converter.ConvertImage(*outType, path, dist)
	}
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
		if imageType == converter.GetImageType(path) {
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

func getConvertToPath(outType string, path string) string {
	return path[:len(path)-len(filepath.Ext(path))] + "." + outType
}
