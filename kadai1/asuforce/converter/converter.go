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
	FromExt string
	ToExt   string
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

	img, err := c.decodeImage(file)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(i.getFileName(c.ToExt))
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

// CrawlFile function found image file and append Converter.Files
func (c *Converter) CrawlFile(path string, info os.FileInfo, err error) error {
	if checkExtension(filepath.Ext(path)) == ("." + c.FromExt) {
		if !info.IsDir() {
			c.Files = append(c.Files, newImage(path))
		}
	}
	return nil
}

func checkExtension(path string) string {
	if path == ".jpeg" {
		return ".jpg"
	}
	return path
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
