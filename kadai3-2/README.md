
## 課題

```
分割ダウンロードを行う
Rangeアクセスを用いる
いくつかのゴルーチンでダウンロードしてマージする
エラー処理を工夫する
golang.org/x/sync/errgourpパッケージなどを使ってみる
キャンセルが発生した場合の実装を行う
```

## 実装

* [x] `signal.Notify` を使ってキャンセルのシグナルをハンドリングする
* [x] `context.Context`, `context.WithCancel` を使って、キャンセルの情報が下流タスクに伝播させる
* [x] `errgroup.WithContext`を使って、並行処理時に`cancel()`が呼ばれたときやエラーが起こったときをハンドリングする
  * [x] `ctxhttp`を使ってhttpリクエスト中のキャンセルに対応
  