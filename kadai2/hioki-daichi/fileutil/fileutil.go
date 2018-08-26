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
func StartsContentsWith(rs io.ReadSeeker, xs []byte) (bool, error) {
	buf := make([]byte, len(xs))

	_, err := rs.Seek(0, 0)
	if err != nil {
		return false, err
	}

	_, err = rs.Read(buf)
	if err != nil {
		return false, err
	}

	_, err = rs.Seek(0, 0)
	if err != nil {
		return false, err
	}

	return bytes.Equal(buf, xs), nil
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

		createCopier := &fileCreateCopier{}
		return Copy(createCopier, path, filepath.Join(dest, strings.TrimLeft(filepath.Dir(path), src), filepath.Base(path)))
	})
	return err
}

// Copy copies src path to dest path.
func Copy(createCopier createCopier, src string, dest string) error {
	err := os.MkdirAll(filepath.Dir(dest), 0755)
	if err != nil {
		return err
	}

	df, err := createCopier.Create(dest)
	if err != nil {
		return err
	}

	sf, err := os.Open(src)
	if err != nil {
		return err
	}

	_, err = createCopier.Copy(df, sf)
	if err != nil {
		return err
	}

	return nil
}

type createCopier interface {
	Create(string) (*os.File, error)
	Copy(io.Writer, io.Reader) (written int64, err error)
}

type fileCreateCopier struct{}

func (c *fileCreateCopier) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func (c *fileCreateCopier) Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	return io.Copy(dst, src)
}
