package main

import (
	"flag"
	"os"

	_ "github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert/gif"
	_ "github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert/jpg"
	_ "github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert/png"
	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/di"
	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/option"
)

func main() {
	var (
		from = flag.String("from", "jpg", "Input file extension.")
		to   = flag.String("to", "png", "Output file extension.")
	)
	flag.Parse()

	dirName := flag.Arg(0)

	option := &option.Option{DirName: dirName, FromExtension: *from, ToExtension: *to}

	convert := di.InjectConvert(option)
	cli := di.InjectCli(convert, option)
	os.Exit(cli.Run())
}
