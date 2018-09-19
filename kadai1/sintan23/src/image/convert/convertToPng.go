/*
ConvertImage
*/
package convert

import (
	"image/jpeg"
	"image/png"
	"os"
)

func convertJpegToPng(name string, toExt string) (err error) {
	// output name
	baseName := name[:len(name)-len(baseExtension)]
	outputName := baseName + toExt
	debug("     Output: " + outputName)

	reader, err := os.Open(name)
	logErr(err)
	defer reader.Close()

	// jpgをデコードimage/imageに
	img, err := jpeg.Decode(reader)
	logErr(err)

	// 空ファイル作成
	out, err := os.Create(outputName)
	logErr(err)

	// pngフォーマットで書き込み
	err = png.Encode(out, img)

	return err
}
