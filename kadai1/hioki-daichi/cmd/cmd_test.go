package cmd

import (
	"bytes"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os/exec"
	"regexp"
	"testing"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/conversion"
	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/fileutil"
)

func TestJpegToPng(t *testing.T) {
	t.Parallel()

	withTempDir(t, func(t *testing.T, tempdir string) {
		buf := &bytes.Buffer{}

		runner := &Runner{OutStream: buf, Decoder: &conversion.Jpeg{}, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}, Force: true}
		runner.Run(tempdir)

		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/jpeg/sample1.png"`)
		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/jpeg/sample2.png"`)
		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/jpeg/sample3.png"`)
	})
}

func TestJpegToGif(t *testing.T) {
	t.Parallel()

	withTempDir(t, func(t *testing.T, tempdir string) {
		buf := &bytes.Buffer{}

		runner := &Runner{OutStream: buf, Decoder: &conversion.Jpeg{}, Encoder: &conversion.Gif{Options: &gif.Options{NumColors: 256}}, Force: true}
		runner.Run(tempdir)

		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/jpeg/sample1.gif"`)
		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/jpeg/sample2.gif"`)
		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/jpeg/sample3.gif"`)
	})
}

func TestPngToJpeg(t *testing.T) {
	t.Parallel()

	withTempDir(t, func(t *testing.T, tempdir string) {
		buf := &bytes.Buffer{}

		runner := &Runner{OutStream: buf, Decoder: &conversion.Png{}, Encoder: &conversion.Jpeg{Options: &jpeg.Options{Quality: 100}}, Force: true}
		runner.Run(tempdir)

		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/png/sample1.jpg"`)
		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/png/sample2.jpg"`)
	})
}

func TestPngToGif(t *testing.T) {
	t.Parallel()

	withTempDir(t, func(t *testing.T, tempdir string) {
		buf := &bytes.Buffer{}

		runner := &Runner{OutStream: buf, Decoder: &conversion.Png{}, Encoder: &conversion.Gif{Options: &gif.Options{NumColors: 256}}, Force: true}
		runner.Run(tempdir)

		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/png/sample1.gif"`)
		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/png/sample2.gif"`)
	})
}

func TestGifToJpeg(t *testing.T) {
	t.Parallel()

	withTempDir(t, func(t *testing.T, tempdir string) {
		buf := &bytes.Buffer{}

		runner := &Runner{OutStream: buf, Decoder: &conversion.Gif{}, Encoder: &conversion.Jpeg{Options: &jpeg.Options{Quality: 100}}, Force: true}
		runner.Run(tempdir)

		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/gif/sample1.jpg"`)
	})
}

func TestGifToPng(t *testing.T) {
	t.Parallel()

	withTempDir(t, func(t *testing.T, tempdir string) {
		buf := &bytes.Buffer{}

		runner := &Runner{OutStream: buf, Decoder: &conversion.Gif{}, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}, Force: true}
		runner.Run(tempdir)

		expectToMatchBuffer(t, buf, `Converted: "`+tempdir+`/gif/sample1.png"`)
	})
}

func TestConflict(t *testing.T) {
	t.Parallel()

	withTempDir(t, func(t *testing.T, tempdir string) {
		buf := &bytes.Buffer{}

		runner := &Runner{OutStream: buf, Decoder: &conversion.Jpeg{}, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}, Force: true}
		runner.Run(tempdir)

		runner = &Runner{OutStream: buf, Decoder: &conversion.Jpeg{}, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}, Force: false}
		err := runner.Run(tempdir)

		expected := "File already exists: " + tempdir + "/jpeg/sample1.png"
		if err.Error() != expected {
			t.Errorf("expected: %s, actual: %s", expected, err)
		}
	})
}

func withTempDir(t *testing.T, f func(t *testing.T, tempdir string)) {
	tempdir, _ := ioutil.TempDir("", "imgconv-testing-")
	fileutil.CopyDirRec("../testdata/", tempdir)
	defer exec.Command("rm", "-r", tempdir).Run()
	f(t, tempdir)
}

func expectToMatchBuffer(t *testing.T, buffer *bytes.Buffer, expected string) {
	if !regexp.MustCompile(expected).MatchString(buffer.String()) {
		t.Errorf("expected: %s, actual: %s", expected, buffer.String())
	}
}
