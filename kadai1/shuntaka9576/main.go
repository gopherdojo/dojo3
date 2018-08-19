package main

import (
	"os"

	"github.com/gopherdojo/dojo3/kadai1/shuntaka9576/cli"
	"github.com/gopherdojo/dojo3/kadai1/shuntaka9576/convert"
)

func main() {
	cliProcess := cli.CLI{os.Stdout, os.Stderr}

	if option, err := cliProcess.GetOption(os.Args); err != nil {
		switch e := err.(type) {
		case *cli.CliStatus:
			switch e.Code {
			case cli.ExitCodeOK:
				os.Exit(0)
			default:
				os.Exit(e.Code)
			}
		}
	} else {
		// Start img files convert process
		convert.ConvertImagesProcess(option)
	}
}
