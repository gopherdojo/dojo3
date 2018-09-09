package color

type Color int

const (
	Red = iota
	Green
	Yellow
	Blue
)

func (c Color) Code() (code string) {
	switch c {
	case Red:
		code = "31"
	case Green:
		code = "32"
	case Yellow:
		code = "33"
	case Blue:
		code = "34"
	}
	return
}
