package cmd

import (
	"bytes"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os/exec"
	"regexp"
	"testing"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/conversion"
)

func TestJpegToPng(t *testing.T) {
	t.Parallel()

	tmpdir := "../tmp/TestJpegToPng"

	exec.Command("cp", "-r", "../testdata/", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	runner := &Runner{OutStream: buf, Decoder: &conversion.Jpeg{}, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}, Force: true}
	runner.Run(tmpdir)

	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestJpegToPng/jpeg/sample1.png"`)
	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestJpegToPng/jpeg/sample2.png"`)
	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestJpegToPng/jpeg/sample3.png"`)
}

func TestJpegToGif(t *testing.T) {
	t.Parallel()

	tmpdir := "../tmp/TestJpegToGif"

	exec.Command("cp", "-r", "../testdata/", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	runner := &Runner{OutStream: buf, Decoder: &conversion.Jpeg{}, Encoder: &conversion.Gif{Options: &gif.Options{NumColors: 256}}, Force: true}
	runner.Run(tmpdir)

	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestJpegToGif/jpeg/sample1.gif"`)
	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestJpegToGif/jpeg/sample2.gif"`)
	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestJpegToGif/jpeg/sample3.gif"`)
}

func TestPngToJpeg(t *testing.T) {
	t.Parallel()

	tmpdir := "../tmp/TestPngToJpeg"

	exec.Command("cp", "-r", "../testdata/", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	runner := &Runner{OutStream: buf, Decoder: &conversion.Png{}, Encoder: &conversion.Jpeg{Options: &jpeg.Options{Quality: 100}}, Force: true}
	runner.Run(tmpdir)

	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestPngToJpeg/png/sample1.jpg"`)
	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestPngToJpeg/png/sample2.jpg"`)
}

func TestPngToGif(t *testing.T) {
	t.Parallel()

	tmpdir := "../tmp/TestPngToGif"

	exec.Command("cp", "-r", "../testdata/", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	runner := &Runner{OutStream: buf, Decoder: &conversion.Png{}, Encoder: &conversion.Gif{Options: &gif.Options{NumColors: 256}}, Force: true}
	runner.Run(tmpdir)

	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestPngToGif/png/sample1.gif"`)
	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestPngToGif/png/sample2.gif"`)
}

func TestGifToJpeg(t *testing.T) {
	t.Parallel()

	tmpdir := "../tmp/TestGifToJpeg"

	exec.Command("cp", "-r", "../testdata/", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	runner := &Runner{OutStream: buf, Decoder: &conversion.Gif{}, Encoder: &conversion.Jpeg{Options: &jpeg.Options{Quality: 100}}, Force: true}
	runner.Run(tmpdir)

	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestGifToJpeg/gif/sample1.jpg"`)
}

func TestGifToPng(t *testing.T) {
	t.Parallel()

	tmpdir := "../tmp/TestGifToPng"

	exec.Command("cp", "-r", "../testdata/", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	runner := &Runner{OutStream: buf, Decoder: &conversion.Gif{}, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}, Force: true}
	runner.Run(tmpdir)

	expectToMatchBuffer(t, buf, `Converted: "../tmp/TestGifToPng/gif/sample1.png"`)
}

func TestConflict(t *testing.T) {
	t.Parallel()

	tmpdir := "../tmp/TestConflict"

	exec.Command("cp", "-r", "../testdata/", tmpdir).Run()
	defer exec.Command("rm", "-r", tmpdir).Run()

	buf := &bytes.Buffer{}

	runner := &Runner{OutStream: buf, Decoder: &conversion.Jpeg{}, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}, Force: true}
	runner.Run(tmpdir)

	runner = &Runner{OutStream: buf, Decoder: &conversion.Jpeg{}, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}, Force: false}
	err := runner.Run(tmpdir)

	expected := "File already exists: ../tmp/TestConflict/jpeg/sample1.png"
	if err.Error() != expected {
		t.Errorf("expected: %s, actual: %s", expected, err)
	}
}

func expectToMatchBuffer(t *testing.T, buffer *bytes.Buffer, expected string) {
	if !regexp.MustCompile(expected).MatchString(buffer.String()) {
		t.Errorf("expected: %s, actual: %s", expected, buffer.String())
	}
}
