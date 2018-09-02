package cli_test

import (
	"bytes"
	"strings"
	"testing"

	c "github.com/gopherdojo/dojo3/kadai2/daikurosawa/cli"
	_ "github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/gif"
	_ "github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/jpg"
	_ "github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/png"
)

func TestCLI_Run(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &c.CLI{OutStream: outStream, ErrStream: errStream}
	args := strings.Split("convert ./../testdata", " ")
	exitCode := cli.Run(args)

	if exitCode != c.ExitCodeOK {
		t.Errorf("failed cli run, exit_code: %d", exitCode)
	}

	if errStream.Len() > 0 {
		t.Errorf("failed cli run, output: %q", errStream.String())
	}
}

func TestCLI_Run_ParseError(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &c.CLI{OutStream: outStream, ErrStream: errStream}
	args := strings.Split("convert -foo", " ") // undefined option
	exitCode := cli.Run(args)

	if exitCode != c.ExitCodeParseFlagError {
		t.Errorf("failed cli run, exit_code: %d", exitCode)
	}

	if errStream.Len() == 0 {
		t.Errorf("failed error message is not output")
	}
}

func TestCLI_Run_ProcessError_NotExistDirectory(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &c.CLI{OutStream: outStream, ErrStream: errStream}
	args := strings.Split("convert ./foo", " ")
	exitCode := cli.Run(args)

	if exitCode != c.ExitCodeInvalidArgsError {
		t.Errorf("failed cli run, exit_code: %d", exitCode)
	}

	if errStream.Len() == 0 {
		t.Errorf("failed error message is not output")
	}
}

func TestCLI_Run_ProcessError_NotDirectory(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &c.CLI{OutStream: outStream, ErrStream: errStream}
	args := strings.Split("convert ./../testdata/gopher.jpg", " ")
	exitCode := cli.Run(args)

	if exitCode != c.ExitCodeInvalidArgsError {
		t.Errorf("failed cli run, exit_code: %d", exitCode)
	}

	if errStream.Len() == 0 {
		t.Errorf("failed error message is not output")
	}
}
