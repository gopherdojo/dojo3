package cmd

import (
	"bytes"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/hioki-daichi/conversion"
	"github.com/gopherdojo/dojo3/kadai2/hioki-daichi/fileutil"
)

// Decoder
var jpegD conversion.Jpeg
var pngD conversion.Png
var gifD conversion.Gif

// Encoder
var jpegE = &conversion.Jpeg{Options: &jpeg.Options{Quality: 1}}
var pngE = &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.NoCompression}}
var gifE = &conversion.Gif{Options: &gif.Options{NumColors: 1}}

func TestCmd_Run(t *testing.T) {
	cases := map[string]struct {
		decoder conversion.Decoder
		encoder conversion.Encoder
		force   bool

		// Reason: To be able to handle "tempdir"
		expected func(string) string
	}{
		"JPEG to PNG": {decoder: &jpegD, encoder: pngE, force: true, expected: func(tempdir string) string {
			return `Converted: "` + tempdir + `/jpeg/sample1.png"
Converted: "` + tempdir + `/jpeg/sample2.png"
Converted: "` + tempdir + `/jpeg/sample3.png"
`
		}},
		"JPEG to GIF": {decoder: &jpegD, encoder: gifE, force: true, expected: func(tempdir string) string {
			return `Converted: "` + tempdir + `/jpeg/sample1.gif"
Converted: "` + tempdir + `/jpeg/sample2.gif"
Converted: "` + tempdir + `/jpeg/sample3.gif"
`
		}},
		"PNG to JPEG": {decoder: &pngD, encoder: jpegE, force: true, expected: func(tempdir string) string {
			return `Converted: "` + tempdir + `/png/sample1.jpg"
Converted: "` + tempdir + `/png/sample2.jpg"
`
		}},
		"PNG to GIF": {decoder: &pngD, encoder: gifE, force: true, expected: func(tempdir string) string {
			return `Converted: "` + tempdir + `/png/sample1.gif"
Converted: "` + tempdir + `/png/sample2.gif"
`
		}},
		"GIF to JPEG": {decoder: &gifD, encoder: jpegE, force: true, expected: func(tempdir string) string {
			return `Converted: "` + tempdir + `/gif/sample1.jpg"
`
		}},
		"GIF to PNG": {decoder: &gifD, encoder: pngE, force: true, expected: func(tempdir string) string {
			return `Converted: "` + tempdir + `/gif/sample1.png"
`
		}},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			buf := &bytes.Buffer{}

			runner := Runner{OutStream: buf, Decoder: c.decoder, Encoder: c.encoder, Force: c.force}

			withTempDir(t, func(t *testing.T, tempdir string) {
				expected := c.expected(tempdir)

				err := runner.Run(tempdir)
				if err != nil {
					t.Fatalf("err %s", err)
				}

				actual := buf.String()
				if actual != expected {
					t.Errorf(`expected="%s" actual="%s"`, expected, actual)
				}
			})
		})
	}
}

func TestCmd_Run_Nonexistence(t *testing.T) {
	t.Parallel()

	expected := "lstat nonexistent_path: no such file or directory"

	runner := Runner{OutStream: ioutil.Discard, Decoder: &jpegD, Encoder: pngE, Force: true}

	err := runner.Run("nonexistent_path")

	actual := err.Error()
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestCmd_Run_Conflict(t *testing.T) {
	t.Parallel()

	w := ioutil.Discard

	withTempDir(t, func(t *testing.T, tempdir string) {
		expected := "File already exists: " + tempdir + "/jpeg/sample1.png"

		var runner Runner
		var err error

		runner = Runner{OutStream: w, Decoder: &jpegD, Encoder: pngE, Force: true}
		err = runner.Run(tempdir)
		if err != nil {
			t.Fatalf("err %s", err)
		}

		runner = Runner{OutStream: w, Decoder: &jpegD, Encoder: pngE, Force: false}
		err = runner.Run(tempdir)
		actual := err.Error()
		if actual != expected {
			t.Errorf(`expected="%s" actual="%s"`, expected, actual)
		}
	})
}

func withTempDir(t *testing.T, f func(t *testing.T, tempdir string)) {
	t.Helper()

	tempdir, _ := ioutil.TempDir("", "imgconv")

	err := fileutil.CopyDirRec("../testdata/", tempdir)
	if err != nil {
		t.Fatalf("err %s", err)
	}
	defer os.RemoveAll(tempdir)

	f(t, tempdir)
}
