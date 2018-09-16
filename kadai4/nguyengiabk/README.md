# Gopher Dojo 3 - Kadai4
## Problem

おみくじAPIを作ってみよう
* [x] JSON形式でおみくじの結果を返す
* [x] 正月（1/1-1/3）だけ大吉にする
* [x] ハンドラのテストを書いてみる

## Build
```
$ go build -o kadai4 .
```

## Usage

Start server
```
$ ./kadai4
```

Client send request
```
$ curl localhost:8080
{"result":"凶"}
```
