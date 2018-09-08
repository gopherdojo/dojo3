// Package opt provides functions and type to parse the arguments and run the command
package opt

import (
	"errors"
	"flag"
	"os"
)

// Parameter holds options to run command
type Parameter struct {
	Timeout int
}

// Parse parses the commandline arguments to Parameter
func Parse(args []string) (*Parameter, error) {
	flg := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	timeout := flg.Int("t", 15, "Time per game (second)")
	flg.Parse(args)

	if *timeout <= 0 {
		return nil, errors.New("Invalid timeout")
	}
	return &Parameter{
		Timeout: *timeout,
	}, nil
}
