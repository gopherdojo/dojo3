package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

type ConvertImage struct {
	outType string
	inPath  string
	outPath string
	dispLog bool
}

func (ci *ConvertImage) ConvertImage() {

	if ci.dispLog {
		fmt.Printf("%v => %v\n", ci.inPath, ci.outPath)
	}

	f, err := os.Open(ci.inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(ci.outPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	switch ci.outType {
	case "jpeg":
		//XXX control by option
		opts := &jpeg.Options{Quality: 100}
		jpeg.Encode(out, img, opts)
	case "png":
		png.Encode(out, img)
	case "gif":
		//XXX control by option
		opts := &gif.Options{}
		gif.Encode(out, img, opts)
	}
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

//XXX もっとスマートなやり方があるのでは
func IsValidImageType(imageType string) bool {
	var isValidType bool
	switch imageType {
	case "jpeg":
		isValidType = true
	case "png":
		isValidType = true
	case "gif":
		isValidType = true
	default:
		isValidType = false
	}
	return isValidType
}
