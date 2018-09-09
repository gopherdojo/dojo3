package words_test

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo3/kadai3-1/shuntaka9576/words"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		testtype string
		input    io.Reader
		want     words.Words
	}{
		"normal-case": {"normal", strings.NewReader("abc\nabcd"), words.Words{words.Word{"abc", "abc"}, words.Word{"abcd", "abcd"}}},
		"normal":      {"non-normal", strings.NewReader(""), words.Words{words.Word{"abc", "abc"}, words.Word{"abcd", "abcd"}}},
	}

	for name, tt := range tests {
		if tt.testtype == "normal" {
			t.Run(name, func(t *testing.T) {
				wlist, err := words.New(tt.input)
				if err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(wlist, tt.want) {
					t.Errorf(`actual="%v" want="%v"`, tt.input, tt.want)
				}
			})
		} else {
		}
	}
}
