package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gopherdojo/dojo3/kadai4/matsumatsu20/omikuji"
)

var port = flag.String("p", "8080", "listen port")

func main() {
	http.HandleFunc("/kuji", omikuji.Handler)

	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
