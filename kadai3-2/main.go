package main

import (
	"flag"
	"net/url"
	"log"
	"os"
	"net/http"
)

func main() {
	flag.Parse()
	target := flag.Arg(0)
	validateTarget(target)
}

func validateTarget(target string) {
	u, err := url.ParseRequestURI(target)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	res, err := http.Head(u.String())

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		log.Println(err)
		return
	}
}
