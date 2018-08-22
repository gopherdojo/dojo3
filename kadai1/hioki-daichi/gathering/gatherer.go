/*
Package gathering is a package that summarizes the processing necessary for collecting the files to be decoded.
*/
package gathering

import (
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/conversion"
	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/fileutil"
)

// Gatherer represents decodable.
type Gatherer struct {
	Decoder   conversion.Decoder
	Pathnames []string
}

// Gather searches under the specified directory and collects files to be decoded.
func (g *Gatherer) Gather(dirname string) ([]string, error) {
	err := filepath.Walk(dirname, g.walkFn)

	return g.Pathnames, err
}

func (g *Gatherer) walkFn(path string, info os.FileInfo, err error) error {
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

	isDecodable := false
	for _, magicBytes := range g.Decoder.MagicBytesSlice() {
		if fileutil.StartsContentsWith(fp, magicBytes) {
			isDecodable = true
			break
		}
	}
	if !isDecodable {
		return nil
	}

	g.Pathnames = append(g.Pathnames, path)

	return nil
}
