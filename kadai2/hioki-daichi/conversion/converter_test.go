package conversion

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/hioki-daichi/fileutil"
)

func TestConversion_Convert(t *testing.T) {
	cases := map[string]struct {
		decoder  Decoder
		encoder  Encoder
		path     string
		force    bool
		expected error
	}{
		"JPEG to PNG": {decoder: jpegDecoder(), encoder: pngEncoder(), path: "./jpeg/sample1.jpg", force: true, expected: nil},
		"JPEG to GIF": {decoder: jpegDecoder(), encoder: gifEncoder(), path: "./jpeg/sample1.jpg", force: true, expected: nil},
		"PNG to JPEG": {decoder: pngDecoder(), encoder: jpegEncoder(), path: "./png/sample1.png", force: true, expected: nil},
		"PNG to GIF":  {decoder: pngDecoder(), encoder: gifEncoder(), path: "./png/sample1.png", force: true, expected: nil},
		"GIF to JPEG": {decoder: gifDecoder(), encoder: jpegEncoder(), path: "./gif/sample1.gif", force: true, expected: nil},
		"GIF to PNG":  {decoder: gifDecoder(), encoder: pngEncoder(), path: "./gif/sample1.gif", force: true, expected: nil},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			converter := &Converter{Decoder: c.decoder, Encoder: c.encoder}

			withTempDir(t, func(t *testing.T, tempdir string) {
				_, actual := converter.Convert(filepath.Join(tempdir, c.path), c.force)
				if actual != c.expected {
					t.Errorf(`expected="%s" actual="%s"`, c.expected, actual)
				}
			})
		})
	}
}

func TestConversion_Convert_Conflict(t *testing.T) {
	t.Parallel()

	converter := &Converter{Decoder: jpegDecoder(), Encoder: pngEncoder()}

	withTempDir(t, func(t *testing.T, tempdir string) {
		expected := "File already exists: " + tempdir + "/jpeg/sample1.png"

		path := filepath.Join(tempdir, "./jpeg/sample1.jpg")

		_, err := converter.Convert(path, false)
		if err != nil {
			t.Fatalf("err %s", err)
		}
		_, err = converter.Convert(path, false)

		actual := err.Error()
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})
}

func TestConversion_Convert_Nonexistence(t *testing.T) {
	t.Parallel()

	expected := "open ./nonexistent_path: no such file or directory"

	converter := &Converter{Decoder: jpegDecoder(), Encoder: pngEncoder()}

	_, err := converter.Convert("./nonexistent_path", true)

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestConversion_Convert_Undecodable(t *testing.T) {
	t.Parallel()

	expected := "unexpected EOF"

	converter := &Converter{Decoder: jpegDecoder(), Encoder: pngEncoder()}

	_, err := converter.Convert("./testdata/undecodable.jpg", true)

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestConversion_Convert_CreationFailure(t *testing.T) {
	t.Parallel()

	converter := &Converter{Decoder: jpegDecoder(), Encoder: pngEncoder()}

	withTempDir(t, func(t *testing.T, tempdir string) {
		expected := "open " + tempdir + "/jpeg/sample1.png: permission denied"

		src := filepath.Join(tempdir, "./jpeg/sample1.jpg")
		dst := filepath.Join(tempdir, "./jpeg/sample1.png")

		// First, create a file without permission in PATH after conversion,
		_, err := os.OpenFile(dst, os.O_CREATE|os.O_EXCL, 0)
		if err != nil {
			t.Fatalf("err %s", err)
		}

		// then convert.
		_, err = converter.Convert(src, true)

		actual := err.Error()
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})
}

func TestConversion_Convert_EncodeFailure(t *testing.T) {
	t.Parallel()

	expected := "error in EncodeMock.Encode"

	converter := &Converter{Decoder: jpegDecoder(), Encoder: mockEncoder()}

	withTempDir(t, func(t *testing.T, tempdir string) {
		path := filepath.Join(tempdir, "./jpeg/sample1.jpg")

		_, err := converter.Convert(path, true)

		actual := err.Error()
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})
}

func jpegDecoder() *Jpeg {
	return &Jpeg{}
}

func pngDecoder() *Png {
	return &Png{}
}

func gifDecoder() *Gif {
	return &Gif{}
}

func jpegEncoder() *Jpeg {
	return &Jpeg{Options: &jpeg.Options{Quality: 1}}
}

func pngEncoder() *Png {
	return &Png{Encoder: &png.Encoder{CompressionLevel: png.NoCompression}}
}

func gifEncoder() *Gif {
	return &Gif{Options: &gif.Options{NumColors: 1}}
}

func withTempDir(t *testing.T, f func(t *testing.T, tempdir string)) {
	t.Helper()

	tempdir, err := ioutil.TempDir("", "imgconv")
	if err != nil {
		t.Fatalf("err %s", err)
	}

	err = fileutil.CopyDirRec("../testdata/", tempdir)
	if err != nil {
		t.Fatalf("err %s", err)
	}
	defer os.RemoveAll(tempdir)

	f(t, tempdir)
}

type EncoderMock struct {
	Png
}

func (m *EncoderMock) Encode(w io.Writer, img image.Image) error {
	return errors.New("error in EncodeMock.Encode")
}

func mockEncoder() *EncoderMock {
	return &EncoderMock{}
}
