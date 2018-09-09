# Gopher Dojo 3 - Kadai 3-2
## Problem
* [x] Rangeアクセスを用いる
* [x] いくつかのゴルーチンでダウンロードしてマージする
* [x] エラー処理を工夫する
  * `golang.org/x/sync/errgourp` パッケージなどを使ってみる
* [x] キャンセルが発生した場合の実装を行う

## TODO
* [ ] プログレスバーを追加
* [ ] テスト追加

## Build
```
$go build -o kadai3-2 .
```

## Usage
```
Usage of ./kadai3-2:
  -p int
        Number of parallel processes (default 4)
```
