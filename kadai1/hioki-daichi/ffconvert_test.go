package main

import "testing"

func TestFFConvert(t *testing.T) {
	force := true
	verbose := true
	ok := execute([]string{"test/images"}, Jpeg, Png, force, verbose)
	if !ok {
		t.FailNow()
	}
}
