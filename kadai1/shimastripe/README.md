# Img Converter

## 仕様1

- ディレクトリを指定する
- 指定したディレクトリ以下􏰁JPGファイルをPNGに変換
- ディレクトリ以下􏰀再帰的に処理する
- 変換前と変換後􏰁画像形式を指定できる

## requirement1

- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージ􏰁み使う
- 準標準パッケージ:golang.org/x以下􏰁パッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる
 
 ## 実行方法1

 ```bash
 $ go run cmd/kadai/main.go -from jpg -to gif ./data
 ```

 ## document1
 
 ```bash
 $ godoc github.com/gopherdojo/dojo3/kadai1/shimastripe
 $ godoc github.com/gopherdojo/dojo3/kadai1/shimastripe/imgconv
 ```

 ## Detail1

- 画像形式は `jpg, gif, png` を相互で変換可能です
  - 拡張子の大文字小文字は区別せず読み込めます
- 標準出力、入力をCLI structを介して出力するように書くことでテストをしやすくしました
- `io`を介するところはinterface型で受けてテストを書きやすくなるようにしました
  - しかし、画像の変換のテストって整合するのかイマイチわからず書いてません
- 画像のformatに関してはImgConverter structのpackage内関数で処理することで外部から使う際依存しないようにしました

## 課題2-1

### io.Reader/io.Writerの使われ方

- json
- bytes.Buffer
- bufio.Reader
- os.File
- image
- jpeg
- net.TCPConn
- png
- base64

### どのように使われていてどのような利点があるか

- 講義でも言われたとおり、ふるまいのみに抽象化していることで、相互の入れ替えが用意に可能である。特にReader/Writerはファイルの読み込みや変換をすべてストリームに抽象化することでデータが流れるストリームに対する読み込み・書き込みと捉えることで標準入出力からTCPコネクションまでbyteのストリームと相互利用することが可能になっている。

## 課題2-2
- [x] テストを書く io.Reader/Writerで抽象化することでtestdataなしでテストできるようにした
- [ ] テーブル駆動テスト
  - 利用するEncode関数がパッケージごと変わってしまってファイルなしのテーブル駆動テストがうまく書けない...。switch文で分岐するようなencodeをテスト内に作ると本体コードでやってることと全くおなじになりそれはテストとして正しいのか疑問
- [ ] テストヘルパー
  - テーブル駆動テストが書けてないため使えてない
