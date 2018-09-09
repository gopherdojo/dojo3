# 課題3-2
## 要求
分割ダウンロードを行う
* [x] Rangeアクセスを用いる
* [x] いくつかのゴルーチンでダウンロードしてマージする
* [x] エラー処理を工夫する
    * golang.org/x/sync/errgourpパッケージなどを使ってみる
* [x] キャンセルが発生した場合の実装を行う

### できてないこと
* [ ] テスト


## How to build
```
$ make
```


## How to run
```
$ ./gget [-n num] [-o outputdir] url
```
