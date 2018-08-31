package converter

import (
	"os"
	"path/filepath"
)

// Collect is survey all directory and pick path
type Collect struct {
	Paths   []string
	FromExt string
}

// CollectPath function found image file and append Converter.Files
func (c *Collect) CollectPath(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		ext, err := checkExtension(filepath.Ext(path))
		if ext == ("." + c.FromExt) {
			if err != nil {
				return err
			}
			c.Paths = append(c.Paths, path)
		}
	}
	return nil
}
