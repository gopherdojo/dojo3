# 第1回課題　画像変換コマンドの作成

## ディレクトリ指定のみ  ( jpg→ png)

```
go run main.go images
```
imagesディレクトリ内のjpgファイルがpngファイルに変換され作成される

## 画像形式を指定する (ex. gif → png)

```
go run main.go -f gif -t png images
```
imagesディレクトリ内の -f オプションで指定されたフォーマットのファイルが -t オプションで指定されたファイルに変換される。

## 補足
今回は jpg, jpeg, png, gif のみ変換対応
