package cli

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai1/matsumatsu20/converter"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
)

var (
	inputFormat  = flag.String("i", "jpeg", "help message for s option")
	outputFormat = flag.String("o", "png", "help message for s option")
)

type CLT struct {
	OutStream, ErrStream io.Writer
}

func (c *CLT) Run(args []string) int {
	flag.Parse()
	dir := flag.Arg(0)

	err := validateArgs(dir, *inputFormat, *outputFormat)
	if err != nil {
		log.Println(err)
		return ExitCodeParseFlagError
	}

	err = filepath.Walk(dir, convertImage)

	if err != nil {
		log.Println(err)
		return ExitCodeParseFlagError
	}

	return ExitCodeOK
}

/*
	Validate arguments of command.
	If there are invalid args, return error.
*/
func validateArgs(dir string, inputFormat string, outputFormat string) error {
	fInfo, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("%v: Invalid Arg\n", dir)
	}

	if !fInfo.IsDir() {
		return fmt.Errorf("%v: Not a directory\n", dir)
	}

	if err := converter.ValidateFormat(inputFormat); err != nil {
		return err
	}
	if err := converter.ValidateFormat(outputFormat); err != nil {
		return err
	}

	return err
}

func convertImage(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	imgConverter := converter.Converter{
		FilePath:     path,
		InputFormat:  *inputFormat,
		OutputFormat: *outputFormat,
	}

	err = imgConverter.Convert()

	return err
}
