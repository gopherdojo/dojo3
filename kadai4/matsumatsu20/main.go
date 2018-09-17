package main

import (
	"flag"
	"net/http"
	"math/rand"
	"time"
	"encoding/json"
	"log"
	"github.com/gopherdojo/dojo3/kadai4/matsumatsu20/dateUtil"
)

var (
	port = flag.String("p", "8080", "resten port")

	luck = map[int]string{
		0: "大吉",
		1: "吉",
		2: "中吉",
		3: "小吉",
		4: "末吉",
		5: "凶",
		6: "大凶",
	}

)

type Response struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func main() {
	http.HandleFunc("/kuji", omikuji)

	http.ListenAndServe(":" + *port, nil)
}

func omikuji(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	i := rand.Intn(len(luck))

	// TODO: マジックナンバーを消したい
	if dateUtil.IsNewYearsHoliday() {
		i = 0
	}

	res := &Response{Status: 200, Result: luck[i]}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {log.Fatal(err)}
}