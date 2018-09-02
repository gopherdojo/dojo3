package converter

import (
	"testing"
)

func TestPng_GetExt(t *testing.T) {
	tests := []struct {
		name string
		p    *Png
		want string
	}{
		{name: "success", want: "png"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Png{}
			if got := p.GetExt(); got != tt.want {
				t.Errorf("Png.GetExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
