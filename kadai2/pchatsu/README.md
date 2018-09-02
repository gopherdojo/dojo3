# 課題

```
1. io.Readerとio.Writerについて調べてみよう
  1. 標準パッケージでどのように使われているか
  2. io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
2. 第一回のテストを書く

```

## 1. io.Readerとio.Writerについて調べよう

### net/http/request.goでの使われ方

`Request`構造体はhttpリクエストを抽象化している。  
`func NewRequest(method, url string, body io.Reader) (*Request, error)` [(github)](https://github.com/golang/go/blob/7a6fc1f30bc7339726cd3f93f96be3e0d36ff7cb/src/net/http/request.go#L792) 
では、
Request Bodyを`io.Reader`で受け取り、`Request`構造体を生成している。  
`func (r *Request) Write(w io.Writer) error` [(github)](https://github.com/golang/go/blob/7a6fc1f30bc7339726cd3f93f96be3e0d36ff7cb/src/net/http/request.go#L494)
では、requestの内容を`io.Writer`に書き込む処理を担当している。
HTTP通信の際は、張られたコネクション上で書き込み内容が送信されている。

### どういう利点があるのか具体例を挙げて考えてみる

以下の２点の利点について、`net/http/request.go` での例を元に考察してみる。

#### モジュラリティ

* 似たような処理を１つの関数で共通化できる
  * 個別の処理が必要な場合は、型アサーションにより吸収することもできる
* 内部ロジックの変更による外部影響が限定的
* 担当ロジックに集中できる

という利点があげられる。
具体的に`func NewRequest(method, url string, body io.Reader) (*Request, error)`では
Bodyは文字列で渡される場合やバイナリデータを送信する場合などがあるが、
`io.Reader`で扱うことでそれらの差異を吸収でき、ひとつの関数だけで扱うことができている。
また、`func NewRequest`で新たに処理を追加しようとしても、現在の利用者に影響が少ない形で修正が可能になっている。

`func (r *Request) Write(w io.Writer) error`では、書き込み処理の抽象化を行っている。`Request`側では、どのような通信が行われているのかについて考慮する必要なく、担当のロジックに集中できているといえる。


#### テスタビリティ
* テストが書きやすい

`request_test.go`[(github)](https://github.com/golang/go/blob/7a6fc1f30bc7339726cd3f93f96be3e0d36ff7cb/src/net/http/request_test.go)では、HTTP通信をすることなく、単体テストを実行することができている。


## 2. kadai1のテストを書く

kadai1のテストを作成しました。

### 疑問点/詰まったところ

https://budougumi0617.github.io/2018/08/19/go-testing2018/ の資料を見て、テストに使いたい画像ファイルなどを `.golden` 形式でおき使ってみました。
ロジックを拡張子に依存させた処理があるところはうまくテストを書くことができないなど、不便な点を感じたため、`.golden`形式で扱うとどういう利点があるのかをちゃんと理解しておきたいと思いました。