/*
Package provides cli process functions.
These are necessary to run imageConverter.
*/
package cli

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeError
)

// This cli program status object
type CliStatus struct {
	Code    int
	Message string
}

// Return this program error object
func (e *CliStatus) Error() string {
	return fmt.Sprintf("%v[code=%v]", e.Message, e.Code)
}

type CLI struct {
	OutStream, ErrStream io.Writer
}

type ConvertOption struct {
	From, To, Targetdir string
}

const Version string = "v0.1.0"

// Return ConvertOption object function
func (c *CLI) GetOption(args []string) (option ConvertOption, err error) {
	var version bool

	flags := flag.NewFlagSet("imageConverter", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)

	// set argument options
	flags.BoolVar(&version, "version", false, "Print version information and quit")
	flags.StringVar(&option.From, "f", "jpg", "Print version information and quit")
	flags.StringVar(&option.From, "from", "jpg", "Print version information and quit")
	flags.StringVar(&option.To, "t", "png", "specified To files type")
	flags.StringVar(&option.To, "to", "png", "specified To files type")

	// flags parse error handling
	if err := flags.Parse(args[1:]); err != nil {
		return option, &CliStatus{ExitCodeParseFlagError, fmt.Sprintf("Parse Error please check argument[%v]", err)}
	}

	// no-flag argument handling
	noflagArgs := flags.Args()
	switch {
	case len(noflagArgs) <= 0:
		option.Targetdir = "."
	case len(noflagArgs) == 1:
		option.Targetdir = noflagArgs[0]
	default:
		return option, &CliStatus{ExitCodeError, ""}
	}

	// version option
	if version {
		fmt.Fprintf(c.OutStream, "imageConverter version %s\n", Version)
		return option, &CliStatus{ExitCodeOK, ""}
	}

	// check option object
	if err := ValidationCheck(option); err != nil {
		fmt.Fprint(c.ErrStream, "Validation check Error. Please check your from or to options\n")
		return option, err
	}

	// return option obj
	fmt.Fprintf(c.OutStream, "Start Convert %v to %v[%v]!\n", option.From, option.To, option.Targetdir)
	return option, nil
}

// Option object validation check function
func ValidationCheck(option ConvertOption) error {
	extentionOptions := []string{"jpeg", "jpg", "png"}
	type ValidationCheck struct {
		fromCheckflag bool
		toCheckflag   bool
	}
	check := ValidationCheck{false, false}

	for _, extention := range extentionOptions {
		if option.From == extention {
			check.fromCheckflag = true
		}
		if option.To == extention {
			check.toCheckflag = true
		}
	}

	if check.fromCheckflag == false || check.toCheckflag == false {
		return &CliStatus{ExitCodeError, "Validation check error! Please check your from or to option arguments"}
	}
	return nil
}
