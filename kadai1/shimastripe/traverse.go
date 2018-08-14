package kadai1

import (
	"os"
	"path/filepath"
	"strings"
)

// Traverse is converting image files recursively in the specified directory
func Traverse(srcDir, fromExt, toExt string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err // error about reading file info
		}

		if info.IsDir() {
			return nil
		}

		ext := strings.TrimPrefix(filepath.Ext(path), ".")
		if strings.ToLower(ext) != fromExt {
			return nil
		}

		r, err := os.Open(path)
		if err != nil {
			return err
		}
		defer r.Close()

		w, err := os.Create(strings.TrimSuffix(path, ext) + toExt)
		if err != nil {
			return err
		}
		defer w.Close()

		imgconv := ImageConverter{}

		if err := imgconv.Convert(r, w, fromExt, toExt); err != nil {
			return err
		}

		return nil
	})
}
