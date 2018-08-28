# kadai2

##io.Readerとio.Writerについて調べてみよう
- 標準パッケージでどのように使われているか
- io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

##1回目の宿題のテストを作ってみて下さい
- テストのしやすさを考えてリファクタリングしてみる
- テストのカバレッジを取ってみる
- テーブル駆動テストを行う
- テストヘルパーを作ってみる

# install
```
make
```

# usage

```
bin/imgconv [OPTION] dir_path
  -i string
    	変換対象の画像形式(jpeg|gif|png) (default "jpeg")
  -o string
    	変換語の画像形式(jpeg|gif|png) (default "png")
  -v	詳細なログを表示
```

#io.Readerとio.Writerについて調べてみよう
##標準パッケージでどのように使われているか?

io.Reader と io.Writerを用いることによって、ファイル、ネットワーク、文字列、バイト配列等々について入出力に関するインターフェイスが抽象化されています。

以下go1.11において使われているパッケージを数例示します。

###bufioパッケージ
バッファー化されたI/0への読み書き。

- [bufio/bufio.go : func (b *Reader) Read(p []byte) (n int, err error) {](https://github.com/golang/go/blob/go1.11/src/bufio/bufio.go#L190#L232)
- [bufio/bufio.go : func (b *Writer) Write(p []byte) (nn int, err error) {](https://github.com/golang/go/blob/go1.11/src/bufio/bufio.go#L601#L623)

###stringsパッケージ
文字列への読み書き。

- [strings/reader.go : func (r *Reader) Read(b []byte) (n int, err error) {](https://github.com/golang/go/blob/go1.11/src/strings/reader.go#L37#L45)
- [strings/builder.go : func (b *Builder) Write(p []byte) (int, error) {](https://github.com/golang/go/blob/go1.11/src/strings/builder.go#L82#L86)

###bytesパッケージ
バイト配列への読み書き。Bufferは読み書き可能なバッファとして動作するのに対し、Readerはリードオンリーでシークが出来る。

- [bytes/buffer.go : func (b *Buffer) Read(p []byte) (n int, err error) {](https://github.com/golang/go/blob/go1.11/src/bytes/buffer.go#L298#L314)
- [bytes/buffer.go : func (b *Buffer) Write(p []byte) (n int, err error) {](https://github.com/golang/go/blob/go1.11/src/bytes/buffer.go#L170#L177)
- [bytes/reader.go : func (r *Reader) Read(b []byte) (n int, err error) {](https://github.com/golang/go/blob/go1.11/src/bytes/reader.go#L39#L47)

###osパッケージ
os/fileにて、ファイルへの読み書き。

- [os/file.go : func (f *File) Read(b []byte) (n int, err error) {](https://github.com/golang/go/blob/go1.11/src/os/file.go#L104#L110)
- [os/file.go : func (f *File) Write(b []byte) (n int, err error) {](https://github.com/golang/go/blob/go1.11/src/os/file.go#L141#L160)

##io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

### 具象型毎に関数を作る必要がなくなる
io.Readerがある事によって型が異なっていても同じio.Readerインターフェイスを持っているものであれば、同一関数で処理を行うことができるようになる。逆にないと型毎に関数を作らなければならなくなる。

- io.Readerがある場合
https://play.golang.org/p/6IW_oU8xIw5

- io.Readerがない場合
https://play.golang.org/p/XPu2SCfIa88

### 実装が入れ替えられるので、テストがしやすくなる

インターフェイスがio.Reader,io.Writerであれば動くようにしておけば、テストの時だけ入力を標準入力からファイルにしたりとか、出力をファイルから標準出力にしたりでき、テストがしやすくなる。

また、自前のio.Reader,io.Writerを作成し、入出力部分をモックに差し替えたりする事もできる。


