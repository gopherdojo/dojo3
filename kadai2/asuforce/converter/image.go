package converter

import (
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
)

// Image information struct
type Image struct {
	path string
	name string
	ext  string
}

// NewImage is creating Image struct
func NewImage(path string) (Image, error) {
	if path == "" {
		return Image{}, errors.New("path must not be empty")
	}

	ext := filepath.Ext(path)
	rep := regexp.MustCompile(ext + "$")
	name := filepath.Base(rep.ReplaceAllString(path, ""))

	return Image{
		path: path,
		name: name,
		ext:  ext,
	}, nil
}

// GetFileName bind filename and extension
func (i *Image) GetFileName(ext string) string {
	return i.name + "." + ext
}
