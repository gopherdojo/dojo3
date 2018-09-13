package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gopherdojo/dojo3/kadai4/gosagawa/omikuzi"
)

//OmikuziResult おみくじの実行結果
type OmikuziResult struct {
	Result string `json:"result"`
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	date := r.FormValue("date")
	var data *OmikuziResult
	if date == "" {
		data = &OmikuziResult{Result: omikuzi.Draw()}
	} else {
		data = &OmikuziResult{Result: omikuzi.DrawByDate(date)}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error:", err)
	}
}
