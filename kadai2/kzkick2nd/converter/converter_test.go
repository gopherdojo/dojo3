package converter

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/decoder"
	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/encoder"
)

func TestRun(t *testing.T) {

	// TODO: 複雑なディレクトリ構成を生成したい
	dir, err := ioutil.TempDir("", "tempdir")
	if err != nil {
		log.Fatal(err)
	}

	src := "../testdata/1px.jpg"
	dest := filepath.Join(dir, "tmpfile.png")

	input, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		log.Fatal(err)
	}

	cases := map[string]struct {
		log          io.Writer
		src          []string
		encoder      encoder.Encoder
		decoder      decoder.Decoder
		expectedErr  error
		expectedDest []string
	}{
		"starndard": {
			log:          &bytes.Buffer{},
			src:          []string{src},
			decoder:      &decoder.Jpg{},
			encoder:      &encoder.Png{},
			expectedErr:  nil,
			expectedDest: []string{dest},
		},
	}

	for n, c := range cases {

		c := c
		t.Run(n, func(t *testing.T) {
			q := Queue{
				Log:     c.log,
				Src:     c.src,
				Encoder: c.encoder,
				Decoder: c.decoder,
			}
			err = q.Run()
			if err != c.expectedErr {
				t.Errorf("converter.Run() wont %s but got %s", c.expectedErr, err)
			}
		})
	}

	// FIXME: 前後処理挟み込んでいる事を表現したい
	os.RemoveAll(dir)
}
