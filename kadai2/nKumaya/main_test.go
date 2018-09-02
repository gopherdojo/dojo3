package main

import (
	"os"
	"strings"
	"testing"
)

var cli = &CLI{os.Stdout, os.Stdout}

func TestRunFlagParse(t *testing.T) {

	t.Run("Flag option error", func(t *testing.T) {
		args := strings.Split("./main -m hoge images", " ")
		errNum := cli.Execute(args)
		if errNum != ExitCodeParseFlagError {
			t.Error(errNum)
		}
	})
	t.Run("Unspecified directory", func(t *testing.T) {
		args := strings.Split("./main", " ")
		errNum := cli.Execute(args)
		if errNum != ExitCodeParseArgError {
			t.Error(errNum)
		}
	})
}

func TestCreateFormat(t *testing.T) {
	args := strings.Split("./main -f testS -t testD images", " ")
	errNum := cli.Execute(args)
	if errNum != ExitCodeFormatError {
		t.Error(errNum)
	}

}

func TestRunDirectory(t *testing.T) {
	args := strings.Split("./main hogeDir", " ")
	errNum := cli.Execute(args)
	if errNum != ExitCodeConvertError {
		t.Error(errNum)
	}
}

func TestRun(t *testing.T) {
	argsPatterns := []string{
		"./main images",
		"./main -f jpg -t png images",
		"./main -f jpeg -t png images",
		"./main -f png -t jpg images",
		"./main -f png -t jpeg images",
		"./main -f jpg -t gif images",
		"./main -f jpeg -t gif images",
		"./main -f gif -t jpg images",
		"./main -f gif -t jpeg images",
		"./main -f png -t gif images",
		"./main -f gif -t png images",
	}
	for _, argsPattern := range argsPatterns {
		args := strings.Split(argsPattern, " ")
		code := cli.Execute(args)
		if code != ExitCodeOK {
			t.Error(code)
		}
	}
}
