package fortune

import (
	"math/rand"
	"testing"
)

func TestFortune_DrawFortune(t *testing.T) {
	cases := map[string]struct {
		seed     int64
		expected Fortune
	}{
		"KYOU":     {seed: 0, expected: Kyo},
		"DAIKYOU":  {seed: 1, expected: Daikyo},
		"SUEKICHI": {seed: 2, expected: Suekichi},
		"KICHI":    {seed: 3, expected: Kichi},
		"CHUKICHI": {seed: 4, expected: Chukichi},
		"SHOKICHI": {seed: 5, expected: Shokichi},
		"DAICHIKI": {seed: 9, expected: Daikichi},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			rand.Seed(c.seed)

			expected := c.expected
			actual := DrawFortune()
			if actual != expected {
				t.Errorf(`unexpected response body: expected: "%s" actual: "%s"`, expected, actual)
			}
		})
	}
}
