# ffconvert

`ffconvert` is a command to convert the file format under the specified directory.

By default, it converts from `JPEG` to `PNG`.

## How to try ffconvert

```shell
$ make build
$ ./ffconvert -J -p -f --compression-level=best-speed testdata/
$ ./ffconvert -P -g -f --num-colors=128 testdata/
$ ./ffconvert -G -j -f --quality=50 testdata/
$ make clean
```

## How to run the test

```shell
$ make test
```

## How to read GoDoc

```shell
$ make doc
```

## How to specify the input/output file format

**Input file format**

Specify the option with uppercase initials.

| Option | File Format |
| ---    | ---         |
| `-J`   | `JPEG`      |
| `-P`   | `PNG`       |
| `-G`   | `GIF`       |

**Output file format**

Specify the option with lowercase initials.

| Option | File Format |
| ---    | ---         |
| `-j`   | `JPEG`      |
| `-p`   | `PNG`       |
| `-g`   | `GIF`       |

For example, if you want to convert from GIF to JPEG, specify it like `-G -j`.

## How to specify the encoding option

As options for encoding, you can specify `--quality` for JPEG, `--num-colors` for GIF and `--compression-level` for PNG.

| Option                | Possible Values                           | Description                                    |
| ---                   | ---                                       | ---                                            |
| `--quality`           | 1 to 100                                  | JPEG Quality                                   |
| `--num-colors`        | 1 to 256                                  | Maximum number of colors used in the GIF image |
| `--compression-level` | default, no, best-speed, best-compression | PNG Compression Level                          |

## How to overwrite duplicate files

If the generated file name is duplicated, if you specify the `-f` option, it will overwrite the existing file without causing an error.

If `-f` is not specified, the following error will be displayed.

```shell
$ ./ffconvert testdata/images
File already exists: testdata/images/2018/07/001.png
```

`-f` overwrites them.

```shell
$ ./ffconvert -f testdata/images
Converted: "testdata/images/2018/07/001.jpg"
Converted: "testdata/images/2018/08/001.jpg"
$
```
