package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// Converter struct
type Converter struct {
	Path    string
	Files   []Image
	FromExt string
	ToExt   string
}

// NewConverter is creating Converter struct
func NewConverter(path string, fromExt string, toExt string) Converter {
	return Converter{
		Path:    path,
		FromExt: fromExt,
		ToExt:   toExt,
	}
}

// FetchConverter is queuing image
func (c *Converter) FetchConverter(q chan Image, wg *sync.WaitGroup) {
	for {
		image, more := <-q
		if more {
			err := c.Convert(image)
			if err != nil {
				//TODO: error hundling
			}
		} else {
			wg.Done()
			return
		}
	}
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

	outputFile, err := os.Create(i.GetFileName(c.ToExt))
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
			i, err := NewImage(path)
			if err != nil {
				return err
			}
			c.Files = append(c.Files, i)
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
