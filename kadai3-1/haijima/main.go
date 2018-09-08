package main

import (
	"github.com/gopherdojo/dojo3/kadai3-1/haijima/typing"
	"os"
	"time"
)

func main() {
	timerCh := time.After(15 * time.Second)
	words := []string{"hello", "world", "gopher"}
	typing.Exec(os.Stdin, os.Stdout, timerCh, words)
}
