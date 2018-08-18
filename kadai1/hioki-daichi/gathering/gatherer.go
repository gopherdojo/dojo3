/*
Package gathering is a package that summarizes the processing necessary for collecting the files to be decoded.
*/
package gathering

import (
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/conversion"
)

// Gatherer represents decodable.
type Gatherer struct {
	Decoder conversion.Decoder
}

// Gather searches under the specified directory and collects files to be decoded.
func (g *Gatherer) Gather(dirname string) ([]string, error) {
	var paths []string

	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !g.Decoder.HasProcessableExtname(path) {
			return nil
		}

		fp, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fp.Close()

		if !g.Decoder.IsDecodable(fp) {
			return nil
		}

		paths = append(paths, path)

		return nil
	})

	return paths, err
}
