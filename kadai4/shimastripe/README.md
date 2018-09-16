# おみくじAPIを作ってみよう

## 仕様

- JSON形式でおみくじ􏰁結果を返す
- 正月(1/1-1/3)だけ大吉にする
- ハンドラ􏰁テストを書いてみる

## 実行方法

 ```bash
 $ go run cmd/kadai4/main.go
 $ curl 'http://localhost:8080'

 # => {"fortune":"大吉"}
 ```

## Test

 ```bash
 $ go test -v

  === RUN   TestHandlerRandomly
  --- PASS: TestHandlerRandomly (0.00s)
  === RUN   TestHandlerWhenNewYear
  === RUN   TestHandlerWhenNewYear/NewYearCase
  === RUN   TestHandlerWhenNewYear/NewYearCase#01
  === RUN   TestHandlerWhenNewYear/NewYearCase#02
  --- PASS: TestHandlerWhenNewYear (0.00s)
      --- PASS: TestHandlerWhenNewYear/NewYearCase (0.00s)
      --- PASS: TestHandlerWhenNewYear/NewYearCase#01 (0.00s)
      --- PASS: TestHandlerWhenNewYear/NewYearCase#02 (0.00s)
  PASS
  ok      github.com/gopherdojo/dojo3/kadai4/shimastripe  0.018s
```

## Detail

- 時間は講義で受けたClock interfaceを用意して抽象化
- ハンドラ側のテストを用意
  - ランダムで決定する要素と決定しない要素を分けて書いた。ランダム要素側は不安定なので実質フレークテスト