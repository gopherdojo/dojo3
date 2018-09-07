package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	url := "https://dummyimage.com/600x400.jpg"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	fmt.Println("status:", res.Status)

	_, filename := path.Split(url)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := file.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(file, res.Body)
}
