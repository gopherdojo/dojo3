package main

import (
	"net/http"

	"github.com/gopherdojo/dojo3/kadai4/tk3fftk/omikuji"
)

func main() {
	handler := omikuji.New(nil)
	http.Handle("/lot", &handler)
	http.ListenAndServe(":8080", nil)
}
