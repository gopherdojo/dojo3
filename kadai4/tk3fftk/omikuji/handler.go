package omikuji

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const bigLuckey = "大吉"

type response struct {
	Result string `json:"result"`
}

type Clock interface {
	Now() time.Time
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}

type OmikujiHandler struct {
	Clock   Clock
	Omikuji Omikuji
}

func New(clock Clock) OmikujiHandler {
	o := NewOmikuji()
	return OmikujiHandler{
		Clock:   clock,
		Omikuji: o,
	}
}

func (o *OmikujiHandler) createResponseJSON(result string) (string, error) {
	res := &response{Result: result}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(res); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (o *OmikujiHandler) handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var res string
	var err error

	if o.isNewYear() {
		res, err = o.createResponseJSON(bigLuckey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		lot := o.Omikuji.Do([]string{})
		res, err = o.createResponseJSON(lot)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	fmt.Fprintf(w, res)
}

func (o *OmikujiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	o.handlerFunc(w, r)
}

func (o *OmikujiHandler) now() time.Time {
	if o.Clock == nil {
		return time.Now()
	}
	return o.Clock.Now()
}

func (o *OmikujiHandler) isNewYear() bool {
	_, m, d := o.now().Date()
	if m == 1 && d <= 3 {
		return true
	}
	return false
}
