# kadai2
## kadai2-1
### 標準パッケージでどのように使われているのか
#### 概要  
読み込みに関わる処理にio.Reader  
書き込みに関わる処理にio.Writerが利用されている。  
io.Readerインターフェースは、以下のメソッドを生やすことで実装可能  
```go
Read(p []byte)(n int, err error)
```
io.Writerインターフェースは、以下のメソッドを生やすことで実装可能
```go
Write(p []byte)(n int, err error)
```
#### どのように使われているか
* 標準出力  
os.Stdin(implements io.Reader)
* ファイル入力  
os.File(implements io.Reader,io.Writer)
* インターネット通信  
net.Conn(implements io.Reader,io.Writer)
* メモリのバッファ  
bytes.Buffer(implements io.Writer)

### io.Readerとio.Writerがあることで、どういう利点があるのか具体的に挙げて考えてみる
* 拡張性の高い実装が可能  
様々な出力先を抽象化出来るので、
標準出力に書き込んでいたCLIツールをファイル出力に変更したい場合など、少ない改修の規模で拡張可能
* テスト時に出力先変更が可能  
io.Writerで引数を抽象化することで、  
実際に動かすときは標準出力に書き込み、テスト時はメモリのバッファに書き込むといった切り替えが可能  
* あらゆる出力機能をに対して自分のコードを提供可能  
golangのライブラリの出力機構として、io.Writer,io.Readerが使われていることから  
自分でio.Writer,io.Readerを実装した構造体を定義した際に、活用の幅が広い。
## kadai2-2
* テストのしやすさを考えてリファクタリングしてみる  
  * 所感  
    * Cliパッケージは、他パッケージの構造体をモックで差し込める設計にするべきだった  
      Cliパッケージの処理部分が、main関数のような位置づけになってしまい、  
      パッケージに着目したテストが書けなかった。
* テストのカバレッジを取ってみる
  * cliパッケージ  
  76.5%
  * converterパッケージ  
  91.7%
  * imagetypes  
  TODO
  * 所感  
    * テストが書き辛いケースが幾つかあった
      * 画像のエンコード・デコードに失敗する処理  
    * カバーできていない箇所についてのアプローチ  
      網羅できていないことを把握して、手動でテストを実施するとか?  
      そもそもユニットテストでどこまで動作確認の保証ができるのか?  
    * その他  
      結果をhtml形式で出力するとどこが網羅されていないか一目でわかり、便利。
* テーブル駆動テストを行う
  * 所感  
  あらゆる引数のパターンを網羅的にテストすることができ、動作確認で便利。  
  サブテスト機能を使うと、テーブルのどのケースで落ちたかが分かりやすかった。
* テストヘルパーを作ってみる
  * 所感  
  テストが落ちたとき、test.goコードの何行目で落ちたか分かるので便利。
# Kadai1
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
## Build
make build
## Command Line Options
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
