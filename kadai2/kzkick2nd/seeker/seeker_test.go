package seeker

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSeek(t *testing.T) {

	dir, err := ioutil.TempDir("", "tempdir")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: 複雑なディレクトリ構成を生成したい
	content := []byte("temporary file's content")
	tmpfn := filepath.Join(dir, "tmpfile.jpg")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}

	cases := map[string]struct {
		dir           string
		ext           string
		expectedPaths []string
		expectedErr   error
	}{
		"default option": {
			dir:           dir,
			ext:           "jpg",
			expectedPaths: []string{tmpfn},
			expectedErr:   nil,
		},
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

	// FIXME: 前後処理挟み込んでいる事を表現したい
	os.RemoveAll(dir)
}
