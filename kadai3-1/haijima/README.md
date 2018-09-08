# 課題3-1

## 要求

タイピングゲームを作ろう
* [x] 標準出力に英単語を出す（出すものは自由）
* [x] 標準入力から1行受け取る
* [x] 制限時間内に何問解けたか表示する

`main()`で定義した文字列が問題として順に表示される。
タイプした文字列は1行ずつ評価される。

## How to build
```
$ make
```

## Result
```
$ ./type
--- Typing Game Start!! ---
hello
hello
world
world
gopher
gopher
hello
hola
world
mundo
gopher


--- Time Over!! ---
--- Result ---
Typed 3 words
Succeed Rate 60.0%
--------------
```

## test
```
$ go test --cover ./...
?   	github.com/gopherdojo/dojo3/kadai3-1/haijima	[no test files]
ok  	github.com/gopherdojo/dojo3/kadai3-1/haijima/typing	0.109s	coverage: 100.0% of statements
```
