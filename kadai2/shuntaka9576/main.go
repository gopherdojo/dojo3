package main

import (
	"os"

	"github.com/gopherdojo/dojo3/kadai2/shuntaka9576/cli"
)

const (
	Version = "v0.1.0"
	Name    = "imageConverter"
)

func main() {
	os.Exit(newApp().Run(os.Args))
}

func newApp() *cli.Cli {
	app := cli.NewApp(os.Stdout, os.Stderr)
	app.Version = Version
	app.Name = Name
	return app
}
