package fileutil

import (
	"bytes"
	"os"
)

// StartsContentsWith returns whether file contents start with specified bytes.
func StartsContentsWith(fp *os.File, xs []uint8) bool {
	buf := make([]byte, len(xs))
	fp.Seek(0, 0)
	fp.Read(buf)
	fp.Seek(0, 0)
	return bytes.Equal(buf, xs)
}
