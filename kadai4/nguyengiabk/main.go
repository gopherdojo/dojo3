package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gopherdojo/dojo3/kadai4/nguyengiabk/omikuji"
)

// Result represents api response structure
type Result struct {
	Result omikuji.Fortune `json:"result"`
}

// JSONEncoder is used to encode JSON
type JSONEncoder interface {
	Encode(interface{}, io.Writer) error
}

// JSONEncodeFunc is function that used to encode JSON
type JSONEncodeFunc func(interface{}, io.Writer) error

// Encode encode JSON and write result to writer
func (f JSONEncodeFunc) Encode(data interface{}, w io.Writer) error {
	return f(data, w)
}

// OmikujiServer handles request and response as JSON
type OmikujiServer struct {
	jsonEncoder JSONEncoder
	omikuji     omikuji.Omikuji
}

func (os *OmikujiServer) encode(data interface{}, w io.Writer) error {
	if os.jsonEncoder != nil {
		return os.jsonEncoder.Encode(data, w)
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}

// ServerErrorMessage is message that will be returned to user in case of error
const serverErrorMessage = "Server error"

func (os *OmikujiServer) handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := Result{os.omikuji.GetResult()}
	if err := os.encode(res, w); err != nil {
		http.Error(w, serverErrorMessage, http.StatusInternalServerError)
	}
}

func main() {
	os := OmikujiServer{omikuji: omikuji.Omikuji{}}
	http.HandleFunc("/", os.handler)
	http.ListenAndServe(":8080", nil)
}
