// Package img provides function to convert images
package img

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

// Convert converts the extension of images
// if includes directory, images of the directory are converted
func Convert(dir, from, to string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("faild to ioutil.ReadDir(dir): %s", err)
	}

	for _, file := range files {
		fileName := file.Name()

		if file.IsDir() {
			if err := Convert(dir+fileName+"/", from, to); err != nil {
				return err
			}
			continue
		}

		s := strings.LastIndex(fileName, ".")
		if fileName[s+1:] != from {
			continue
		}

		f, err := os.Open(dir + fileName)
		if err != nil {
			return fmt.Errorf("faild to open file: %s", err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			return fmt.Errorf("faild to image.Decode(f): %s", err)
		}

		out, err := os.Create(dir + fileName[:s] + "." + to)
		if err != nil {
			return fmt.Errorf("faild to os.Create(): %s", err)
		}
		defer out.Close()

		switch to {
		case "jpg":
			err = jpeg.Encode(out, img, nil)
		case "png":
			err = png.Encode(out, img)
		case "gif":
			err = gif.Encode(out, img, nil)
		}
	}

	return err
}
