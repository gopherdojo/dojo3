package jsonhelper

import (
	"errors"
	"io"
	"testing"
)

func TestJsonhelper_ToJSON(t *testing.T) {
	foo := struct {
		Bar string `json:"bar"`
		Baz int    `json:"baz"`
	}{
		Bar: "barbar",
		Baz: 1,
	}

	actual, err := ToJSON(foo)
	if err != nil {
		t.Fatalf("err %s", err)
	}
	expected := "{\"bar\":\"barbar\",\"baz\":1}\n"
	if actual != expected {
		t.Errorf(`unexpected : expected: "%s" actual: "%s"`, expected, actual)
	}
}

func TestJsonhelper_ToJSON_Error(t *testing.T) {
	expected := errInMock

	newEncoderFunc = func(w io.Writer) encoder { return &mockEncoder{} }

	_, actual := ToJSON(struct{}{})
	if actual != expected {
		t.Errorf(`unexpected : expected: "%s" actual: "%s"`, expected, actual)
	}
}

var errInMock = errors.New("error in mock")

type mockEncoder struct{}

func (m *mockEncoder) Encode(v interface{}) error { return errInMock }
