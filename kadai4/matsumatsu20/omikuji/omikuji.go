package omikuji

import (
	"net/http"
	"math/rand"
	"time"
	"github.com/gopherdojo/dojo3/kadai4/matsumatsu20/dateUtil"
	"encoding/json"
	"log"
)

type response struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

const (
	daikichi = "大吉"
	kichi    = "吉"
	chukichi = "中吉"
	shokichi = "小吉"
	suekichi = "末吉"
	kyo      = "凶"
	daikyo   = "大凶"
)

var(
	luck                  = []string{daikichi, kichi, chukichi, shokichi, suekichi, kyo, daikyo}
	isNewYearsHolidayFunc = dateUtil.IsNewYearsHoliday
	fetchKujiFunc         = fetchKuji
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var result string

	if isNewYearsHolidayFunc() {
		result = daikichi
	} else {
		rand.Seed(time.Now().UnixNano())
		result = fetchKujiFunc()
	}

	res := &response{Status: 200, Result: result}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {
		log.Fatal(err)
	}
}

func fetchKuji() string {
	rand.Seed(time.Now().UnixNano())
	return luck[rand.Intn(len(luck))]
}