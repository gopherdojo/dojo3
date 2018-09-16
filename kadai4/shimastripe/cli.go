package shimastripe

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

const success = iota

type CLI struct {
	Clock Clock
}

// Generate a seed only once
func (c *CLI) generateSeed() {
	rand.Seed(c.Clock.Now().Unix())
}

// handler
func (c *CLI) handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	respBody := &FortuneRepository{Fortune: DrawRandomly()}

	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		log.Printf("Encode error: %v\n", err)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
	}
}

// Run a server
func (c *CLI) Run(args []string) int {
	c.generateSeed()
	http.HandleFunc("/", c.handler)
	http.ListenAndServe(":8080", nil)
	return success
}
