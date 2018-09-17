package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gopherdojo/dojo3/kadai4/daikurosawa/lot"
)

func main() {
	port := flag.String("p", "8080", "listen port")
	flag.Parse()

	fmt.Fprintf(os.Stdout, "GET http://localhost:%s/lot\n", *port)

	lot := lot.NewLot(nil)
	http.HandleFunc("/lot", lot.Handler)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		fmt.Fprintln(os.Stderr, "listen and serve error.", err)
	}
	return
}
