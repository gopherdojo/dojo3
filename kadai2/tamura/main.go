package main

import (
	"fmt"
	"os"

	"./convert"
)

func main() {
	if err := convert.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
