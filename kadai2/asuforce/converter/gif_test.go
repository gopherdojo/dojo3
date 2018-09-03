package converter

import (
	"testing"
)

func TestGif_GetExt(t *testing.T) {
	tests := []struct {
		name string
		g    *Gif
		want string
	}{
		{name: "success", want: ".gif"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gif{}
			if got := g.GetExt(); got != tt.want {
				t.Errorf("Gif.GetExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
