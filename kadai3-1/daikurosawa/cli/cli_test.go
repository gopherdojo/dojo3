package cli_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/gopherdojo/dojo3/kadai3-1/daikurosawa/cli"
)

type Mock struct{}

func (Mock) Generate() error {
	return nil
}

func (Mock) GetWord() (string, error) {
	return "dummy", nil
}

func TestCLI_Play(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	ch := make(chan string)
	c := cli.NewExportCLI(inStream, outStream, errStream, Mock{}, ch)

	go func() {
		ch <- "dummy"
	}()

	err := cli.ExportPlay(c, "dummy", 1*time.Second)
	if err != nil {
		t.Error("failed type game play.", err)
	}
	output := "> dummy\n" + "\x1b[32mSuccess!\x1b[0m\n" + "> dummy\n" + "Time out!\n" + "Result: 1\n"
	if output != outStream.String() {
		t.Fatalf("output string different. output: %s", outStream.String())
	}
	if len(errStream.String()) > 0 {
		t.Fatalf("output err message. error: %s", errStream.String())
	}
}
