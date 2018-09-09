package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/gopherdojo/dojo3/kadai3-1/shuntaka9576/words"
)

func main() {
	file, err := os.OpenFile("word.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	typingWords,err := words.New(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	//game := typing.New().Set(typingWords)
	//game.Run()
}
