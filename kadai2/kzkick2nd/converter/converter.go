/*
Package converter converts image. open and decode then encode.
*/
package converter

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/decoder"
	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/encoder"
)

// Queue of convert.
type Queue struct {
	Log     io.Writer
	Src     []string
	Encoder encoder.Encoder
	Decoder decoder.Decoder
}

// Run convert.
func (q *Queue) Run() error {
	for i, _ := range q.Src {
		err := q.convert(i)
		if err != nil {
			return err
		}
	}
	return nil
}

// convert() contains Decode and Encode.
func (q *Queue) convert(i int) error {
	s := q.Src[i]
	f, err := os.Open(s)
	if err != nil {
		return err
	}
	defer f.Close()

	m, err := q.Decoder.Run(f)
	if err != nil {
		return err
	}

	p := s[:len(s)-len(filepath.Ext(s))] + "." + q.Encoder.Ext()
	d, err := os.Create(p)
	if err != nil {
		return err
	}

	err = q.Encoder.Run(d, m)
	if err != nil {
		return err
	}

	// FIXME: 出力ファイルのスライスを出した方が良い。テストも楽。
	fmt.Fprintf(q.Log, "converted: %s => %s\n", s, p)

	err = d.Close()
	if err != nil {
		return err
	}

	return nil
}
