package omikuji_test

import (
	"testing"
	"time"

	"github.com/gopherdojo/dojo3/kadai4/shuntaka9576/omikuji"
)

func mockClock(t *testing.T, v string) omikuji.Clock {
	t.Helper()
	now, err := time.Parse("2006/01/02 15:04:05", v)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	return omikuji.ClockFunc(func() time.Time {
		return now
	})
}

func TestOmikuji_Run(t *testing.T) {
	tests := map[string]struct {
		clock omikuji.Clock

		msg       string
		expectErr bool
	}{
		"1月1日": {
			clock:     mockClock(t, "2018/01/01 00:00:00"),
			msg:       "大吉",
			expectErr: true,
		},
		"1月2日": {
			clock:     mockClock(t, "2018/01/02 00:00:00"),
			msg:       "大吉",
			expectErr: true,
		},
		"1月3日": {
			clock:     mockClock(t, "2018/01/03 00:00:00"),
			msg:       "大吉",
			expectErr: true,
		},
		"1月4日": {
			clock:     mockClock(t, "2018/01/04 00:00:00"),
			msg:       "大吉",
			expectErr: false,
		},
		"1月5日": {
			clock:     mockClock(t, "2018/01/04 00:00:00"),
			msg:       "大吉",
			expectErr: false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			o := omikuji.Omikuji{
				Clock: tt.clock,
			}

			if o.Run() != tt.msg {
				if tt.expectErr {
					t.Errorf("omikuji msg wont %s but got %s\n", tt.msg, o.Run())
				}
			}
		})
	}
}
