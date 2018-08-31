package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
)

// Converter struct
type Converter struct {
	FromExt string
	ToExt   string
}

// Convert image functon
func (c *Converter) Convert(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := c.decodeImage(file)
	if err != nil {
		return err
	}

	fileName, err := c.getFileName(path)
	if err != nil {
		return err
	}
	outputFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = c.encodeImage(outputFile, img)
	if err != nil {
		return err
	}

	return nil
}

func checkExtension(ext string) (string, error) {
	if ext == "" {
		return "", errors.New("ext must not be empty")
	} else if ext == ".jpeg" {
		return ".jpg", nil
	}
	return ext, nil
}

func (c *Converter) getFileName(path string) (string, error) {
	ext := c.ToExt
	if ext == "" {
		return "", errors.New("path must not be empty")
	}

	imageExt := filepath.Ext(path)
	rep := regexp.MustCompile(imageExt + "$")
	name := filepath.Base(rep.ReplaceAllString(path, ""))

	return name + "." + ext, nil
}

func (c *Converter) encodeImage(file io.Writer, img image.Image) error {
	var err error
	switch c.ToExt {
	case "jpeg", "jpg":
		err = jpeg.Encode(file, img, nil)
	case "gif":
		err = gif.Encode(file, img, nil)
	case "png":
		err = png.Encode(file, img)
	}
	if err != nil {
		return err
	}

	return nil
}

func (c *Converter) decodeImage(file io.Reader) (image.Image, error) {
	var (
		img image.Image
		err error
	)
	switch c.FromExt {
	case "jpeg", "jpg":
		img, err = jpeg.Decode(file)
	case "gif":
		img, err = gif.Decode(file)
	case "png":
		img, err = png.Decode(file)
	}
	if err != nil {
		return nil, err
	}

	return img, nil
}
