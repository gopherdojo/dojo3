# Gopher Dojo 3 - Kadai 1（画像コンバートCLIツール作成）

## Problem
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


## Command details
* JPEGからPNGへの変換することができます。
* フォルダを指定し、複数ファイルを一括変換できます。
* GoDocのExampleを作りました


## Build
```
$make
```

## Run
```
convertImage [options]
```

### Options
```
-s string
	Convert Directory

-i string
    Input file type (default "jpg")

-o string
    Output file type (default "png")
```


### Examples
```
$./bin/convertImage -s _data/
$./bin/convertImage -i jpg -o png -s _data/
```


## Godoc
```
$godoc -http=:6060
```
You can read the documentation about converter package at:
`http://localhost:6060/pkg/gopher-dojyo/kadai1/src/convert/image/`


## TODO
- Goっぽさがない（構造体などをうまく使えてない）
- テスト作成
- Godoc作成
- 変換対象を増やす(jpg, png, gitの相互変換)
- 変換後のpng軽量化
- 並列処理が重い…
