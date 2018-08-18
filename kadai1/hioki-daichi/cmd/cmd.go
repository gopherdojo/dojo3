package cmd

import (
	"fmt"
	"io"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/conversion"
	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/gathering"
)

// Runner has streams and command line options.
type Runner struct {
	OutStream io.Writer
	Decoder   conversion.Decoder
	Encoder   conversion.Encoder
	Force     bool
	Verbose   bool
}

// Run runs command.
func (r *Runner) Run(dirname string) error {
	gatherer := &gathering.Gatherer{Decoder: r.Decoder}
	paths, err := gatherer.Gather(dirname)
	if err != nil {
		return err
	}

	converter := &conversion.Converter{Decoder: r.Decoder, Encoder: r.Encoder}

	for _, path := range paths {
		fp, err := converter.Convert(path, r.Force)
		if err != nil {
			return err
		}

		if r.Verbose {
			fmt.Fprintf(r.OutStream, "Converted: %q\n", fp.Name())
		}
	}

	return nil
}
