package download

import (
	"io"
	"fmt"
)

type Position struct {
	Range  Range
	Offset int64
}

func (p *Position) Absolute() int64 {
	return p.Range.Start + p.Offset
}

func (p *Position) Forward(n int64) {
	p.Offset += n
}

func (p *Position) CanForward(n int64) bool {
	return p.Absolute()+n-1 <= p.Range.End
}

type RangeWriter struct {
	io.WriterAt
	position Position
}

func newRangeWriter(w io.WriterAt, r Range) *RangeWriter {
	return &RangeWriter{w, Position{r, 0}}
}

func (w *RangeWriter) Write(p []byte) (int, error) {
	if !w.position.CanForward(int64(len(p))) {
		return 0, fmt.Errorf("Write position exceeds the range: len(p)=%d, position=%+v", len(p), w.position)
	}
	n, err := w.WriterAt, w.position.Absolute()
	w.position.Forward(int64(n))
	return n, err
}
