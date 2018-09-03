package main

import (
	"os"
	"github.com/gopherdojo/dojo3/kadai3-1/haijima/typing"
)

func main() {
	words := []string{"hello", "world", "gopher"}
	typing.Exec(os.Stdin, os.Stdout, words, 15)
}
