package converter

import (
	"image"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai2/shuntaka9576/imagetypes"
)

type Converter struct {
	From, To imagetypes.ImageType
}

func GetConverter(from, to string) (converter Converter, err error) {
	converter.From, err = imagetypes.GetSupportImageType("." + from)
	if err != nil {
		return converter, err
	}

	converter.To, err = imagetypes.GetSupportImageType("." + to)
	if err != nil {
		return converter, err
	}
	return converter, nil
}

// If conversion failed, return empty string.
func (c *Converter) Convert(inputImagePath, outputPath string) (string, error) {
	// decode
	var decodeImage image.Image
	if c.From.CheckExtStr(filepath.Ext(inputImagePath)) {
		file, err := os.Open(inputImagePath)
		if err != nil {
			return "", err
		}
		defer file.Close()
		decodeImage, err = c.From.Decode(file)
		if err != nil {
			return "", err
		}
		// encode
		outfile, err := os.Create(outputPath)
		if err != nil {
			return "", err
		}
		err = c.To.Encode(outfile, decodeImage)
		if err != nil {
			return "", err
		}
		return outputPath, nil
	}
	return "", nil
}
