# io.Readerとio.Writerについて調べてみよう
## 標準パッケージでどのように使われているか
* よく使われている。ほとんどの標準パーケージに使われている。
  - io.Readerを使っている[ファイルリスト](ioReaderList.md)
  - io.Writerを使っている[ファイルリスト](ioWriterList.md)

* 引数として関数に渡し、入力および出力を抽象化する。
* 出力および入力が必要なところはほとんどio.Readerとio.Writerを使っている。

### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
net/http/request.goファイルに下記の関数がある

```
// If body is of type *bytes.Buffer, *bytes.Reader, or
// *strings.Reader, the returned request's ContentLength is set to its
// exact value (instead of -1), GetBody is populated (so 307 and 308
// redirects can replay the body), and Body is set to NoBody if the
// ContentLength is 0.

func NewRequest(method, url string, body io.Reader) (*Request, error)
```

  * bodyは色んな型から作成できる。汎用性が高い。その関数のコメントにも書かれているが、bytes.Bufferやbytes.Readerなど色んな型をbodyに渡せる。

  * モックが出来て、テストしやすい。下記のテストコードを見ると、bodyのモックは色々の方法があるとわかる。
  ```
  tests := []struct {
		r    io.Reader
		want int64
	}{
		{bytes.NewReader([]byte("123")), 3},
		{bytes.NewBuffer([]byte("1234")), 4},
		{strings.NewReader("12345"), 5},
		{strings.NewReader(""), 0},
		{NoBody, 0},

		// Not detected. During Go 1.8 we tried to make these set to -1, but
		// due to Issue 18117, we keep these returning 0, even though they're
		// unknown.
		{struct{ io.Reader }{strings.NewReader("xyz")}, 0},
		{io.NewSectionReader(strings.NewReader("x"), 0, 6), 0},
		{readByte(io.NewSectionReader(strings.NewReader("xy"), 0, 6)), 0},
	}
	for i, tt := range tests {
		req, err := NewRequest("POST", "http://localhost/", tt.r)
  ```

  * 入力と出力ロジック（bufferの初期化など）を切り出せて、コードの可読性が上がる。
