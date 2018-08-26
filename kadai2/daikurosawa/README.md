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

## Build

<!-- markdownlint-disable MD014 -->

```bash
$ go build main.go
```

<!-- markdownlint-enable MD014 -->

## Usage

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

## GoDoc

<!-- markdownlint-disable MD014 -->

```bash
$ godoc -http=:6060
```

<!-- markdownlint-enable MD014 -->

You can access to read the documentation. See this link:
[http://localhost:6060/pkg/github.com/gopherdojo/dojo3/kadai1/daikurosawa/](http://localhost:6060/pkg/github.com/gopherdojo/dojo3/kadai1/daikurosawa/)
