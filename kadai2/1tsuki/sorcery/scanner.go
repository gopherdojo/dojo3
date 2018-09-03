package sorcery

import (
	"path/filepath"
	"os"
)

func scan(path string, ext imgExt, callback func(string) error) error {
	return filepath.Walk(path, filter(ext, callback))
}

func filter(ext imgExt, callback func(string) error) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileExt, err := ImgExt(filepath.Ext(path))
		if err != nil {
			// skip unsupported extension
			return nil
		}

		if fileExt != ext {
			return nil
		}

		return callback(path)
	}
}
