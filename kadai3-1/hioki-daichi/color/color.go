package color

type Color int

const (
	Green = iota
	Red
	Cyan
)

// Code returns ANSI escape code.
// see https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
func (c Color) Code() string {
	switch c {
	case Green:
		return "32"
	case Red:
		return "31"
	case Cyan:
		return "36"
	default:
		panic("unknown")
	}
}
