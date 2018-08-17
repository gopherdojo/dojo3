package converter

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
)

func Convert(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("[Error]", err)
		return
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("[Error]", err)
		return
	}

	outputFile, err := os.Create(getFileName(path))
	if err != nil {
		fmt.Println("[Error]", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	if err != nil {
		fmt.Println("[Error]", err)
	}
}

func getFileName(path string) string {
	ext := filepath.Ext(path)
	rep := regexp.MustCompile(ext + "$")
	dest := filepath.Base(rep.ReplaceAllString(path, ""))
	return dest + ".png"
}
