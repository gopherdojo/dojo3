/*
Package provides image convert process and functions.
These are necessary to run imageConverter.
*/
package convert

import (
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gopherdojo/dojo3/kadai1/shuntaka9576/cli"
)

// Convert image process function
func ConvertImagesProcess(option cli.ConvertOption) error {
	outdir := "converted_" + time.Now().Format("20060102-150405")
	if err := os.Mkdir(outdir, 0777); err != nil {
		log.Fatalf("mkdir err %v\n", err)
	}
	outputPath, err := filepath.Abs(outdir)
	if err != nil {
		return err
	}

	filepaths := Dirwalk(option.Targetdir)

	var extractionPaths []string
	for _, filep := range filepaths {
		ex := strings.LastIndex(filep, ".")
		if option.From == filep[ex+1:] {
			p, err := filepath.Abs(filep)
			if err != nil {
				return err
			}
			extractionPaths = append(extractionPaths, p)
		}
	}

	for _, extractionPath := range extractionPaths {
		log.Printf("Start Convert [%v]\n", extractionPath)
		ConvertImage(option.From, option.To, extractionPath, outputPath)
		log.Printf("Convert Success [%v]\n", extractionPath)
	}

	return nil
}

// Recursively directory search function
func Dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, Dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

// Convert per image file function
func ConvertImage(from, to, picpath, outputpath string) {
	imgfile, err := os.Open(picpath)
	img, _, err := image.Decode(imgfile)
	if err != nil {
		log.Fatalf("open err %v", err)
	}
	defer imgfile.Close()

	ex := strings.LastIndex(picpath, `\`)
	_, filename := picpath[:ex], picpath[ex+1:]

	convertimg, err := os.Create(filepath.Join(outputpath, filename+"_c."+to))
	defer convertimg.Close()

	if err != nil {
		log.Fatalf("convert err %v", err)
	}

	switch to {
	case "jpeg", "jpg":
		jpeg.Encode(convertimg, img, nil)
	case "png":
		png.Encode(convertimg, img)
	}
}
