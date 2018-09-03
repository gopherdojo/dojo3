# Gopher道場#3 課題2

## io.Readerとio.Writerについて調べてみよう

#### io.ReaderはRead()が定義されているインタフェース

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```
引数で指定したバイト列を読み込み、読み込んだバイト数とエラーを返す。


#### io.WriterはWrite()が定義されているインタフェース

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
引数で指定したバイト列を書き込み、書き込んだバイト数とエラーを返す


#### 標準パッケージではどのように使われているか

    - ファイル出力(os.File) ファイルを読み込む/書き込む

    - 画面出力(os.Stdout) 画面の入力を受け付ける/表示する
    
    - バッファ(bytes.Buffer) バッファから読み込む/バッファに溜める

#### io.Readerとio.Writerがあることで、どういう利点があるのか具体例を挙げて考えてみる

入出力の機能が抽象化されているので、例えば標準出力に出していたものをファイルに書く場合などに少ない修正で実装を差し替えることができる。また、標準出力とファイルに両方書き出したい場合などにも同様に扱うことができ、差し替えや合成が容易にできる。



## 1回目の宿題のテストを作ってみてください
* テストのしやすさを考えてリファクタリングしてみる
実ファイルを扱わなくてもテストできるように、io.Readerを受け取ってbytes.Bufferを返すようにした

* テストのカバレッジをとってみる

```
ᐅ go test -coverprofile=profile ./convertImage
ok _/Users/yamadanaoyuki/Documents/git/dojo3/kadai2/chokkoyamada/convertImage	0.576s	coverage: 90.0% of statements
```

* テーブル駆動テストを行う
TestConvert()をtestFlagsという変数でテーブルテストしてみた

* テストヘルパーを作ってみる
これがよくわからず、未対応


