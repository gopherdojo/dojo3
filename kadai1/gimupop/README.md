#第1回課題 Img Converter
 ## 仕様
* ディレクトリを指定する
* 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
* ディレクトリ以下は再帰的に処理する
 ## requirement
* mainパッケージと分離する
* 自作パッケージと標準パッケージと準標準パッケージのみ使う
* 準標準パッケージ：golang.org/x以下のパッケージ
* ユーザ定義型を作ってみる
* GoDocを生成してみる
 
 ## 実行方法
  ```bash
 $ go run kadai1/gimupop/main/main.go -from jpg -to gif -target ./kadai1/gimupop/target   
 ```
  ## document
 
 ```bash
 $ godoc github.com/gopherdojo/dojo3/kadai1/gimupop
 $ godoc github.com/gopherdojo/dojo3/kadai1/gimupop/imgconv
 ```
 