package imagetypes

import (
	"errors"
	"image"
	"io"
)

type ImageType interface {
	Decode(r io.Reader) (image.Image, error)
	Encode(w io.Writer, m image.Image) error
	CheckExtStr(ext string) bool
}

var supportImageTypes = []ImageType{}

func ResisterImageType(imageType ImageType) {
	supportImageTypes = append(supportImageTypes, imageType)
}

func GetSupportImageType(extension string) (imagetype ImageType, err error) {
	for _, imagetype := range supportImageTypes {
		if imagetype.CheckExtStr(extension) {
			return imagetype, nil
			break
		}
	}
	return imagetype, errors.New("not found option: " + extension)
}
