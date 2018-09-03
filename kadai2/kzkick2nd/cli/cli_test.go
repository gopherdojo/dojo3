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
      w := &Worker{
        Input:  p.inputExt,
    	  Output: p.outputExt,
      }
      err := w.Run(p.dir)
    	if err != nil {
        t.Errorf("p.name: %v", err)
    	}
    })
  }

}
