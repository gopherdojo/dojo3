package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	ok := execute(flag.Args())
	if ok {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func execute(dirnames []string) (ok bool) {
	ok = true
	if len(dirnames) == 0 {
		fmt.Println("Specify filenames as an arguments")
		ok = false
		return
	}

DIRNAMES_LOOP:
	for _, dirname := range dirnames {
		err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			fp, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fp.Close()

			fmt.Println(path)

			return nil
		})

		if err != nil {
			fmt.Println(err)
			ok = false
			break DIRNAMES_LOOP
		}
	}
	return
}
