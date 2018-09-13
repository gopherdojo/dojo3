package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gopherdojo/dojo3/kadai4/gosagawa/omikuzi"
)

type OmikuziResult struct {
	Result string `json:"result"`
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {

	date := r.FormValue("date")
	var data OmikuziResult
	if date != "" {
		data = &OmikuziResult{Result: omikuzi.Draw()}
	} else {
		data = &OmikuziResult{Result: omikuzi.DrawByDate(date)}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error:", err)
	}
}

func handlerHoge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, net/http!")
}
