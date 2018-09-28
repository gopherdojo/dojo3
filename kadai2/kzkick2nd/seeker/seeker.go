package seeker

import (
	"os"
	"path/filepath"
	"strings"
)

type Target struct {
	Dir   string
	Ext   string
	Paths []string
}

func (t *Target) Seek() ([]string, error) {
	err := filepath.Walk(t.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.ToLower(filepath.Ext(path)) == "."+t.Ext {
			t.Paths = append(t.Paths, path)
		}
		return nil
	})
	return t.Paths, err
}
