package gif

import (
	"image"
	"image/gif"
	"io"

	"github.com/gopherdojo/dojo3/kadai2/shuntaka9576/imagetypes"
)

func init() {
	init := &Gif{[]string{".gif"}}
	imagetypes.ResisterImageType(init)
}

type Gif struct {
	extStrs []string
}

func (*Gif) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

func (*Gif) Encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}

func (g *Gif) CheckExtStr(checkExt string) bool {
	for _, ext := range g.extStrs {
		if ext == checkExt {
			return true
			break
		}
	}
	return false
}
