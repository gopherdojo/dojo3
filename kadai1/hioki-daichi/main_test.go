package main

import (
	"bytes"
	"os/exec"
	"regexp"
	"testing"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/cliopt"
)

func TestJpegToPng(t *testing.T) {
	tmpdir := "testdata/" + "TestJpegToPng"

	exec.Command("cp", "-r", "testdata/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromJpeg = true
	cliopt.ToPng = true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestJpegToPng"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestJpegToPng/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestJpegToPng/2018/07"`)
	expectToMatchBuffer(t, buf, `Converted: "testdata/TestJpegToPng/2018/07/001.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestJpegToPng/2018/07/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestJpegToPng/2018/08"`)
	expectToMatchBuffer(t, buf, `Converted: "testdata/TestJpegToPng/2018/08/001.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestJpegToPng/2018/08/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestJpegToPng/2018/08/003.gif"`)
}

func TestJpegToGif(t *testing.T) {
	tmpdir := "testdata/" + "TestJpegToGif"

	exec.Command("cp", "-r", "testdata/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromJpeg = true
	cliopt.ToGif = true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestJpegToGif"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestJpegToGif/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestJpegToGif/2018/07"`)
	expectToMatchBuffer(t, buf, `Converted: "testdata/TestJpegToGif/2018/07/001.gif"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestJpegToGif/2018/07/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestJpegToGif/2018/08"`)
	expectToMatchBuffer(t, buf, `Converted: "testdata/TestJpegToGif/2018/08/001.gif"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestJpegToGif/2018/08/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestJpegToGif/2018/08/003.gif"`)
}

func TestPngToJpeg(t *testing.T) {
	tmpdir := "testdata/" + "TestPngToJpeg"

	exec.Command("cp", "-r", "testdata/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromPng = true
	cliopt.ToJpeg = true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestPngToJpeg"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestPngToJpeg/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestPngToJpeg/2018/07"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestPngToJpeg/2018/07/001.jpg"`)
	expectToMatchBuffer(t, buf, `Converted: "testdata/TestPngToJpeg/2018/07/002.jpg"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestPngToJpeg/2018/08"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestPngToJpeg/2018/08/001.jpeg"`)
	expectToMatchBuffer(t, buf, `Converted: "testdata/TestPngToJpeg/2018/08/002.jpg"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestPngToJpeg/2018/08/003.gif"`)
}

func TestPngToGif(t *testing.T) {
	tmpdir := "testdata/" + "TestPngToGif"

	exec.Command("cp", "-r", "testdata/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromPng = true
	cliopt.ToGif = true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestPngToGif"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestPngToGif/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestPngToGif/2018/07"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestPngToGif/2018/07/001.jpg"`)
	expectToMatchBuffer(t, buf, `Converted: "testdata/TestPngToGif/2018/07/002.gif"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestPngToGif/2018/08"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestPngToGif/2018/08/001.jpeg"`)
	expectToMatchBuffer(t, buf, `Converted: "testdata/TestPngToGif/2018/08/002.gif"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestPngToGif/2018/08/003.gif"`)
}

func TestGifToJpeg(t *testing.T) {
	tmpdir := "testdata/" + "TestGifToJpeg"

	exec.Command("cp", "-r", "testdata/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromGif = true
	cliopt.ToJpeg = true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestGifToJpeg"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestGifToJpeg/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestGifToJpeg/2018/07"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestGifToJpeg/2018/07/001.jpg"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestGifToJpeg/2018/07/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestGifToJpeg/2018/08"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestGifToJpeg/2018/08/001.jpeg"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestGifToJpeg/2018/08/002.png"`)
	expectToMatchBuffer(t, buf, `Converted: "testdata/TestGifToJpeg/2018/08/003.jpg"`)
}

func TestGifToPng(t *testing.T) {
	tmpdir := "testdata/" + "TestGifToPng"

	exec.Command("cp", "-r", "testdata/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromGif = true
	cliopt.ToPng = true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestGifToPng"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestGifToPng/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestGifToPng/2018/07"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestGifToPng/2018/07/001.jpg"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestGifToPng/2018/07/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "testdata/TestGifToPng/2018/08"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestGifToPng/2018/08/001.jpeg"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "testdata/TestGifToPng/2018/08/002.png"`)
	expectToMatchBuffer(t, buf, `Converted: "testdata/TestGifToPng/2018/08/003.png"`)
}

func TestConflict(t *testing.T) {
	tmpdir := "testdata/" + "TestConflict"

	exec.Command("cp", "-r", "testdata/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromJpeg = true
	cliopt.ToPng = true

	cliopt.Verbose = true
	cliopt.Force = true
	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.execute(tmpdir)

	cliopt.Verbose = true
	cliopt.Force = false
	cli = &CLI{OutStream: buf, ErrStream: buf}
	err := cli.execute(tmpdir)

	expected := "File already exists: testdata/TestConflict/2018/07/001.png"
	if err.Error() != expected {
		t.Errorf("expected: %s, actual: %s", expected, err)
	}
}

func initCliopt() {
	cliopt.FromJpeg = false
	cliopt.FromPng = false
	cliopt.FromGif = false
	cliopt.ToJpeg = false
	cliopt.ToPng = false
	cliopt.ToGif = false
	cliopt.Verbose = true
	cliopt.Force = true
}

func expectToMatchBuffer(t *testing.T, buffer *bytes.Buffer, expected string) {
	if !regexp.MustCompile(expected).MatchString(buffer.String()) {
		t.Errorf("expected: %s, actual: %s", expected, buffer.String())
	}
}
