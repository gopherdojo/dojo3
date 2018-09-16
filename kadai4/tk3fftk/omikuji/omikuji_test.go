package omikuji

import (
	"strings"
	"testing"
)

func TestOmikuji_Do(t *testing.T) {
	cases := map[string][]string{
		"default": defaultLots,
		"one":     {"大吉"},
		"two":     {"大吉", "大凶"},
	}

	o := NewOmikuji()

	for c := range cases {
		c := c
		t.Run(c, func(t *testing.T) {
			t.Helper()

			lots := cases[c]
			lot := o.Do(lots)
			for _, l := range lots {
				if strings.Compare(lot, l) == 0 {
					return
				}
			}
			t.Errorf("%v should be included in %v", lot, lots)
		})
	}
}
