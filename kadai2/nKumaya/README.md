# 第2回課題　テストを書いてみよう

## io.Readerとio.Writerについて調べてみよう
Go1.10においては以下のように定義されている
```
type Reader interface {
        Read(p []byte) (n int, err error)
}

type Writer interface {
        Write(p []byte) (n int, err error)
}
```
入出力に関わるインタフェースを抽象化することによりコードから冗長性を減らしシンプルに保つことができる。
入力元や出力先がファイルなのか標準出力なのかローカルのものなのかリモートにあるものかなど様々にあるが、同じメソッドを用いることができ、コードの切り替えが容易になるというメリットがある。

## コマンドツールの使い方
## ディレクトリ指定のみ  ( jpg→ png)

```
go run main.go images
```
imagesディレクトリ内のjpgファイルがpngファイルに変換され作成される

## 画像形式を指定する (ex. gif → png)

```
go run main.go -f gif -t png images
```
imagesディレクトリ内の -f オプションで指定されたフォーマットのファイルが -t オプションで指定されたファイルに変換される。

## 補足
今回は jpg, jpeg, png, gif のみ変換対応
