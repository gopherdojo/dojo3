package parser

import (
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/decoder"
	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/encoder"
)

func TestParse(t *testing.T) {
	cases := map[string]struct {
		osargs          []string
		expectedDir     string
		expectedDecoder decoder.Decoder
		expectedEncoder encoder.Encoder
		expectedErr     error
	}{
		"empty args": {
			osargs:          []string{"cmd"},
			expectedDir:     "./",
			expectedDecoder: &decoder.Jpg{},
			expectedEncoder: &encoder.Png{},
			expectedErr:     nil,
		},
		// with full args
		// jpg => png, png => jpg
		// with err
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			a, err := Parse(c.osargs)

			if a.Dir != c.expectedDir {
				t.Errorf("parser.Parse wont %s but got %s", c.expectedDir, a.Dir)
			}

			if a.Decoder != c.expectedDecoder {
				t.Errorf("parser.Parse wont %s but got %s", c.expectedDecoder, a.Decoder)
			}

			if a.Encoder != c.expectedEncoder {
				t.Errorf("parser.Parse wont %s but got %s", c.expectedEncoder, a.Encoder)
			}

			if err != c.expectedErr {
				t.Errorf("parser.Parse wont %s but got %s", c.expectedErr, err)
			}
		})
	}
}
