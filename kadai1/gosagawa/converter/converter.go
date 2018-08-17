package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Converter struct {
	inType   string
	outType  string
	dir      string
	dispLog  bool
	inPaths  []string
	outPaths []string
}

func NewConverter(inType string, outType string, dir string, dispLog bool) *Converter {
	c := Converter{}
	c.inType = inType
	c.outType = outType
	c.dispLog = dispLog
	c.dir = dir
	return &c
}

func GetImageType(path string) string {
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

func (c *Converter) ConvertImage() {
	convertPaths := getConvertList(c.inType, c.dir)
	for _, path := range convertPaths {
		dist := getConvertToPath(c.outType, path)

		//TODO check file is already exist
		if c.dispLog {
			fmt.Printf("%v => %v\n", path, dist)
		}
		ConvertImage(c.outType, path, dist)
	}
}

func ConvertImage(outType string, src string, dest string) {
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
		if imageType == GetImageType(path) {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths
}

func getConvertToPath(outType string, path string) string {
	return path[:len(path)-len(filepath.Ext(path))] + "." + outType
}
