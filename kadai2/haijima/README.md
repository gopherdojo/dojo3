# 課題2

## 【TRY】io.Readerとio.Writer
* io.Readerとio.Writerについて調べてみよう
    * 標準パッケージでどのように使われているか
    * io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

### 標準パッケージでどのように使われているか

#### 定義
ioパッケージで定義されている。
インタフェース内で定義されているメソッドは、下記に示すようにそれぞれとてもシンプルである
``` go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
``` go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

#### 最小限
上で示すように保持するメソッドは最小限となっており、入出力の最小単位として機能している。

ioパッケージでは、`io.Closer`、`io.Seeker`と合わせてコンポジションが多く定義されている。

* `io.ReadWriter`
* `io.ReadCloser`
* `io.WriteCloser`
* `io.ReadWriteCloser`
* `io.ReadSeeker`
* `io.WriteSeeker`
* `io.ReadWriteSeeker`

#### 抽象的
`io.multi.go`の`multiReader`/`multiWriter`やbufioパッケージ配下の`Reader`/`Writer`なども用意されており、入出力に機能が付加されているものも多くある。

また、File descriptor、ソケット、Request Body、Http ConnectionなどもReader/Writerを実装しており、抽象かされた入出力インターフェースとして機能することで、入出力を行う関数に制限や個別ロジックなどが入り込みにくくなっている。


最小でかつ抽象的なため汎用性が高く、標準パッケージ内でもとても多く利用されている
[io.Reader - The Go Programming Language](https://golang.org/search?q=io.Reader)で確認すると、794件（テストやコメントも含まれているのでだいぶ雑ではある）


## 【TRY】テストを書いてみよう

* 1回目の宿題のテストを作ってみて下さい
    * テストのしやすさを考えてリファクタリングしてみる
    * テストのカバレッジを取ってみる
    * テーブル駆動テストを行う
    * テストヘルパーを作ってみる

### kadai-1での対応事項
* [x] テストのしやすさを考えてリファクタリングしてみる
* [x] テストのカバレッジを取ってみる
    * 87.2%
* [x] テーブル駆動テストを行う
* Exampleテストを作成
* testdataディレクトリにテスト用データを格納

### kadai-2での対応事項
* [x] テストヘルパーを作ってみる
    * `converter_test#assertConvert()`
