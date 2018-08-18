/*
Package cmd is a package for executing various things after evaluating flag etc in main().
*/
package cmd

import (
	"fmt"
	"io"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/conversion"
	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/gathering"
)

// Runner configures run-needed settings.
type Runner struct {
	// Usually, stdout is specified, and at the time of testing, buffer is specified.
	OutStream io.Writer

	// See conversion.{Jpeg,Png,Gif}.
	Decoder conversion.Decoder
	Encoder conversion.Encoder

	// Overwrite when the converted file name duplicates.
	Force bool
}

// Run gathers and converts the target files.
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

		fmt.Fprintf(r.OutStream, "Converted: %q\n", fp.Name())
	}

	return nil
}
