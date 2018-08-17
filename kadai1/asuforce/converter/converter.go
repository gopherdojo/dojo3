package converter

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
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

	outputFile, err := os.Create("output.png")
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
