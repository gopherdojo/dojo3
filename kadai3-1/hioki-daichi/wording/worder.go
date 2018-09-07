package wording

import (
	"bufio"
	"os"
)

// Worder has a file path as a generation source.
type Worder struct {
	path string
}

// NewWorder generates a Worder.
func NewWorder(path string) *Worder {
	return &Worder{path: path}
}

// Words generates words from the file corresponding to their own path.
func (w *Worder) Words() ([]string, error) {
	words := make([]string, 0)

	fp, err := os.Open(w.path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}
