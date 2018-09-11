package wording

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestWording_Words(t *testing.T) {
	w := NewWorder("./testdata/words.txt")
	actual, err := w.Words()
	if err != nil {
		t.Fatalf("err %s", err)
	}
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestWording_Words_Unopenable(t *testing.T) {
	dir, err := ioutil.TempDir("", "splatoon2-weapons-typing")
	if err != nil {
		t.Fatalf("err %s", err)
	}
	path := filepath.Join(dir, "unopenable.txt")

	expected := "open " + path + ": permission denied"

	_, err = os.OpenFile(path, os.O_CREATE, 000)
	if err != nil {
		t.Fatalf("err %s", err)
	}
	defer os.Remove(path)

	_, err = NewWorder(path).Words()

	actual := err.Error()
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}
