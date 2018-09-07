package color

import (
	"testing"
)

func TestColor_Code(t *testing.T) {
	cases := map[string]struct {
		clr      Color
		expected string
	}{
		"Green": {clr: Green, expected: "32"},
		"Red":   {clr: Red, expected: "31"},
		"Cyan":  {clr: Cyan, expected: "36"},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			expected := c.expected

			actual := c.clr.Code()
			if actual != expected {
				t.Errorf(`expected="%s" actual="%s"`, expected, actual)
			}
		})
	}
}

func TestColor_Code_Default(t *testing.T) {
	var c Color = 127
	defer func() {
		if recover() == nil {
			t.Errorf("did not panic")
		}
	}()
	c.Code()
}
