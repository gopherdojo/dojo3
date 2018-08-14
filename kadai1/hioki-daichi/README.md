# ffconvert

`ffconvert` は指定したディレクトリ以下のファイルフォーマットを変換します。

デフォルトでは `JPEG` から `PNG` に変換します。

## 画像形式の指定

画像形式を指定することもできます。

**変換前の指定方法**

| オプション | 画像形式 |
| ---        | ---      |
| `J`        | `JPEG`   |
| `P`        | `PNG`    |
| `G`        | `GIF`    |

**変換後の指定方法**

| オプション | 画像形式 |
| ---        | ---      |
| `j`        | `JPEG`   |
| `p`        | `PNG`    |
| `g`        | `GIF`    |

取りうるパターンが少ないためパターンをまとめてみますと、以下の通りです。

| パターン             | オプション指定方法 |
| ---                  | ---                |
| `JPEG` から `PNG` へ | `-J -p`            |
| `JPEG` から `GIF` へ | `-J -g`            |
| `PNG` から `JPEG` へ | `-P -j`            |
| `PNG` から `GIF` へ  | `-P -g`            |
| `GIF` から `JPEG` へ | `-G -j`            |
| `GIF` から `PNG` へ  | `-G -p`            |

## `-f` で上書き

変換後の名前が重複する場合、

- `-f` オプションを指定する場合、既存のファイルを上書きします。
- `-f` オプションを指定しない場合、エラーになります。

```shell
$ ./ffconvert test/images
File already exists: test/images/2018/07/001.png

$ ./ffconvert -f test/images
$
```

## `-v` でログ出力

`-v` を指定するとログを出力します。

```shell
$ ./ffconvert -v test/images
Skipped because the path is directory: "test/images"
Skipped because the path is directory: "test/images/2018"
Skipped because the path is directory: "test/images/2018/07"
Converted: "test/images/2018/07/001.png"
Skipped because the file is not applicable: "test/images/2018/07/001.png"
Skipped because the file is not applicable: "test/images/2018/07/002.png"
Skipped because the path is directory: "test/images/2018/08"
Converted: "test/images/2018/08/001.png"
Skipped because the file is not applicable: "test/images/2018/08/001.png"
Skipped because the file is not applicable: "test/images/2018/08/002.png"
Skipped because the file is not applicable: "test/images/2018/08/003.gif"

$ ./ffconvert test/images
$
```
