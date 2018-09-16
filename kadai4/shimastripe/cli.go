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

func (c *CLI) IsNewYear() bool {
	_, m, d := c.Clock.Now().Date()
	if m == 1 && d <= 3 {
		return true
	}
	return false
}

// handler
func (c *CLI) handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	var respBody *FortuneRepository

	if c.IsNewYear() {
		respBody = &FortuneRepository{Fortune: daikichi}
	} else {
		respBody = &FortuneRepository{Fortune: DrawRandomly()}
	}

	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		log.Printf("Encode error: %v\n", err)
		http.Error(w, "{\"result\":\"Internal server error\"}\n", http.StatusInternalServerError)
	}
}

// Run a server
func (c *CLI) Run(args []string) int {
	c.generateSeed()
	http.HandleFunc("/", c.handler)
	http.ListenAndServe(":8080", nil)
	return success
}
