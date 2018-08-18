package main

import (
	"os"

	c "github.com/gopherdojo/dojo3/kadai1/daikurosawa/cli"
	_ "github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert/gif"
	_ "github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert/jpg"
	_ "github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert/png"
)

func main() {
	cli := &c.Cli{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
