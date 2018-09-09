package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestTimeUp(t *testing.T) {
	ch := make(chan struct{}, 1)
	TimeUp(0, ch)
	if _, closed := <-ch; closed {
		t.Error("channel should be closed")
	}
}

func TestQuestion(t *testing.T) {
	update := make(chan bool, 1)
	update <- true
	output := new(bytes.Buffer)
	Question(update, output)

	o := output.String()
	for _, d := range dictionary {
		if strings.Compare(d+"\n", o) == 0 {
			return
		}
	}
	t.Errorf("output should be included in dictionary. output is %v", o)
}

func TestRun(t *testing.T) {
	org := dictionary
	defer func() {
		dictionary = org
	}()
	dictionary = []string{
		"test",
	}

	cases := map[string]struct {
		expect string
		score  int
	}{
		"success": {
			"ok",
			1,
		},
		"failure": {
			"ng",
			0,
		},
	}

	for c := range cases {
		c := c
		input := new(bytes.Buffer)
		output := new(bytes.Buffer)

		t.Run(c, func(t *testing.T) {
			var err error
			if cases[c].score == 1 {
				_, err = input.WriteString("test")
			} else {
				_, err = input.WriteString("wrong")
			}
			if err != nil {
				t.Fatal("should not have error")
			}
			score := <-Run(input, output)

			out := output.String()
			if strings.Index(out, cases[c].expect) < 0 {
				t.Errorf("expect=%v, actual=%v", cases[c].expect, out)
			}
			if score != cases[c].score {
				println(cases[c].score, score)
				t.Errorf("expect=%v, actual=%v", cases[c].score, score)
			}
		})
	}
}
