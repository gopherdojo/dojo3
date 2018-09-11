package opt

import (
	"reflect"
	"testing"
)

func TestOpt_Parse(t *testing.T) {
	cases := map[string]struct {
		args            []string
		expectedOptions *Options
	}{
		"no options":     {args: []string{}, expectedOptions: &Options{Timeout: 15, Path: "./weapons.txt"}},
		"--timeout=1":    {args: []string{"--timeout=1"}, expectedOptions: &Options{Timeout: 1, Path: "./weapons.txt"}},
		"--path=foo.txt": {args: []string{"--path=foo.txt"}, expectedOptions: &Options{Timeout: 15, Path: "foo.txt"}},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			expected := c.expectedOptions

			actual := Parse(c.args...)
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf(`expected="%v" actual="%v"`, expected, actual)
			}
		})

	}
}
