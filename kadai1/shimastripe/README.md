# Img Converter

## 仕様

- ディレクトリを指定する
- 指定したディレクトリ以下􏰁JPGファイルをPNGに変換
- ディレクトリ以下􏰀再帰的に処理する
- 変換前と変換後􏰁画像形式を指定できる

## requirement

- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージ􏰁み使う
- 準標準パッケージ:golang.org/x以下􏰁パッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる
 
 ## 実行方法

 ```bash
 $ go run cmd/kadai1/main.go -from jpg -to gif ./data
 ```

 ## document
 
 ```bash
 $ godoc github.com/gopherdojo/dojo3/kadai1/shimastripe
 $ godoc github.com/gopherdojo/dojo3/kadai1/shimastripe/imgconv
 ```

 ## Detail

- 画像形式は `jpg, gif, png` を相互で変換可能です
  - 拡張子の大文字小文字は区別せず読み込めます
- 標準出力、入力をCLI structを介して出力するように書くことでテストをしやすくしました
- `io`を介するところはinterface型で受けてテストを書きやすくなるようにしました
  - しかし、画像の変換のテストって整合するのかイマイチわからず書いてません
- 画像のformatに関してはImgConverter structのpackage内関数で処理することで外部から使う際依存しないようにしました