package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	flag.Parse()
	file, err := os.Open(flag.Arg(0))
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
