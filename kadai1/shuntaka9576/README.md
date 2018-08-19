# Kadai1
Gopherdojo kadai1 tool  
Convert png and jpeg files to jpeg and png.  
# Assignment
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
# Build
make build
# Command Line Options
## Description
### -version
Display the version of ImageConverter.
### -f,-from
Specifies the original file extention.
### -t,-target
Specifies want to convert file extention.
### no-flag
Specifies convert images directory.
## SYNOPSIS
imageConverter [-f] [-t] [directory]
# 注意点
既に何点かバグを発見してます。  
とりあえず暫定版として、提出いたします。  
時間を見つけて修正予定です。
