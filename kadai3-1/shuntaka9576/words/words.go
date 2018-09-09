package words

import (
	"bufio"
	"io"
)

type Word struct {
	Input, Expected string
}

type Words []Word

func New(reader io.Reader) (wlist Words, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		var w Word
		w.Input, w.Expected = scanner.Text(), scanner.Text()
		wlist = append(wlist, w)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return
}
