package typing

import "fmt"

type stat struct {
	s int
	f int
}

func (s *stat) Succeed() {
	s.s++
}

func (s *stat) Fail() {
	s.f++
}

func (s *stat) Count() int {
	return s.s + s.f
}

func (s *stat) successRate() float64 {
	if s.Count() == 0 {
		return 0
	}
	return float64(s.s) / float64(s.Count())
}

func (s *stat) String() string {
	return "--- Result ---\n" +
		fmt.Sprintf("Typed %d words\n", s.s) +
		fmt.Sprintf("Succeed Rate %.1f%%\n", 100*s.successRate()) +
		"--------------"
}
