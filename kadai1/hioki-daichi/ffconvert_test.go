package main

import (
	"bytes"
	"os/exec"
	"regexp"
	"testing"
)

func TestJpegToPng(t *testing.T) {
	t.Parallel()

	in := Jpeg
	out := Png
	tmpdir := "test/" + "TestJpegToPng"

	exec.Command("cp", "-r", "test/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}
	force := true
	verbose := true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.Execute([]string{tmpdir}, in, out, force, verbose)

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
	force := true
	verbose := true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.Execute([]string{tmpdir}, in, out, force, verbose)

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
	force := true
	verbose := true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.Execute([]string{tmpdir}, in, out, force, verbose)

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
	force := true
	verbose := true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.Execute([]string{tmpdir}, in, out, force, verbose)

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
	force := true
	verbose := true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.Execute([]string{tmpdir}, in, out, force, verbose)

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
	force := true
	verbose := true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.Execute([]string{tmpdir}, in, out, force, verbose)

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

func TestNoArgs(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.Execute([]string{}, Jpeg, Png, true, true)
	expectToMatchBuffer(t, buf, "Specify filenames as an arguments")
}

func TestConflict(t *testing.T) {
	t.Parallel()

	in := Jpeg
	out := Png
	tmpdir := "test/" + "TestConflict"

	exec.Command("cp", "-r", "test/images", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}
	verbose := true

	cli := &CLI{OutStream: buf, ErrStream: buf}
	cli.Execute([]string{tmpdir}, in, out, true, verbose)
	cli.Execute([]string{tmpdir}, in, out, false, verbose)

	expectToMatchBuffer(t, buf, "File already exists: test/TestConflict/2018/07/001.png")
}

func expectToMatchBuffer(t *testing.T, buffer *bytes.Buffer, expected string) {
	if !regexp.MustCompile(expected).MatchString(buffer.String()) {
		t.Errorf("expected: %s, actual: %s", expected, buffer.String())
	}
}
