package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

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
