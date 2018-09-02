# Gopher Dojo 3 - Kadai 2

## Problem

### Kadai 1
* 次の仕様を満たすコマンドを作って下さい
  - ディレクトリを指定する
  - 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
  - ディレクトリ以下は再帰的に処理する
  - 変換前と変換後の画像形式を指定できる（オプション）

* 以下を満たすように開発してください
  - mainパッケージと分離する
  - 自作パッケージと標準パッケージと準標準パッケージのみ使う
    - 準標準パッケージ：golang.org/x以下のパッケージ
  - ユーザ定義型を作ってみる
  - GoDocを生成してみる

### Kadai 2

* io.Readerとio.Writerについて調べてみよう
  - [Answer](Answer.md)に回答を書きました。

* 1回目の宿題のテストを作ってみて下さい
  - テストのしやすさを考えてリファクタリングしてみる
  - テストのカバレッジを取ってみる
  - テーブル駆動テストを行う
  - テストヘルパーを作ってみる

## Command details
* JPEG, PNG, GIFを対応しました。
* デコード出来ない場合はログを出して、次の処理へ進みます。
* JPEGのQuaility, GIFのNumColorsのパラメーターが指定できるようにしました。
* GoDocのExampleを作りました

## Build
```
$go build -o kadai2 .
```

## Run
```
$./kadai2 [options] [directories]
```

### Options
```
-i string
    Input file type (default "jpg")

-o string
    Output file type (default "png")

-q int
    JPG Quality, ranges from 1 to 100, (only used for encoding jpg) (default 100)

-n int
    Maximum number of colors, ranges from 1 to 256, (only used for encoding gif) (default 256)    
```

### Examples
```
$./kadai2 -i jpg -o png testdata
$./kadai2 -i jpg -o png testdata1 testdata2
$./kadai2 -i png -o gif -n 100 testdata
$./kadai2 -i png -o jpg -q 50 testdata
```

## Godoc
```
$godoc -http=:6060
```
You can read the documentation about converter package at:
`http://localhost:6060/pkg/github.com/gopherdojo/dojo3/kadai2/nguyengiabk/converter`

## Test coverage
```
$go test -coverprofile=coverage.out ./...
$go tool cover -func=coverage.out
$go tool cover -html=coverage.out
```
