package word

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Type game word interface
type Word interface {
	Generate() error
	GetWord() (string, error)
}

type wordFile struct {
	path  string
	words []string
}

// Create word interface
func NewWordFile(path string) Word {
	return &wordFile{path: path}
}

// Generate type game word for file
func (w *wordFile) Generate() error {
	words := make([]string, 0)
	file, err := os.Open(w.path)
	if err != nil {
		return err
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		words = append(words, scan.Text())
	}

	if err := scan.Err(); err != nil {
		return err
	}

	w.words = words
	return nil
}

// Get word
func (w *wordFile) GetWord() (string, error) {
	len := len(w.words)
	if len == 0 {
		return "", errors.New("word was not found")
	}
	return w.words[rand.Intn(len-1)], nil
}
