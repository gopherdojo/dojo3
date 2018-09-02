package cli

import (
	"os"
	"fmt"
	"errors"
	"github.com/gopherdojo/dojo3/kadai2/pchatsu"
)

var (
	validFormat = map[string]struct{}{"gif": {}, "jpeg": {}, "png": {}}
)

var (
	ErrSameFormat = errors.New("convert to the same format")
)

func Run(path string, src string, dst string) error {
	if err := validate(path, src, dst); err != nil {
		return err
	}
	if err := imgconv.Convert(path, src, dst); err != nil {
		return err
	}

	return nil
}

func validate(path string, srcExt string, dstExt string) error {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return fmt.Errorf("%s no such directory", path)
	}

	if !isValidFormatType(srcExt) || !isValidFormatType(dstExt) {
		return errors.New("available formats are gif, jpeg and png")
	}

	if srcExt == dstExt {
		return ErrSameFormat
	}

	return nil
}

func isValidFormatType(f string) bool {
	_, ok := validFormat[f]
	return ok
}
