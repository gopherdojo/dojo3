package converter

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

type Converter struct {
	FilePath     string
	InputFormat  string
	OutputFormat string
}

var convertibleFormat = [...]string{"jpg", "jpeg", "png", "gif"}

/*
	Convert image to OutputFormat from InputFormat.
	If it is failed, return errors
*/
func (c *Converter) Convert() error {
	f, err := os.Open(c.FilePath)
	defer f.Close()

	if err != nil {
		log.Fatalf("%v: Failed to open file.\n", c.FilePath)
	}

	img, format, err := image.Decode(f)

	if err != nil {
		log.Fatalf("%v: Failed to decode file.\n", c.FilePath)
	}

	err = nil

	if c.isConvertTargetFormat(format) {
		dst, _ := os.Create(c.GetDistFileName())
		defer dst.Close()

		switch c.OutputFormat {
		case "jpeg", "jpg":
			err = jpeg.Encode(dst, img, nil)
		case "png":
			err = png.Encode(dst, img)
		case "gif":
			err = gif.Encode(dst, img, nil)
		}
	}

	return err
}

/*
	Return whether the argument format equals input format.
	In addition, perform name identification of jpeg.
*/
func (c *Converter) isConvertTargetFormat(format string) bool {
	if format == c.InputFormat || (format == "jpeg" && c.InputFormat == "jpg") {
		return true
	}
	return false
}

/*
	Get name of output file after converting.
	If file has no extension, add extension.
*/
func (c *Converter) GetDistFileName() string {
	var filePath string

	i := strings.LastIndex(c.FilePath, ".")

	if i == -1 {
		filePath = c.FilePath
	} else {
		filePath = c.FilePath[:i]
	}

	return filePath + "." + c.OutputFormat
}

// Validate argument as convertible format. If not, return error.
func ValidateFormat(format string) error {
	for _, v := range convertibleFormat {
		if format == v {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("%v: invalid format", format))
}
