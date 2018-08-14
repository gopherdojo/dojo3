package main

import "testing"

func TestFFConvert(t *testing.T) {
	ok := execute([]string{"test/images"}, Jpeg, Png)
	if !ok {
		t.FailNow()
	}
}
