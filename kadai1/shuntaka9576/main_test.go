package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo3/kadai1/shuntaka9576/cli"
	"github.com/gopherdojo/dojo3/kadai1/shuntaka9576/convert"
	"time"
)

func TestCli(t *testing.T) {
	var tests = []struct {
		output   string
		args     []string
		expected string
	}{
		{"standard", []string{"imageConverter", "-version"}, fmt.Sprintf("imageConverter version %s", cli.Version)},
		{"standard", []string{"imageConverter", "-f", "png", "-t", "jpg", "."}, fmt.Sprint("Start Convert png to jpg[.]!\n")},
		{"standard", []string{"imageConverter", "-f", "jpg", "-t", "png", "."}, fmt.Sprint("Start Convert jpg to png[.]!\n")},
		{"standard", []string{"imageConverter", "-f", "jpg", "-t", "png", "./testdata"}, fmt.Sprint("Start Convert jpg to png[./testdata]!\n")},
		{"standard", []string{"imageConverter", "-f", "jpg", "-t", "png", `C:\Users\hozi576\go\src\github.com\gopherdojo\dojo3\kadai1\shuntaka9576\testdata`}, fmt.Sprint("Start Convert jpg to png[C:\\Users\\hozi576\\go\\src\\github.com\\gopherdojo\\dojo3\\kadai1\\shuntaka9576\\testdata]!\n")},
		{"standard", []string{"imageConverter", "-f", "png", "-t", "jpg", `C:\Users\hozi576\go\src\github.com\gopherdojo\dojo3\kadai1\shuntaka9576\testdata`}, fmt.Sprint("Start Convert png to jpg[C:\\Users\\hozi576\\go\\src\\github.com\\gopherdojo\\dojo3\\kadai1\\shuntaka9576\\testdata]!\n")},
		{"error", []string{"imageConverter", "-f", "jpg", "-t", "a", `C:\Users\hozi576\go\src\github.com\gopherdojo\dojo3\kadai1\shuntaka9576\testdata`}, fmt.Sprint("Validation check Error. Please check your from or to options\n")},
	}

	for _, test := range tests {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cliProcess := &cli.CLI{outStream, errStream}
		option, _ := cliProcess.GetOption(test.args)

		switch test.output {
		case "standard":
			if !strings.Contains(outStream.String(), test.expected) {
				t.Errorf("Output=%q, want %q", outStream.String(), test.expected)
			}
			convert.ConvertImagesProcess(option)
		case "error":
			if !strings.Contains(errStream.String(), test.expected) {
				t.Errorf("Output=%q, want %q", outStream.String(), test.expected)
			}
		}
		time.Sleep(1 * time.Second)
	}
}
