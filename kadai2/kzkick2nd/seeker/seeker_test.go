package seeker

import (
	"reflect"
	"testing"
)

func TestSeek(t *testing.T) {
	cases := map[string]struct {
		dir           string
		ext           string
		expectedPaths []string
		expectedErr   error
	}{
		"default option": {
			dir:           "../testdata",
			ext:           "jpg",
			expectedPaths: []string{"../testdata/1px.jpg", "../testdata/1px_2.jpg", "../testdata/subdir/sub-dir-1px.jpg"},
			expectedErr:   nil,
		},
		// TODO: create test temp dir on testing seek()
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			o := Target{
				Dir: c.dir,
				Ext: c.ext,
			}

			p, err := o.Seek()

			if !reflect.DeepEqual(p, c.expectedPaths) {
				t.Errorf("target.Seek() wont %s but got %s", c.expectedPaths, p)
			}

			if err != c.expectedErr {
				t.Errorf("target.Seek() wont %s but got %s", c.expectedErr, err)
			}

		})
	}
}
