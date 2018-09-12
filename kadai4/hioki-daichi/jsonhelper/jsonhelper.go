/*
Package jsonhelper is a collection of convenient functions for manipulating JSON.
*/
package jsonhelper

import (
	"bytes"
	"encoding/json"
	"io"
)

// ToJSON converts v to JSON.
func ToJSON(v interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := newEncoderFunc(&buf)
	if err := encoder.Encode(v); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// for testing
type encoder interface {
	Encode(v interface{}) error
}

// for testing
var newEncoderFunc = func(w io.Writer) encoder {
	return json.NewEncoder(w)
}
