package main

import (
	"flag"
	"net/http"
	"github.com/gopherdojo/dojo3/kadai4/matsumatsu20/omikuji"
	"log"
)

var port = flag.String("p", "8080", "listen port")

func main() {
	http.HandleFunc("/kuji", omikuji.Handler)

	if err := http.ListenAndServe(":" + *port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}