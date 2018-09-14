package kuji_test

import (
	"testing"

	"github.com/gopherdojo/dojo3/kadai4/shuntaka9576/kuji"
)

func TestKuji_PrintFortune(t *testing.T) {
	tests := map[string]struct {
		fortune kuji.Kuji

		msg string
	}{
		"大吉": {
			fortune: kuji.Dikichi,
			msg:     "大吉",
		},
		"吉": {
			fortune: kuji.Kichi,
			msg:     "吉",
		},
		"中吉": {
			fortune: kuji.Chuukichi,
			msg:     "中吉",
		},
		"小吉": {
			fortune: kuji.Syoukichi,
			msg:     "小吉",
		},
		"末吉": {
			fortune: kuji.Suekichi,
			msg:     "末吉",
		},
		"凶": {
			fortune: kuji.Kyou,
			msg:     "凶",
		},
		"大凶": {
			fortune: kuji.Daikyou,
			msg:     "大凶",
		},
		"no": {
			fortune: 8,
			msg:     "Error",
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if tt.fortune.PrintFortune() != tt.msg {
				t.Errorf("kuji msgs wont %s but got %s\n", tt.msg, tt.fortune.PrintFortune())
			}
		})
	}
}

func TestRandomFortuneExpected(t *testing.T) {
	tests := map[string]struct {
		fortune []kuji.Kuji

		msgs []string
	}{
		"大吉/小吉": {
			fortune: []kuji.Kuji{
				kuji.Dikichi,
				kuji.Syoukichi,
			},
			msgs: []string{
				"大吉",
				"小吉",
			},
		},
		"大吉/小吉/中吉": {
			fortune: []kuji.Kuji{
				kuji.Dikichi,
				kuji.Syoukichi,
				kuji.Chuukichi,
			},
			msgs: []string{
				"大吉",
				"小吉",
				"中吉",
			},
		},
		"大吉/小吉/中吉/吉": {
			fortune: []kuji.Kuji{
				kuji.Dikichi,
				kuji.Syoukichi,
				kuji.Chuukichi,
				kuji.Kichi,
			},
			msgs: []string{
				"大吉",
				"小吉",
				"中吉",
				"吉",
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			testflag := false
			rfortune := kuji.RandomFortuneExpected(tt.fortune)
			for _, msg := range tt.msgs {
				if rfortune == msg {
					testflag = true
				}
			}
			if !testflag {
				t.Errorf("kuji msgs wont %s but got %s\n", tt.msgs, kuji.RandomFortuneExpected(tt.fortune))
			}
		})
	}
}
