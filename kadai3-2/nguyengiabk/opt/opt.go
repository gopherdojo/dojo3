// Package opt provides functions and type to parse the arguments and run the command
package opt

import (
	"errors"
	"flag"
	"os"
)

// Parameter holds options to run command
type Parameter struct {
	ProcNum int
	URL     string
}

// Parse parses the commandline arguments to Parameter
func Parse(args []string) (*Parameter, error) {
	flg := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	procNum := flg.Int("p", 4, "Number of parallel processes")
	flg.Parse(args)
	if *procNum <= 0 {
		return nil, errors.New("Invalid number of parallel processes")
	}
	if len(flg.Args()) < 1 {
		return nil, errors.New("URL was not specified")
	}
	return &Parameter{
		ProcNum: *procNum,
		URL:     flg.Args()[0],
	}, nil
}
