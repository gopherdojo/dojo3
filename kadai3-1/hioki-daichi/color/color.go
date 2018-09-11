package color

// Color deals with color.
type Color int

const (
	// Green represents green color.
	Green Color = 0

	// Red represents red color.
	Red Color = 1

	// Cyan represents cyan color.
	Cyan Color = 2
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
