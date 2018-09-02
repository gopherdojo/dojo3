package jpg

import (
	"image"
	"image/jpeg"
	"io"

	"github.com/gopherdojo/dojo3/kadai2/shuntaka9576/imagetypes"
)

func init() {
	init := &Jpeg{[]string{".jpeg", ".jpg"}}
	imagetypes.ResisterImageType(init)
}

type Jpeg struct {
	extStrs []string
}

func (*Jpeg) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

func (*Jpeg) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}

func (g *Jpeg) CheckExtStr(checkExt string) bool {
	for _, ext := range g.extStrs {
		if ext == checkExt {
			return true
			break
		}
	}
	return false
}
