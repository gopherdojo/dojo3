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
)

// Converter struct
type Converter struct {
	Path    string
	Files   []Image
	DestExt string
}

// Image information struct
type Image struct {
	path string
	name string
	ext  string
}

// Convert image functon
func (c *Converter) Convert(i Image) error {
	file, err := os.Open(i.path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(i.getFileName(c.DestExt))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = c.decodeImage(outputFile, img)
	if err != nil {
		return err
	}

	return nil
}

func (c *Converter) CrawlFile(path string, info os.FileInfo, err error) error {
	if filepath.Ext(path) == ".jpg" {
		if !info.IsDir() {
			var i = newImage(path)
			c.Files = append(c.Files, i)
		}
	}
	return nil
}

func (c *Converter) decodeImage(file io.Writer, img image.Image) error {
	var err error
	switch c.DestExt {
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

func newImage(path string) Image {
	ext := filepath.Ext(path)
	rep := regexp.MustCompile(ext + "$")
	name := filepath.Base(rep.ReplaceAllString(path, ""))

	return Image{
		path: path,
		name: name,
		ext:  ext,
	}
}

func (i *Image) getFileName(ext string) string {
	return i.name + "." + ext
}
