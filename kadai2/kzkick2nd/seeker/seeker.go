package seeker

import (
	"os"
	"path/filepath"
)

type Dest struct {
	Dir   string
	Ext   string
	Paths []string
}

func (d *Dest) Seek() ([]string, error) {
	err := filepath.Walk(d.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		// TODO 大文字小文字と表記揺れも対応したい
		if filepath.Ext(path) == d.Ext {
			d.Paths = append(d.Paths, path)
		}
		return nil
	})
	return d.Paths, err
}
