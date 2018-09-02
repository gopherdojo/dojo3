# 画像変換コマンド
## kadai2-1
### io.Readerとio.Writerについて
#### 標準パッケージでどのように使われているか
入出力が発生するほぼ全ての箇所では`io.Reader`, `io.Writer`が実装されており、パッケージに応じて`io.Closer`や`io.Seeker`も合わせて実装されている。
例えば、`os.File`は上記4つのinterfaceを全て満たしており、`os.Stdin`は`io.ReadCloser`を実装している。

#### io.Readerとio.Writerがあることでどういう利点があるのか
利用者側がio.Readerや口とio.Writerの口を用意すればいいので、同じコードで様々なユースケースに対応できる。
また、今回のように本番ではファイルを書き出しつつ、テストではbufferに書き出すようにする、といったこともできる。
再利用を考えていないコードでもテストのためにinterfaceの引数を用意するのはちょっとつらいかも？

## kadai2-2
- テストを書いてみよう
  - リファクタリング
  - カバレッジ取得
    ```
    $ go test ./... -coverprofile cover.out; go tool cover -func cover.out
    ```
  - テーブル駆動テスト
  - テストヘルパー

---

## kadai1
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

