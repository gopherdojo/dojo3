package main

import (
	"os"

	"github.com/gopherdojo/dojo3/kadai3-1/pchatsu"
)

func main() {
	tg := typinggame.NewTypingGame(60, os.Stdin, os.Stdout)
	tg.Play()
}
