package gathering

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/cliopt"
	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/conversion"
)

// Gatherer has Decoder.
type Gatherer struct {
	Decoder   conversion.Decoder
	OutStream io.Writer
}

// Gather gathers files whose format is own Decoder.
func (g *Gatherer) Gather(dirname string) ([]string, error) {
	var paths []string

	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if cliopt.Verbose {
				fmt.Fprintf(g.OutStream, "Skipped because the path is directory: %q\n", path)
			}
			return nil
		}

		fp, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fp.Close()

		if !g.Decoder.IsValid(fp) {
			if cliopt.Verbose {
				fmt.Fprintf(g.OutStream, "Skipped because the file is not applicable: %q\n", path)
			}
			return nil
		}

		paths = append(paths, path)

		return nil
	})

	return paths, err
}
