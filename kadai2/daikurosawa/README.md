# ＃3 Gopher道場

＃3 Gopher道場用のリポジトリです。

connpass: [https://mercari.connpass.com/event/95886/](https://mercari.connpass.com/event/95886/)

## kadai1

* 次の仕様を満たすコマンドを作って下さい
  * [x] ディレクトリを指定する
  * [x] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
  * [x] ディレクトリ以下は再帰的に処理する
  * [x] 変換前と変換後の画像形式を指定できる（オプション）
* 以下を満たすように開発してください
  * [x] main パッケージと分離する
  * [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う
  * [x] 準標準パッケージ：`golang.org/x` 以下のパッケージ
  * [x] ユーザ定義型を作ってみる
  * [x] GoDoc を生成してみる

### Build

<!-- markdownlint-disable MD014 -->

```bash
$ go build main.go
```

<!-- markdownlint-enable MD014 -->

### Usage

<!-- markdownlint-disable MD014 -->

```bash
$ ./main ~/Desktop/hoge/
```

<!-- markdownlint-enable MD014 -->

### Option

```text
Usage of ./main:
  -from string
        Input file extension. (default "jpg")
  -to string
        Output file extension. (default "png")
```

### GoDoc

<!-- markdownlint-disable MD014 -->

```bash
$ godoc -http=:6060
```

<!-- markdownlint-enable MD014 -->

You can access to read the documentation. See this link:
[http://localhost:6060/pkg/github.com/gopherdojo/dojo3/kadai1/daikurosawa/](http://localhost:6060/pkg/github.com/gopherdojo/dojo3/kadai1/daikurosawa/)

## kadai2

### Test

* [x] リファクタ
* [x] テストカバレッジ
* [x] テーブル駆動テスト
* [x] テストヘルパー

#### カバレッジ

```html
<option value="file0">github.com/gopherdojo/dojo3/kadai2/daikurosawa/cli/cli.go (97.1%)</option>
<option value="file1">github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/convert.go (82.8%)</option>
<option value="file2">github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/gif/gif.go (100.0%)</option>
<option value="file3">github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/jpg/jpg.go (100.0%)</option>
<option value="file4">github.com/gopherdojo/dojo3/kadai2/daikurosawa/convert/png/png.go (100.0%)</option>
```

### io.Readerとio.Writerについて調べてみよう

#### 標準パッケージでどのように使われているか

`io.Reader`, `io.Writer` は入力, 出力を抽象化しているインタフェース。

* [io.Reader](https://golang.org/pkg/io/#Reader)

```go
type Reader interface {
        Read(p []byte) (n int, err error)
}
```

* [io.Writer](https://golang.org/pkg/io/#Writer)

```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

色々な標準パッケージがこのインタフェースを実装していたり、引数として使えるように作られている。

* [os](https://golang.org/pkg/os/)
* [json](https://golang.org/pkg/encoding/json/)
* [jpeg](https://golang.org/pkg/image/jpeg/)
* [net](https://golang.org/pkg/net/)

#### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

様々な入力, 出力がインタフェースとして抽象化されており、コード上ではそのインターフェースに依存しているため
その先がファイル, メモリ, データ送信etc なのかをコード側で意識せずにテストの中で切り替えを行うことができる。

自分でインタフェースを実装することによりあらゆる入力, 出力に対して自分のコードを提供することも可能になる。
