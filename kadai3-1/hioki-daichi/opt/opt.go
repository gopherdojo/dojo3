package opt

import (
	"flag"
	"os"
)

// Options has Timeout and Path.
type Options struct {
	Timeout int
	Path    string
}

// Parse parses the command line option and returns Options.
func Parse(args ...string) *Options {
	flg := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	timeout := flg.Int("timeout", 15, "Time limit in this typing game")
	path := flg.String("path", "./weapons.txt", "File path in which a list of words used for this typing game is described")

	flg.Parse(args)

	return &Options{Timeout: *timeout, Path: *path}
}
