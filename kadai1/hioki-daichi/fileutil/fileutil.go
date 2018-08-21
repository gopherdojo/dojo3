/*
Package fileutil is a collection of convenient functions for manipulating files
*/
package fileutil

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// StartsContentsWith returns whether file contents start with specified bytes.
func StartsContentsWith(rs io.ReadSeeker, xs []uint8) bool {
	buf := make([]byte, len(xs))
	rs.Seek(0, 0)
	rs.Read(buf)
	rs.Seek(0, 0)
	return bytes.Equal(buf, xs)
}

// CopyDirRec copies src directory to dest recursively.
func CopyDirRec(src string, dest string) error {
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		sf, err := os.Open(path)
		if err != nil {
			return err
		}

		destDir := filepath.Join(dest, strings.TrimLeft(filepath.Dir(path), src))

		err = os.MkdirAll(destDir, 0755)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, filepath.Base(path))

		df, err := os.Create(destPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(df, sf)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
