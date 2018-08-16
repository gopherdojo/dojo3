package main

import (
	"os"
	"flag"
	"github.com/haijima/go-imgconv/imgconv"
	"fmt"
)

func main() {
	arg, err := imgconv.CreateArgument(os.Args, os.Stdout)
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(imgconv.SuccessExitCode)
		} else {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(imgconv.FailExitCode)
		}
	}
	if err := arg.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(imgconv.FailExitCode)
	}
	conv := imgconv.ImgConverter(&arg.Option)
	cli := imgconv.Cli{Out: os.Stdout, Err: os.Stderr, Conv: conv, Dir: arg.Dir, Option: arg.Option}
	os.Exit(cli.Run())
}
