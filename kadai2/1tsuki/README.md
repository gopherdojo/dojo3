Kadai1
====

Overview

## Assignment
次の仕様を満たすコマンドを作成せよ
- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）

以下を満たすように開発せよ
- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- 準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる

## Dependencies
- fmt
- strings
- io
- os
- flag
- path/filepath
- image
- image/jpeg
- image/png
- image/gif
- github.com/gopherdojo/dojo3/kadai1/1tsuki/sorcery

## Usage
Usage of imgconv

NAME
  imgconv -- convert image extensions from something to another

SYNOPSIS
  imgconv [-from] [-to] [directory ...]

DESCRIPTION
  -from string
    target image extension to convert. jpeg, png, gif are supported. default would be jpeg.
    
  -to string
    image extension to output. jpeg, png, gif are supported. default would be png.
