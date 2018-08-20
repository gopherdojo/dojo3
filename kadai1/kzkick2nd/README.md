# 課題1 画像変換コマンドを作ろう

指定ディレクトリ以下のjpgをpngに変換します。

変換前後フォーマットをオプション指定できます。(jpg|png)

```
$ go run main.go [-i (jpg|png)] [-o (jpg|png)] <path/to/dir>
```

## TODO
- コマンドライン引数バリデーション作成
- テスト作成
