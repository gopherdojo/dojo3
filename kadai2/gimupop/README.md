#第2回課題

1. io.Readerとio.Writerについて調べてみよう
  1. 標準パッケージでどのように使われているか
    io.Readerとio.Writerは入出力を抽象化しているインタフェースです。
      * [io.Reader](https://golang.org/pkg/io/#Reader)
      * [io.Writer](https://golang.org/pkg/io/#Writer)
    以下の標準パッケージを始めいろいろなライブラリがこのインタフェースを実装しています。
      * [json](https://golang.org/pkg/encoding/json/)
      * [jpeg](https://golang.org/pkg/image/jpeg/)
      * [net](https://golang.org/pkg/net/)
      * [os](https://golang.org/pkg/os/)

  2. io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
   io.Readerとio.Writerは様々な入出力をインタフェースとして抽象化している。
   したがって、コード上でそのインターフェースを使用すると、その先がファイル, メモリ, データ送信等なのかを
   コード側で意識せずに使用することが出来る。
   またお決まりパターンとして標準パッケージ等でこのようなインターフェースを使用しておけば楽だしわかりやすい。

2. 第1回のテストを書く
  *テストのしやすさを考えてリファクタリングしてみる
  *テストのカバレッジを取ってみる
  *テーブル駆動テストを行う
  *テストヘルパーを作ってみる

  課題1のディレクトリに追加しときました。
