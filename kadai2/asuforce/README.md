# Kadai2

## io.Readerとio.Writerについて調べてみよう

### 標準パッケージでどのように使われているか

`io.Reader` を実装している

```txt
src/archive/tar/reader.go
src/archive/zip/reader.go
src/compress/zlib/reader.go
src/debug/elf/reader.go
src/encoding/csv/reader.go
src/image/gif/reader.go
src/image/jpeg/reader.go
src/image/png/reader.go
src/mime/quotedprintable/reader.go
src/strings/reader.go
src/testing/iotest/reader.go
```

`io.Writer` を実装している

```txt
src/archive/tar/writer.go
src/archive/zip/writer.go
src/compress/zlib/writer.go
src/encoding/csv/writer.go
src/image/gif/writer.go
src/image/jpeg/writer.go
src/image/png/writer.go
src/mime/quotedprintable/writer.go
src/testing/iotest/writer.go
```

これらの実装から、幅広い入出力に対応していることがわかる。

### io.Reader と io.Writer があることでどういう利点があるのか具体例を挙げて考えてみる

Reader と Writer が抽象化されていることで、様々な形式の入出力に対応できる柔軟性を備える事ができる。
Read, Write が別れている事でそれぞれの処理の実装に専念することができる。
また、これらのインターフェースをユーザ自身で実装することで、独自の機能追加が行いやすくなる拡張性も併せ持つ事ができる。
