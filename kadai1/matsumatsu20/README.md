# 課題1
## 次の仕様を満たすコマンドを作って下さい
- [x] ディレクトリを指定する
- [x] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- [x] ディレクトリ以下は再帰的に処理する
- [x] 変換前と変換後の画像形式を指定できる（オプション）

## 以下を満たすように開発してください
- [x] mainパッケージと分離する
  - mainとは別に`cli`, `converter` を作ってみました。
- [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う
  - 以下のパッケージを利用（準標準は利用なし）
  	- "errors"
  	- "flag"
	- "fmt"
	- "image"
	- "image/gif"
	- "image/jpeg"
	- "image/png"
	- "log"
	- "os"
	- "path/filepath"
	- "strings"
	- "github.com/gopherdojo/dojo3/kadai1/cli" (自作)
	- "github.com/gopherdojo/dojo3/kadai1/converter" (自作)
- [x] ユーザ定義型を作ってみる
  - CLIやConverterを作って見ました。またテストでも一部使っております。
- [x] GoDocを生成してみる
  - 生成してみました

## 使い方
### Build
```
$ make
```

### Run
```
$ ./ozu -i jpg -o gif images/
```

### Godoc
```
 $ godoc github.com/gopherdojo/dojo3/kadai1/matsumatsu20/
 $ godoc github.com/gopherdojo/dojo3/kadai1/matsumatsu20/cli/
 $ godoc github.com/gopherdojo/dojo3/kadai1/matsumatsu20/converter/
```
