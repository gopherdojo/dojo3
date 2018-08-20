/*
	Change image format.
	Available format are png, jpeg, gif.
 */
package convertImage

import (
	"os"
	"fmt"
	"net/http"
	"path/filepath"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
)

//Abstracted image
type AbstractImage struct {
	in image.Image
}

//encode image according to the destination type
func(ai AbstractImage) encode(out *os.File, destType *string) {
	switch *destType {
	case "png":
		png.Encode(out, ai.in)
	case "jpeg":
		jpeg.Encode(out, ai.in, nil)
	case "gif":
		gif.Encode(out, ai.in, nil)
	}
}

//convert image file from one to another
func Convert(targetDir *string, src string, srcType *string, destType *string) {
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileType, err := getFileContentType(file)
	if err != nil {
		panic(err)
	}

	if fileType != "image/"+*srcType {
		fmt.Println(src + ": image format is not " + *srcType)
		return
	}

	file.Seek(0, 0)
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	ai := AbstractImage{in:img}

	out, err := os.Create(*targetDir + getFileNameWithoutExt(src) + "." + *destType)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	ai.encode(out, destType)
}

func getFileContentType(out *os.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
