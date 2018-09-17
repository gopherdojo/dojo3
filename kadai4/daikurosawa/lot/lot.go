package lot

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gopherdojo/dojo3/kadai4/daikurosawa/clock"
)

const (
	daikiti = "大吉"
	tyukiti = "中吉"
	syokiti = "小吉"
	kiti    = "吉"
	suekiti = "末吉"
	kyo     = "凶"
	daikyo  = "大凶"
)

var results = []string{
	daikiti,
	tyukiti,
	syokiti,
	kiti,
	suekiti,
	kyo,
	daikyo}

type response struct {
	Result string `json:"result"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type lot struct {
	clock.Clock
}

func NewLot(c clock.Clock) *lot {
	return &lot{c}
}

func (l *lot) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	result := l.fetchLot()
	res := &response{Result: result}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func (l *lot) now() time.Time {
	if l.Clock == nil {
		return time.Now()
	}
	return l.Clock.Now()
}

func (l *lot) isNewYear() bool {
	date := l.now()
	_, month, day := date.Date()
	return month == time.January && (day <= 3)
}

func (l *lot) fetchLot() string {
	if l.isNewYear() {
		return daikiti
	} else {
		len := len(results)
		return results[rand.Intn(len-1)]
	}
}
