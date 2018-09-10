# 分割ダウンロードを行う
## 課題事項
- [ ] Rangeアクセスを用いる 
- [ ] いくつかのゴルーチンでダウンロードしてマージする
- [ ] エラー処理を工夫する
  - [ ] golang.org/x/sync/errgourpパッケージなどを使ってみる
- [ ] キャンセルが発生した場合の実装を行う 

 ## Usage
### build
```bash
% make build
```
 ### run
```bash
% ./matsuget http://example.com
```