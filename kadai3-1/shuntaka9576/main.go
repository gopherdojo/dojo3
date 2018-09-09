package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo3/kadai3-1/shuntaka9576/typing"
	"github.com/gopherdojo/dojo3/kadai3-1/shuntaka9576/words"
)

func main() {
	file, err := os.OpenFile("words.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	typingWords, err := words.New(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	game := typing.New().SetWords(typingWords)

	ctx := context.Background()
	game.Run(ctx, time.Second*30)
}
