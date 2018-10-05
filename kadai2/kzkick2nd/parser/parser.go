/*
Package parser is parsing Args for converter.
*/
package parser

import (
	"flag"
	"os"

	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/decoder"
	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/encoder"
)

// Args is parsed os.Args for convert source.
type Args struct {
	Dir     string
	Decoder decoder.Decoder
	Encoder encoder.Encoder
}

// Parse args with flag and contruct.
func Parse(s []string) (Args, error) {

	f := flag.NewFlagSet(s[0], flag.ExitOnError)
	from := f.String("i", "jpg", "convert from (jpg|png)")
	to := f.String("o", "png", "convert to (jpg|png)")
	f.Parse(s[1:])
	dir := f.Arg(0)

	return Args{
		Dir:     validDir(dir),
		Decoder: identifyDecoder(from),
		Encoder: identifyEncoder(to),
	}, nil
}

// validDir check dir exist.
func validDir(s string) string {
	if _, err := os.Stat(s); os.IsNotExist(err) {
		return "./"
	}
	return s
}

// identify Decoder.
func identifyDecoder(s *string) decoder.Decoder {
	switch *s {
	case "jpg", "jpeg":
		return &decoder.Jpg{}
	case "png":
		return &decoder.Png{}
	default:
		return &decoder.Jpg{}
	}
}

// identify Encoder.
func identifyEncoder(s *string) encoder.Encoder {
	switch *s {
	case "jpg", "jpeg":
		return &encoder.Jpg{}
	case "png":
		return &encoder.Png{}
	default:
		return &encoder.Png{}
	}
}
