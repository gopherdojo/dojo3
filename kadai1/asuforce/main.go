package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

var path string

func init() {
	flag.Parse()
	path = flag.Arg(0)
}

func main() {
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
