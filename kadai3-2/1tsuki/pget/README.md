# 分割ダウンロードを行う
- [x] Rangeアクセスを用いる
- [x] いくつかのゴルーチンでダウンロードしてマージする
- [x] エラー処理を工夫する
- [x] golang.org/x/sync/errgourpパッケージなどを使ってみる
- [x] キャンセルが発生した場合の実装を行う

## 提供機能
- 処理中断時に平行リクエストを終了
- ダウンロード再開機能

## How to run
```bash
% cd main 
% make build
% ./pget -p 10 -t 60 https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.6.4.tar.gz
```
