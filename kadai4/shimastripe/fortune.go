package shimastripe

import "math/rand"

type FortuneElement int
type FortuneRepository struct {
	Fortune FortuneElement `json:"fortune"`
}

const (
	daikichi FortuneElement = iota
	chukichi
	shokichi
	suekichi
	kichi
	kyo
	daikyo
	threshold // use FortuneElement length
)

// Behave Stringer interface
func (oe FortuneElement) String() string {
	switch oe {
	case daikichi:
		return "大吉"
	case chukichi:
		return "中吉"
	case shokichi:
		return "小吉"
	case suekichi:
		return "末吉"
	case kichi:
		return "吉"
	case kyo:
		return "凶"
	case daikyo:
		return "大凶"
	default:
		return "強運"
	}
}

func (oe FortuneElement) MarshalJSON() ([]byte, error) {
	return []byte(`"` + oe.String() + `"`), nil
}

// Draw a fortune randomly
func DrawRandomly() FortuneElement {
	return FortuneElement(rand.Intn(int(threshold)))
}
