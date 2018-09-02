package converter

import (
	"testing"
)

func TestJpg_GetExt(t *testing.T) {
	tests := []struct {
		name string
		j    *Jpg
		want string
	}{
		{name: "success", want: ".jpg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Jpg{}
			if got := j.GetExt(); got != tt.want {
				t.Errorf("Jpg.GetExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
