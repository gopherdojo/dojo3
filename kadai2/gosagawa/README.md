# kadai2

##io.Readerとio.Writerについて調べてみよう
- 標準パッケージでどのように使われているか
- io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

##1回目の宿題のテストを作ってみて下さい
- テストのしやすさを考えてリファクタリングしてみる
- テストのカバレッジを取ってみる
- テーブル駆動テストを行う
- テストヘルパーを作ってみる

# install
```
make
```

# usage

```
bin/imgconv [OPTION] dir_path
  -i string
    	変換対象の画像形式(jpeg|gif|png) (default "jpeg")
  -o string
    	変換語の画像形式(jpeg|gif|png) (default "png")
  -v	詳細なログを表示
```

