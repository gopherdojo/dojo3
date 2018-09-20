package convert

import (
	"image/jpeg"
	"image/png"
	"os"
)

func (c Converter) ConvertJpegToPng(toExt string) (err error) {
	img, err := jpeg.Decode(c.reader)
	logErr(err)

	// 空ファイル作成
	out, err := os.Create(c.getFilename(toExt))
	logErr(err)

	// pngフォーマットで書き込み
	err = png.Encode(out, img)

	return err
}
