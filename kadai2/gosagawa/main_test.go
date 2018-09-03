package main

import (
	"testing"
)

type AssertFn func(name string, err error)

func TestIsValidInput(t *testing.T) {

	noError := func(name string, err error) {
		if err != nil {
			t.Errorf("%v: expected no error", name)
		}
	}
	withError := func(name string, err error) {
		if err == nil {
			t.Errorf("%v: expected returning error", name)
		}
	}

	nilArgs := []string{}
	correctArgs := []string{"test"}
	overArgs := []string{"test1", "test2"}

	cases := []struct {
		name     string
		inType   string
		outType  string
		args     []string
		assertFn AssertFn
	}{
		{name: "all correct", inType: "jpeg", outType: "png", args: correctArgs, assertFn: noError},
		{name: "no args", inType: "jpeg", outType: "png", args: nilArgs, assertFn: withError},
		{name: "more than two args", inType: "jpeg", outType: "png", args: overArgs, assertFn: withError},
		{name: "unsupported inType", inType: "unsuppport", outType: "png", args: correctArgs, assertFn: withError},
		{name: "unsupported outType", inType: "png", outType: "unsupport", args: correctArgs, assertFn: withError},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			err := IsValidInput(v.inType, v.outType, v.args)
			v.assertFn(v.name, err)
		})
	}
}
