package main

import "testing"

func TestFFConvert(t *testing.T) {
	ok := execute([]string{"test/images"})
	if !ok {
		t.FailNow()
	}
}
