package main

import (
	"bytes"
	"os/exec"
	"regexp"
	"testing"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/cliopt"
)

func TestJpegToPng(t *testing.T) {
	t.Parallel()

	in := Jpeg
	out := Png
	tmpdir := "test/" + "TestJpegToPng"

	exec.Command("cp", "-r", "test/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromJpeg = true
	cliopt.ToPng = true

	cli := &CLI{OutStream: buf, ErrStream: buf, in: in, out: out}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestJpegToPng"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestJpegToPng/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestJpegToPng/2018/07"`)
	expectToMatchBuffer(t, buf, `Converted: "test/TestJpegToPng/2018/07/001.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestJpegToPng/2018/07/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestJpegToPng/2018/08"`)
	expectToMatchBuffer(t, buf, `Converted: "test/TestJpegToPng/2018/08/001.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestJpegToPng/2018/08/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestJpegToPng/2018/08/003.gif"`)
}

func TestJpegToGif(t *testing.T) {
	t.Parallel()

	in := Jpeg
	out := Gif
	tmpdir := "test/" + "TestJpegToGif"

	exec.Command("cp", "-r", "test/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromJpeg = true
	cliopt.ToGif = true

	cli := &CLI{OutStream: buf, ErrStream: buf, in: in, out: out}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestJpegToGif"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestJpegToGif/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestJpegToGif/2018/07"`)
	expectToMatchBuffer(t, buf, `Converted: "test/TestJpegToGif/2018/07/001.gif"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestJpegToGif/2018/07/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestJpegToGif/2018/08"`)
	expectToMatchBuffer(t, buf, `Converted: "test/TestJpegToGif/2018/08/001.gif"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestJpegToGif/2018/08/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestJpegToGif/2018/08/003.gif"`)
}

func TestPngToJpeg(t *testing.T) {
	t.Parallel()

	in := Png
	out := Jpeg
	tmpdir := "test/" + "TestPngToJpeg"

	exec.Command("cp", "-r", "test/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromPng = true
	cliopt.ToJpeg = true

	cli := &CLI{OutStream: buf, ErrStream: buf, in: in, out: out}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestPngToJpeg"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestPngToJpeg/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestPngToJpeg/2018/07"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestPngToJpeg/2018/07/001.jpg"`)
	expectToMatchBuffer(t, buf, `Converted: "test/TestPngToJpeg/2018/07/002.jpg"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestPngToJpeg/2018/08"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestPngToJpeg/2018/08/001.jpg"`)
	expectToMatchBuffer(t, buf, `Converted: "test/TestPngToJpeg/2018/08/002.jpg"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestPngToJpeg/2018/08/003.gif"`)
}

func TestPngToGif(t *testing.T) {
	t.Parallel()

	in := Png
	out := Gif
	tmpdir := "test/" + "TestPngToGif"

	exec.Command("cp", "-r", "test/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromPng = true
	cliopt.ToGif = true

	cli := &CLI{OutStream: buf, ErrStream: buf, in: in, out: out}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestPngToGif"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestPngToGif/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestPngToGif/2018/07"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestPngToGif/2018/07/001.jpg"`)
	expectToMatchBuffer(t, buf, `Converted: "test/TestPngToGif/2018/07/002.gif"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestPngToGif/2018/08"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestPngToGif/2018/08/001.jpg"`)
	expectToMatchBuffer(t, buf, `Converted: "test/TestPngToGif/2018/08/002.gif"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestPngToGif/2018/08/003.gif"`)
}

func TestGifToJpeg(t *testing.T) {
	t.Parallel()

	in := Gif
	out := Jpeg
	tmpdir := "test/" + "TestGifToJpeg"

	exec.Command("cp", "-r", "test/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromGif = true
	cliopt.ToJpeg = true

	cli := &CLI{OutStream: buf, ErrStream: buf, in: in, out: out}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestGifToJpeg"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestGifToJpeg/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestGifToJpeg/2018/07"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestGifToJpeg/2018/07/001.jpg"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestGifToJpeg/2018/07/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestGifToJpeg/2018/08"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestGifToJpeg/2018/08/001.jpg"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestGifToJpeg/2018/08/002.png"`)
	expectToMatchBuffer(t, buf, `Converted: "test/TestGifToJpeg/2018/08/003.jpg"`)
}

func TestGifToPng(t *testing.T) {
	t.Parallel()

	in := Gif
	out := Png
	tmpdir := "test/" + "TestGifToPng"

	exec.Command("cp", "-r", "test/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromGif = true
	cliopt.ToPng = true

	cli := &CLI{OutStream: buf, ErrStream: buf, in: in, out: out}
	cli.execute(tmpdir)

	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestGifToPng"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestGifToPng/2018"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestGifToPng/2018/07"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestGifToPng/2018/07/001.jpg"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestGifToPng/2018/07/002.png"`)
	expectToMatchBuffer(t, buf, `Skipped because the path is directory: "test/TestGifToPng/2018/08"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestGifToPng/2018/08/001.jpg"`)
	expectToMatchBuffer(t, buf, `Skipped because the file is not applicable: "test/TestGifToPng/2018/08/002.png"`)
	expectToMatchBuffer(t, buf, `Converted: "test/TestGifToPng/2018/08/003.png"`)
}

func TestConflict(t *testing.T) {
	t.Parallel()

	in := Jpeg
	out := Png
	tmpdir := "test/" + "TestConflict"

	exec.Command("cp", "-r", "test/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	initCliopt()
	cliopt.FromJpeg = true
	cliopt.ToPng = true

	cliopt.Verbose = true
	cliopt.Force = true
	cli := &CLI{OutStream: buf, ErrStream: buf, in: in, out: out}
	cli.execute(tmpdir)

	cliopt.Verbose = true
	cliopt.Force = false
	cli = &CLI{OutStream: buf, ErrStream: buf, in: in, out: out}
	err := cli.execute(tmpdir)

	expected := "File already exists: test/TestConflict/2018/07/001.png"
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
