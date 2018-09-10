# 課題3-2
- 分割ダウンロードを行う
  - Rangeアクセスを用いる
  - いくつかのゴルーチンでダウンロードしてマージする
  - エラー処理を工夫する
    - golang.org/x/sync/errgourpパッケージなどを使ってみる
  - キャンセルが発生した場合の実装を行う

## 学習ステップ
- コンテキストを理解する
  - 参考記事 contextパッケージとは何か https://deeeet.com/writing/2016/07/22/context/
    - 「キャンセルのためのシグナルの受け渡しの標準的なインターフェース」
    - NOTE ネットワーク周り（比較的死ぬ）に便利っぽいって事
  - 参考記事 contextのvalueに注意 https://deeeet.com/writing/2017/02/23/go-context-value/
- errgourpパッケージを理解する
  - "sync.ErrGroupで複数のgoroutineを制御する" https://deeeet.com/writing/2016/10/12/errgroup/
  - 非同期処理の同期時にエラーも取得可能にする
- Rangeアクセスを理解する
  - "「どの位置からどの位置までダウンロードする」と要求を送ることでその範囲分のファイルをダウンロード" https://qiita.com/codehex/items/d0a500ac387d39a34401
    - サーバー側の実装に因る
    - 範囲を指定してリクエストすればOK
- キャンセル実装を考える
  - とりま、エラーorキャンセルで全ルーチンを終了させ途中までのファイルを削除しようか

## TODO
- DONE 大枠設計
