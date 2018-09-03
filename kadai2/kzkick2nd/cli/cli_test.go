package cli

import (
  "testing"
)

// FIXME フォルダ構造のassert追加

func TestRunWithArgs(t *testing.T) {

  patterns := []struct {
		name string
		inputExt string
		outputExt string
		dir string
	}{
		{
      name: "jpg to png with dir",
      inputExt: "jpg",
      outputExt: "png",
      dir: "../convert/testdata/",
    },
    {
      name: "jpg to png with 1 file",
      inputExt: "jpg",
      outputExt: "png",
      dir: "../convert/testdata/1px.jpg",
    },
	}
  for _, p := range patterns {
    t.Run(p.name, func(t *testing.T) {
      workerRun(t, p.name, p.inputExt, p.outputExt, p.dir)
    })
  }
}

func workerRun(t *testing.T, name string, inputExt string, outputExt string, dir string ){
  t.Helper()
  w := &Worker{
    Input:  inputExt,
    Output: outputExt,
  }
  err := w.Run(dir)
  if err != nil {
    t.Errorf("%s: %v",name, err)
  }
}
