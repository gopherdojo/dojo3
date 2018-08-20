package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

type Convert struct {
	dst string
	src string
}

func (c *Convert) Convert() error {
	sf, err := os.Open(c.src)
	if err != nil {
		return fmt.Errorf("画像ファイルが開けませんでした。%s", c.src)
	}
	defer sf.Close()

	df, err := os.Create(c.dst)
	if err != nil {
		return fmt.Errorf("画像ファイルが書き出せませんでした。%s", c.dst)
	}
	defer df.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	ext := strings.ToLower(path.Ext(c.dst))
	if ext == ".png" {
		err = png.Encode(df, img)
	} else if ext == ".jpeg" || ext == ".jpg" {
		err = jpeg.Encode(df, img, &jpeg.Options{jpeg.DefaultQuality})
	}
	return nil
}
