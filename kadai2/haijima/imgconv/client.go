package imgconv

import (
	"path/filepath"
	"os"
	"fmt"
	"io"
	"strings"
)

// SuccessExitCode is zero exit code for success
const SuccessExitCode = 0

// FailExitCode is non-zero exit code for failure (1)
const FailExitCode = 1

// Cli is Command line client struct
// Search the directory and convert files
type Cli struct {
	Out, Err io.Writer
	Conv     Converter
	Dir      string
	Option   Option
}

// Run is main function of client. Returns exit status.
func (c *Cli) Run() int {
	if err := filepath.Walk(c.Dir, c.execFile); err != nil {
		fmt.Fprintln(c.Err, err)
		return FailExitCode
	}
	return SuccessExitCode
}

// execFile image format
func (c *Cli) execFile(path string, info os.FileInfo, err error) error {
	opt := c.Option

	// not target
	if !check(path, info, err, opt) {
		return nil
	}

	// create Output file path
	outPath := changeExt(path, opt.Output)

	// print message
	if !opt.Quiet {
		c.print(path, outPath, opt)
	}

	// skip condition
	if opt.DryRun || (exists(outPath) && !opt.Overwrite) {
		return nil
	}

	// Open Input file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create Output file
	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return c.Conv.Convert(file, outFile)
}

// Check whether the path is target
func check(path string, info os.FileInfo, err error, option Option) bool {
	if err != nil || info == nil || info.IsDir() {
		return false
	}

	ext := filepath.Ext(path)
	for _, targetExt := range option.Input {
		if option.CaseSensitive && ext == "."+targetExt {
			return true
		}
		if !option.CaseSensitive && strings.EqualFold(ext, "."+targetExt) {
			return true
		}
	}
	return false
}

func (c *Cli) print(in string, out string, opt Option) {
	if exists(out) {
		if opt.Overwrite {
			fmt.Fprintf(c.Out, "%s -> %s (Overwrite)\n", in, out)
		} else {
			fmt.Fprintf(c.Out, "%s -> %s (skip: %s already exists)\n", in, out, out)
		}
	} else {
		fmt.Fprintf(c.Out, "%s -> %s\n", in, out)
	}
}

func changeExt(path string, ext string) string {
	return path[0:len(path)-len(filepath.Ext(path))] + "." + ext
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
