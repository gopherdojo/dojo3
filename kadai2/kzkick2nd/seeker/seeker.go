package seeker

import (
	"os"
	"path/filepath"
	"strings"
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
		if strings.ToLower(filepath.Ext(path)) == "."+d.Ext {
			d.Paths = append(d.Paths, path)
		}
		return nil
	})
	return d.Paths, err
}
