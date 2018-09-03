package cli

import (
  "os"
	"path/filepath"

  "github.com/gopherdojo/dojo3/kadai2/kzkick2nd/convert"
)

type Worker struct {
  Input  string
  Output string
}

func (w *Worker) Run(dir string) error {
  err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) != "."+w.Input {
			return nil
		}
		c := convert.Converter{
			Src:       path,
			OutputExt: w.Output,
		}
		errConvert := c.Convert()
		if errConvert != nil {
			return errConvert
		}
		return nil
	})
  return err
}
