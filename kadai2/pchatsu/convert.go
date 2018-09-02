package imgconv

import (
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"
)

// Convert converts the format of image files.
// d is the target dir and it's children images are recursively converted except hidden files.
// When a converting process occurs error, errgroup.Group as a supervisor makes the others stop.
func Convert(d string, src string, dst string) error {
	var eg errgroup.Group
	if err := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if !isTargetFile(info, path, src) {
			return nil
		}

		eg.Go(func() error {
			img, err := NewImg(path, src)
			if err != nil {
				return err
			}

			newPath := getOutputPath(path, dst)
			f, err := os.OpenFile(newPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				return err
			}
			defer f.Close()

			if err := img.Encode(f, dst); err != nil {
				return err
			}
			return nil
		})

		return nil
	}); err != nil {
		return err
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func getOutputPath(path, dst string) string {
	var newExt string
	if dst == "jpeg" {
		newExt = ".jpg"
	} else {
		newExt = "." + dst
	}

	ext := filepath.Ext(path)

	return strings.TrimSuffix(path, ext) + newExt
}

// isTargetFile returns a boolean indicating whether the file should be converted.
func isTargetFile(info os.FileInfo, path string, format string) bool {
	// skip dir
	if info.IsDir() {
		return false
	}
	// skip hidden files
	if filepath.Base(path)[0:1] == "." {
		return false
	}

	// skip unmatched files with the target format
	ext := filepath.Ext(path)
	f := strings.ToLower(strings.TrimPrefix(ext, "."))
	if f != format {
		if f != "jpg" || format != "jpeg" {
			return false
		}
	}

	return true
}
