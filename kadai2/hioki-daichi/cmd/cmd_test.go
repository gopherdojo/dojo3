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

func TestCmd_Run(t *testing.T) {
	cases := map[string]struct {
		decoder conversion.Decoder
		encoder conversion.Encoder
		force   bool

		// Reason: To be able to handle "tempdir"
		expected func(string) string
	}{
		"JPEG to PNG": {decoder: jpegDecoder(t), encoder: pngEncoder(t), force: true, expected: func(tempdir string) string {
			return `Converted: "` + tempdir + `/jpeg/sample1.png"
Converted: "` + tempdir + `/jpeg/sample2.png"
Converted: "` + tempdir + `/jpeg/sample3.png"
`
		}},
		"JPEG to GIF": {decoder: jpegDecoder(t), encoder: gifEncoder(t), force: true, expected: func(tempdir string) string {
			return `Converted: "` + tempdir + `/jpeg/sample1.gif"
Converted: "` + tempdir + `/jpeg/sample2.gif"
Converted: "` + tempdir + `/jpeg/sample3.gif"
`
		}},
		"PNG to JPEG": {decoder: pngDecoder(t), encoder: jpegEncoder(t), force: true, expected: func(tempdir string) string {
			return `Converted: "` + tempdir + `/png/sample1.jpg"
Converted: "` + tempdir + `/png/sample2.jpg"
`
		}},
		"PNG to GIF": {decoder: pngDecoder(t), encoder: gifEncoder(t), force: true, expected: func(tempdir string) string {
			return `Converted: "` + tempdir + `/png/sample1.gif"
Converted: "` + tempdir + `/png/sample2.gif"
`
		}},
		"GIF to JPEG": {decoder: gifDecoder(t), encoder: jpegEncoder(t), force: true, expected: func(tempdir string) string {
			return `Converted: "` + tempdir + `/gif/sample1.jpg"
`
		}},
		"GIF to PNG": {decoder: gifDecoder(t), encoder: pngEncoder(t), force: true, expected: func(tempdir string) string {
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

			tempdir, cleanFn := withTempDir(t)
			defer cleanFn()

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
	}
}

func TestCmd_Run_Nonexistence(t *testing.T) {
	t.Parallel()

	expected := "lstat nonexistent_path: no such file or directory"

	runner := Runner{OutStream: ioutil.Discard, Decoder: jpegDecoder(t), Encoder: pngEncoder(t), Force: true}

	err := runner.Run("nonexistent_path")

	actual := err.Error()
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestCmd_Run_Conflict(t *testing.T) {
	t.Parallel()

	w := ioutil.Discard

	tempdir, cleanFn := withTempDir(t)
	defer cleanFn()

	expected := "File already exists: " + tempdir + "/jpeg/sample1.png"

	var runner Runner
	var err error

	runner = Runner{OutStream: w, Decoder: jpegDecoder(t), Encoder: pngEncoder(t), Force: true}
	err = runner.Run(tempdir)
	if err != nil {
		t.Fatalf("err %s", err)
	}

	runner = Runner{OutStream: w, Decoder: jpegDecoder(t), Encoder: pngEncoder(t), Force: false}
	err = runner.Run(tempdir)
	actual := err.Error()
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func withTempDir(t *testing.T) (string, func()) {
	t.Helper()

	tempdir, err := ioutil.TempDir("", "imgconv")
	if err != nil {
		t.Fatalf("err %s", err)
	}

	err = fileutil.CopyDirRec("../testdata/", tempdir)
	if err != nil {
		t.Fatalf("err %s", err)
	}

	return tempdir, func() { os.RemoveAll(tempdir) }
}

func jpegDecoder(t *testing.T) *conversion.Jpeg {
	t.Helper()
	var d conversion.Jpeg
	return &d
}

func pngDecoder(t *testing.T) *conversion.Png {
	t.Helper()
	var d conversion.Png
	return &d
}

func gifDecoder(t *testing.T) *conversion.Gif {
	t.Helper()
	var d conversion.Gif
	return &d
}

func jpegEncoder(t *testing.T) *conversion.Jpeg {
	t.Helper()
	return &conversion.Jpeg{Options: &jpeg.Options{Quality: 1}}
}

func pngEncoder(t *testing.T) *conversion.Png {
	t.Helper()
	return &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.NoCompression}}
}

func gifEncoder(t *testing.T) *conversion.Gif {
	t.Helper()
	return &conversion.Gif{Options: &gif.Options{NumColors: 1}}
}
