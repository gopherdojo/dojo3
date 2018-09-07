package typing_lesson

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type TypingLesson struct {
	Words  []string
	Reader io.Reader
	Writer io.Writer
}

func NewTypingLesson(dict, input io.Reader, output io.Writer) (*TypingLesson, error) {
	rand.Seed(time.Now().Unix())

	words, err := loadWords(dict)
	if err != nil {
		return nil, err
	}

	return &TypingLesson{
		Words:  words,
		Reader: input,
		Writer: output,
	}, nil
}

func loadWords(dict io.Reader) ([]string, error) {
	// FIXME: スライスの初期化サイズうまく推測できないか？もしくは一括で読み込んでからSplitかけたほうが効率的？
	var words []string
	scanner := bufio.NewScanner(dict)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}

func (t *TypingLesson) Start(limit time.Duration) (ok int, ng int) {
	ch := input(t.Reader)
	timeout := timer(limit)
	for {
		word := t.randomWord()
		t.printf("%s\n>", word)

		select {
		case in := <-ch:
			if word == in {
				ok++
				t.printf("...OK\n")
			} else {
				ng++
				t.printf("...NG\n")
			}
			continue
		case <-timeout:
			t.printf("\n==========timed out==========\n")
			break
		}
		break
	}
	return ok, ng
}

func (t *TypingLesson) randomWord() string {
	return t.Words[rand.Intn(len(t.Words))]
}

func (t *TypingLesson) printf(format string, a ...interface{}) {
	fmt.Fprintf(t.Writer, format, a...)
}

func timer(limit time.Duration) <-chan bool {
	timeout := make(chan bool)
	go func() {
		time.Sleep(limit)
		timeout <- true
	}()
	return timeout
}

func input(r io.Reader) <-chan string {
	in := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			in <- s.Text()
		}
		close(in)
	}()
	return in
}
