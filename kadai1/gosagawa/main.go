package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
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
		convertImage(*outType, path, dist)
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

func getConvertToPath(outType string, path string) string {
	return path[:len(path)-len(filepath.Ext(path))] + "." + outType
}

func convertImage(outType string, src string, dest string) {
	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	switch outType {
	case "jpeg":
		//TODO control by option
		opts := &jpeg.Options{}
		jpeg.Encode(out, img, opts)
	case "png":
		png.Encode(out, img)
	case "gif":
		//TODO control by option
		opts := &gif.Options{}
		gif.Encode(out, img, opts)
	}
}
