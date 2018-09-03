/* Change image format. Available format are png, jpeg, gif. */
package convertImage

import (
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"io"
	"bytes"
)

//Abstracted image
type AbstractImage struct {
	in image.Image
}

//encode image according to the destination type
func (ai AbstractImage) encode(destType *string) (buf bytes.Buffer) {
	switch *destType {
	case "png":
		png.Encode(&buf, ai.in)
	case "jpeg":
		jpeg.Encode(&buf, ai.in, nil)
	case "gif":
		gif.Encode(&buf, ai.in, nil)
	}
	return buf
}

//convert image file from one to another
func Convert(in io.Reader, destType *string) (b []byte) {
	img, _, err := image.Decode(in)
	if err != nil {
		panic(err)
	}
	ai := AbstractImage{in: img}
	return ai.encode(destType).Bytes()
}
