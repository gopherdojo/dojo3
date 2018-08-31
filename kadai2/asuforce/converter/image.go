package converter

import (
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
)

// Image information struct
type Image struct {
	path string
}

// GetFileName bind filename and extension
func (i *Image) GetFileName(ext string) (string, error) {
	if ext == "" {
		return "", errors.New("path must not be empty")
	}

	imageExt := filepath.Ext(i.path)
	rep := regexp.MustCompile(imageExt + "$")
	name := filepath.Base(rep.ReplaceAllString(i.path, ""))

	return name + "." + ext, nil
}
