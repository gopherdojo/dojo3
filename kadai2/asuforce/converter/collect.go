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

		ext, err := checkExtension(filepath.Ext(path))
		if ext == ("." + c.FromExt) {
			if err != nil {
				return err
			}
			c.Paths = append(c.Paths, path)
		}
		return nil
	})
	return err
}

func checkExtension(ext string) (string, error) {
	if ext == "" {
		return "", errors.New("ext must not be empty")
	} else if ext == ".jpeg" {
		return ".jpg", nil
	}
	return ext, nil
}
