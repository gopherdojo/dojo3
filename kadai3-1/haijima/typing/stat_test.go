package typing

import (
	"testing"
)

func Test_stat_Succeed(t *testing.T) {
	t.Run("Initial value should be zero value", func(t *testing.T) {
		s := stat{}
		if got := s.s; got != 0 {
			t.Errorf("stat.s expected to be 0, but: %d", got)
		}
	})

	t.Run("Succeed() should count up success count by 1", func(t *testing.T) {
		s := stat{}
		s.Succeed()
		if got := s.s; got != 1 {
			t.Errorf("stat.s expected to be 1, but: %d", got)
		}
	})

	t.Run("Succeed() should also count up total count by 1", func(t *testing.T) {
		s := stat{}
		s.Succeed()
		if got := s.Count(); got != 1 {
			t.Errorf("stat.Count() expected to be 1, but: %d", got)
		}
	})

	t.Run("Succeed() should not change failure value", func(t *testing.T) {
		s := stat{}
		b := s.f
		s.Succeed()
		a := s.f
		if b != a {
			t.Errorf("stat.f has changed from %d to %d", b, a)
		}
	})
}

func Test_stat_Fail(t *testing.T) {
	t.Run("Initial value should be zero value", func(t *testing.T) {
		s := stat{}
		if got := s.f; got != 0 {
			t.Errorf("stat.f expected to be 0, but: %d", got)
		}
	})

	t.Run("Fail() should count up failure count by 1", func(t *testing.T) {
		s := stat{}
		s.Fail()
		if got := s.f; got != 1 {
			t.Errorf("stat.f expected to be 1, but: %d", got)
		}
	})

	t.Run("Fail() should also count up total count by 1", func(t *testing.T) {
		s := stat{}
		s.Fail()
		if got := s.Count(); got != 1 {
			t.Errorf("stat.Count() expected to be 1, but: %d", got)
		}
	})

	t.Run("Fail() should not change success value", func(t *testing.T) {
		s := stat{}
		b := s.s
		s.Fail()
		a := s.s
		if b != a {
			t.Errorf("stat.s has changed from %d to %d", b, a)
		}
	})
}

func Test_stat_Count(t *testing.T) {
	t.Run("Initial value should be zero value", func(t *testing.T) {
		s := stat{}
		if got := s.Count(); got != 0 {
			t.Errorf("stat.Count() expected to be 0, but: %d", got)
		}
	})
}

func Test_stat_successRate(t *testing.T) {
	type fields struct {
		s int
		f int
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{name: "normal case", fields: fields{s: 1, f: 2}, want: float64(1) / 3},
		{name: "zero division", fields: fields{s: 0, f: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stat{
				s: tt.fields.s,
				f: tt.fields.f,
			}
			if got := s.successRate(); got != tt.want {
				t.Errorf("stat.successRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stat_String(t *testing.T) {
	s := stat{}
	s.Succeed()
	s.Succeed()
	s.Fail()
	want := "--- Result ---\n" +
		"Typed 2 words\n" +
		"Succeed Rate 66.7%\n" +
		"--------------"
	if got := s.String(); got != want {
		t.Errorf("stat.string() = %v, want %v", got, want)
	}
}
