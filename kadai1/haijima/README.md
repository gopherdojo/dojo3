# 課題1 画像変換コマンドを作ろう
```
* 次の仕様を満たすコマンドを作って下さい
  * ディレクトリを指定する
  * 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
  * ディレクトリ以下は再帰的に処理する
  * 変換前と変換後の画像形式を指定できる（オプション）
* 以下を満たすように開発してください
  * mainパッケージと分離する
  * 自作パッケージと標準パッケージと準標準パッケージのみ使う
    * 準標準パッケージ：golang.org/x以下のパッケージ
  * ユーザ定義型を作ってみる
  * GoDocを生成してみる
```

# 使い方
## Build
```
$ make
```

## Usage
```
Usage of conv:
conv [-i srcExts] [-o destExt] [-w] [--dry-run|-q] [-s] [directory]
  -dry-run
        Dry run mode
  -i string
        Input extension. (default jpg|jpeg)
  -o string
        Output extension. (default png)
  -q    Quiet mode. Suppress print
  -s    Matches file extension case-sensitively
  -w    If converted file has already existed, Overwrite old files.
```

# 課題の内容
## mainパッケージと分離する
imgconvパッケージを作りました

## 自作パッケージと標準パッケージと準標準パッケージのみ使う
使用したパッケージは以下の通りです
* 標準パッケージ
  * bytes
  * errors
  * flag
  * fmt
  * image
  * image/jpeg
  * image/png
  * image/gif
  * io
  * os
  * path
  * path/filepath
  * reflect
  * strings
  * testing
  * time
* 自作パッケージ
  * github.com/haijima/go-imgconv/imgconv

## ユーザ定義型を作ってみる
下記の型を作成しました
* Converter
* Cli
* Argument
* Option

## GoDocを生成してみる
* 公開される型・関数へのコメントを記述しました
* Exampleを用いて公開される関数の使用例を記述しました
* `imgconv/doc.go`にパッケージの概要を記述しました

# 気になるところ・見ていただきたいところ・メモ
パッケージの切り方や、ファイルの分け方、インターフェースの切り方のちょうどいい粒度がわからず、色々切りすぎてしまったように感じています。
この辺Goならでは、のような考え方などあれば教えていただきたいです。

テスト、書いてから次回の課題になっていることに気づきました。一応提出しておきます。

