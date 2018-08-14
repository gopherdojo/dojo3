package main

import "testing"

func TestFFConvert(t *testing.T) {
	force := true
	ok := execute([]string{"test/images"}, Jpeg, Png, force)
	if !ok {
		t.FailNow()
	}
}
