package converter

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/decoder"
	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/encoder"
)

type Queue struct {
	Log     io.Writer
	Src     []string
	Encoder encoder.Encoder
	Decoder decoder.Decoder
}

func (q *Queue) Run() error {
	for i, _ := range q.Src {
		err := q.convert(i)
		if err != nil {
			return err
		}
	}
	return nil
}

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

	err = d.Close()
	if err != nil {
		return err
	}

	return nil
}
