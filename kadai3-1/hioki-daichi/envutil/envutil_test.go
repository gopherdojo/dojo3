package envutil

import (
	"os"
	"testing"
)

func TestEnvutil_GetIntEnvOrElse(t *testing.T) {
	err := os.Setenv("FOO", "1")
	if err != nil {
		t.Fatalf("err %s", err)
	}

	expected := 1

	actual, err := GetIntEnvOrElse("FOO", 0)
	if err != nil {
		t.Fatalf("err %s", err)
	}

	if actual != expected {
		t.Errorf(`expected="%d" actual="%d"`, expected, actual)
	}
}

func TestEnvutil_getIntEnvOrElse_DefaultValue(t *testing.T) {
	expected := 99
	actual, err := GetIntEnvOrElse("NON_EXISTENT_ENV_VAR_NAME", expected)
	if err != nil {
		t.Fatalf("err %s", err)
	}
	if actual != expected {
		t.Errorf(`expected="%d" actual="%d"`, expected, actual)
	}
}

func TestEnvutil_getIntEnvOrElse_AtoiFailed(t *testing.T) {
	err := os.Setenv("FOO", "foo")
	if err != nil {
		t.Fatalf("err %s", err)
	}

	expected := `strconv.Atoi: parsing "foo": invalid syntax`

	_, err = GetIntEnvOrElse("FOO", 0)
	actual := err.Error()
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}
