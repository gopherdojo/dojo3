package converter

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Collect is survey all directory and pick path
type Collect struct {
	Paths   []string
	FromExt string
}

// CollectPath function found image file and append Converter.Files
func (c *Collect) CollectPath(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		err = c.appendFiles(path)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (c *Collect) appendFiles(path string) error {
	if path == "" {
		return errors.New("path must not be empty")
	}

	ext := filepath.Ext(path)
	if ext == ".jpeg" {
		ext = ".jpg"
	}

	if ext == ("." + c.FromExt) {
		c.Paths = append(c.Paths, path)
	}
	return nil
}
