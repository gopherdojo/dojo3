package imgconv

import (
	"flag"
	"errors"
	"io"
	"strings"
	"fmt"
)

// Argument is a struct describes passed arguments
type Argument struct {
	Option Option
	Dir    string
}

// Option is options for the command
type Option struct {
	Input         []string
	Output        string
	Overwrite     bool
	DryRun        bool
	Quiet         bool
	CaseSensitive bool
}

// Validate arguments and options
func (a *Argument) Validate() error {
	if a.Option.DryRun && a.Option.Quiet {
		return errors.New("should print when dry run")
	}
	for _, ext := range a.Option.Input {
		if strings.EqualFold(ext, a.Option.Output) {
			return errors.New("output extension should be different from input extension")
		}
	}
	for _, ext := range a.Option.Input {
		lowerExt := strings.ToLower(ext)
		switch lowerExt {
		case "jpg":
		case "jpeg":
		case "png":
		case "gif":
		default:
			return fmt.Errorf("input extension support only jpg, jpeg, png and gif. but was: %s", ext)
		}
	}
	switch a.Option.Output {
	case "jpg":
	case "png":
	case "gif":
	default:
		return fmt.Errorf("output extension support only jpg, png and gif. but was: %s", a.Option.Output)
	}
	return nil
}

func (a *Argument) String() string {
	return fmt.Sprintf("Argument{ "+
		"Option: %v, "+
		"Dir: %s "+
		"}", &a.Option, a.Dir)
}

// CreateArgument creates Argument object from commandline arguments.
func CreateArgument(args []string, out io.Writer) (Argument, error) {
	defaultArg := createDefaultArg()

	flags := flag.NewFlagSet("conv", flag.ContinueOnError)
	flags.SetOutput(out)
	flags.Usage = func() {
		w := flags.Output()
		fmt.Fprintln(w, "Usage of conv:")
		fmt.Fprintln(w, "conv [-i srcExts] [-o destExt] [-w] [--dry-run|-q] [-s] [directory]")
		flags.PrintDefaults()
	}

	var input string
	flags.StringVar(&input, "i", strings.Join(defaultArg.Option.Input, "|"), "Input extension.")
	var output string
	flags.StringVar(&output, "o", defaultArg.Option.Output, "Output extension.")
	var overwrite bool
	flags.BoolVar(&overwrite, "w", defaultArg.Option.Overwrite, "If converted file has already existed, Overwrite old files.")
	var dryRun bool
	flags.BoolVar(&dryRun, "dry-run", defaultArg.Option.DryRun, "Dry run mode")
	var quiet bool
	flags.BoolVar(&quiet, "q", defaultArg.Option.Quiet, "Quiet mode. Suppress print")
	var caseSensitive bool
	flags.BoolVar(&caseSensitive, "s", defaultArg.Option.CaseSensitive, "Matches file extension case-sensitively")

	if err := flags.Parse(args[1:]); err != nil {
		return Argument{}, err
	}

	if len(flags.Args()) > 1 {
		return Argument{}, errors.New("too many dir is specified")
	}

	var dir string
	if len(flags.Args()) == 0 {
		dir = defaultArg.Dir
	} else {
		dir = flags.Args()[0]
	}

	return Argument{
		Dir: dir,
		Option: Option{
			Input:         strings.Split(input, "|"),
			Output:        output,
			Overwrite:     overwrite,
			DryRun:        dryRun,
			Quiet:         quiet,
			CaseSensitive: caseSensitive,
		},
	}, nil
}

// Create default argument and options
func createDefaultArg() Argument {
	defaultDir := "."
	defaultOpt := Option{Input: []string{"jpg", "jpeg"}, Output: "png"}
	return Argument{Dir: defaultDir, Option: defaultOpt}
}

func (o *Option) String() string {
	return fmt.Sprintf("Option{ "+
		"Input: %v, "+
		"Output: %v, "+
		"Overwrite: %v, "+
		"DryRun: %v, "+
		"Quiet: %v, "+
		"CaseSensitive: %v "+
		"}", o.Input, o.Output, o.Overwrite, o.DryRun, o.Quiet, o.CaseSensitive)
}
