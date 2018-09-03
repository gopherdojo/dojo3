package main

import (
	"testing"
)

type AssertFn func(err error)

func TestIsValidInput(t *testing.T) {

	noError := func(err error) {
		if err != nil {
			t.Errorf("expected no error")
		}
	}
	withError := func(err error) {
		if err == nil {
			t.Errorf("expected returning error")
		}
	}

	nilArgs := []string{}
	correctArgs := []string{"test"}
	overArgs := []string{"test1", "test2"}

	cases := []struct {
		inType   string
		outType  string
		args     []string
		assertFn AssertFn
	}{
		{inType: "jpeg", outType: "png", args: correctArgs, assertFn: noError},
		{inType: "jpeg", outType: "png", args: nilArgs, assertFn: withError},
		{inType: "jpeg", outType: "png", args: overArgs, assertFn: withError},
		{inType: "unsuppport", outType: "png", args: correctArgs, assertFn: withError},
		{inType: "png", outType: "unsupport", args: correctArgs, assertFn: withError},
	}

	for _, v := range cases {
		err := IsValidInput(v.inType, v.outType, v.args)
		v.assertFn(err)
	}

}
