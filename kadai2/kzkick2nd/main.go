package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kzkick2nd/golang-sandbox/img-convert/converter"
	"github.com/kzkick2nd/golang-sandbox/img-convert/option"
	"github.com/kzkick2nd/golang-sandbox/img-convert/seeker"
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
	a, err := option.Parse(c.Args)
	if err != nil {
		return err
	}

	d := seeker.Dest{
		Dir: a.Dir,
		Ext: a.Decoder,
	}
	p, err := d.Seek()
	if err != nil {
		return err
	}

	q := converter.Queue{
		Log:     c.Out,
		Src:     p,
		Encoder: a.Encoder,
	}
	err = q.Run()
	if err != nil {
		return err
	}

	return nil
}
