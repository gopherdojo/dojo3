package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/converter"
	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/parser"
	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/seeker"
)

func main() {
	c := Config{
		Out:  os.Stdout,
		Args: os.Args,
	}

	// TODO loggerの追加

	err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

type Config struct {
	Out  io.Writer
	Args []string
}

func (c *Config) Run() error {
	a, err := parser.Parse(c.Args)
	if err != nil {
		return err
	}

	d := seeker.Target{
		Dir: a.Dir,
		Ext: a.Decoder.Ext(),
	}
	p, err := d.Seek()
	if err != nil {
		return err
	}

	q := converter.Queue{
		Log:     c.Out,
		Src:     p,
		Encoder: a.Encoder,
		Decoder: a.Decoder,
	}
	err = q.Run()
	if err != nil {
		return err
	}

	return nil
}
