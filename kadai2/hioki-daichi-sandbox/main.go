/*
io.Reader と io.Writer が標準パッケージでどのように使われているか調べるために、
Read() や Write() を実装しているパッケージをいくつか抜き出し、
それぞれ Read() や Write() を呼び出してみました。

抜き出した対象のパッケージは以下です。

- strings
- os
- bytes
- tar
- zip
- zlib
- iotest
*/
package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"strings"
	"testing/iotest"
)

func main() {
	// strings
	readStringsExample() // strings

	// os
	readWriteFileExample("./testdata/file.txt") // file

	// bytes
	readWriteBufferExample() // bytes

	// tar
	readWriteTarExample("./testdata/ab.tar") // a\nb\n

	// zip
	readWriteZipExample("./testdata/ab.zip") // a\nb\n

	// zlib
	readWriteZlibExample() // foo

	// iotest
	oneByteReaderExample() // bar
	halfReaderExample()    // hoge
	readDataErrExample()   // EOF
	readTimeoutExample()   // 8 <nil>\n0 timeout
}

func readWriteFileExample(path string) {
	var r io.Reader
	var w io.Writer

	p := []byte("file\n")

	w, _ = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	_, _ = w.Write(p)
	_ = w.(io.Closer).Close()

	w = os.Stdout

	r, _ = os.Open(path)
	_, _ = io.Copy(w, r)

	_ = r.(io.Closer).Close()
}

func readStringsExample() {
	var r io.Reader
	var w io.Writer

	r = strings.NewReader("strings\n")
	w = os.Stdout

	_, _ = io.Copy(w, r)
}

func readWriteBufferExample() {
	var r io.Reader
	var w io.Writer

	p := []byte("bytes\n")

	var buf bytes.Buffer

	r = &buf
	w = &buf

	_, _ = w.Write(p)

	w = os.Stdout
	_, _ = io.Copy(w, r)
}

func readWriteTarExample(path string) {
	var r io.Reader
	var w io.Writer

	filepaths := []string{
		"./testdata/a.txt",
		"./testdata/b.txt",
	}

	_ = os.Remove(path)

	tfp, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	tw := tar.NewWriter(tfp)

	for _, filepath := range filepaths {
		fp, _ := os.Open(filepath)

		info, _ := fp.Stat()
		size := info.Size()
		mode := int64(info.Mode())

		b := make([]byte, size)
		_, _ = fp.Read(b)

		_ = fp.Close()

		hdr := &tar.Header{
			Name: filepath,
			Size: size,
			Mode: mode,
		}
		_ = tw.WriteHeader(hdr)
		_, _ = tw.Write(b)
	}

	_ = tw.Close()
	_ = tfp.Close()

	w = os.Stdout

	r, _ = os.Open(path)
	tr := tar.NewReader(r)
	for {
		_, err := tr.Next()
		if err == io.EOF {
			break
		}
		_, _ = io.Copy(w, tr)
	}
	_ = r.(io.Closer).Close()
}

func readWriteZipExample(path string) {
	var w io.Writer

	filepaths := []string{
		"./testdata/a.txt",
		"./testdata/b.txt",
	}

	_ = os.Remove(path)

	tfp, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	zw := zip.NewWriter(tfp)

	for _, filepath := range filepaths {
		fp, _ := os.Open(filepath)

		info, _ := fp.Stat()

		hdr, _ := zip.FileInfoHeader(info)
		hdr.Name = filepath
		hdr.Method = zip.Deflate

		w, _ = zw.CreateHeader(hdr)
		_, _ = io.Copy(w, fp)
		_ = fp.Close()
	}

	_ = zw.Close()
	_ = tfp.Close()

	zrc, _ := zip.OpenReader(path)
	for _, fp := range zrc.File {
		rc, _ := fp.Open()
		_, _ = io.Copy(os.Stdout, rc)
		_ = rc.Close()
	}
	_ = zrc.Close()
}

func readWriteZlibExample() {
	p := []byte("foo\n")

	var r io.Reader
	var w io.Writer

	var buf bytes.Buffer

	zw := zlib.NewWriter(&buf)
	_, _ = zw.Write(p)
	_ = zw.Close()

	r = bytes.NewReader(buf.Bytes())
	zr, _ := zlib.NewReader(r)
	w = os.Stdout
	_, _ = io.Copy(w, zr)
	_ = zr.Close()
}

func oneByteReaderExample() {
	p := []byte("bar")
	br := bytes.NewReader(p)
	obr := iotest.OneByteReader(br)

	for i := 0; i < len(p); i++ {
		b := make([]byte, 1)
		_, _ = obr.Read(b)
		fmt.Print(string(b))
	}
	fmt.Println()
}

func halfReaderExample() {
	p := []byte("hogehoge")

	br := bytes.NewReader(p)
	hr := iotest.HalfReader(br)

	b := make([]byte, len(p))
	_, _ = hr.Read(b)
	fmt.Println(string(b))
}

func readDataErrExample() {
	p := []byte("baz")
	br := bytes.NewReader(p)
	der := iotest.DataErrReader(br)
	b := make([]byte, len(p))
	_, err := der.Read(b)
	fmt.Println(err)
}

func readTimeoutExample() {
	p := []byte("fugafuga")

	var (
		n   int
		err error
	)

	br := bytes.NewReader(p)
	tr := iotest.TimeoutReader(br)

	b := make([]byte, len(p))

	n, err = tr.Read(b)
	fmt.Println(n, err) // 8 <nil>

	n, err = tr.Read(b)
	fmt.Println(n, err) // 0 timeout
}
