package cli

import (
	"io"

	"github.com/gopherdojo/dojo3/kadai3-1/daikurosawa/word"
)

type ExportCLI = CLI

var ExportPlay = (*CLI).play

func NewExportCLI(in io.Reader, out, err io.Writer, word word.Word, ch <-chan string) *ExportCLI {
	return &ExportCLI{InStream: in, OutStream: out, ErrStream: err, Word: word, ch: ch}
}
