package shimastripe

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const success = iota

type CLI struct {
}

func init() {
	rand.Seed(time.Now().Unix())
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	respBody := &FortuneRepository{Fortune: DrawRandomly()}

	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		log.Printf("Encode error: %v\n", err)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
	}
}

func (c *CLI) Run(args []string) int {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	return success
}
