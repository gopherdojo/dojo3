package main

import (
	"os"
	"time"

	"github.com/gopherdojo/dojo3/kadai4/shimastripe"
)

func main() {
	cli := &shimastripe.CLI{Clock: shimastripe.ClockFunc(func() time.Time {
		return time.Now()
	})}
	os.Exit(cli.Run(os.Args))
}
