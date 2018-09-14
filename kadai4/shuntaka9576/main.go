package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gopherdojo/dojo3/kadai4/shuntaka9576/handler"
)

func main() {
	http.HandleFunc("/omikuji", handler.OmikujiHandler)
	fmt.Fprintf(os.Stdout, "Web Server is available at http://localhost:8080/omikuji")
	http.ListenAndServe(":8080", nil)
}
