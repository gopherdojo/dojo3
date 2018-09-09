package words

import (
	"bufio"
	"io"
)

type word struct {
	input, expected string
}

type words []word

func New(reader io.Reader) (wlist words, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		var w word
		w.input, w.expected = scanner.Text(), scanner.Text()
		wlist = append(wlist, w)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
}
