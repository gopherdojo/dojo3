## 画像変換コマンド
### 仕様
- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション

### 要求
- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- 準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる

### 使い方
```bash
$ go build -o main
$ ./main -h
Usage of ./main:
  -d string
    	target directory of conversion
  -f string
    	extension of target file (jpg, jpeg, png, gif) (default "jpg")
  -t string
    	extension of after conversion (jpg, jpeg, png, gif) (default "png")
```