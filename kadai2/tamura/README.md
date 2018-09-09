# 第2回課題

# io.Readerとio.Writerについて調べてみよう
## どのように使用されているのか。

* 読み込みでio.Reader、書き込みでio.Writerが利用していた。
→　os.Stdin、os.File、net.Conn、bytes.Buffer

## io.Readerとio.Writerがあることで、どういう利点があるのか。

* 拡張性の実装
様々な出力先を抽象化出来る。少ない改修の規模で拡張可能。
* テストで出力先変更
io.Writerの引数抽象化で、標準出力に書き込み、テストはメモリのバッファに書き込むの切り替えができる
* コードを提供可能
io.Writer,io.Readerが使われているので、活用の幅が広い。

# テスト

作成中
