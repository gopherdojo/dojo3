package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/gopherdojo/dojo3/kadai3-2/hioki-daichi/downloading"
	"github.com/gopherdojo/dojo3/kadai3-2/hioki-daichi/opt"
	"github.com/gopherdojo/dojo3/kadai3-2/hioki-daichi/termination"
)

func main() {
	err := execute(os.Stdout, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func execute(w io.Writer, args []string) error {
	ctx := context.Background()

	ctx, clean := termination.Listen(ctx, w)
	defer clean()

	opts, err := opt.Parse(args...)
	if err != nil {
		return err
	}

	return downloading.NewDownloader(w, opts).Download(ctx)
}
