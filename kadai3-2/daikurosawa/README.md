# go-parallel-download

The Golang file parallel download command line tool.

## Build

<!-- markdownlint-disable MD014 -->

```bash
$ go build main.go
```

<!-- markdownlint-enable MD014 -->

## Usage

<!-- markdownlint-disable MD014 -->

```bash
$ ./main [directory path to save file] [download file url]
```

<!-- markdownlint-enable MD014 -->

### Example

```bash
$ ./main /Dir/path https://example.com/foo.png
> Content-Length: 74121
> 59300-74120 download success
> 29650-44474 download success
> 44475-59299 download success
> 14825-29649 download success
> 0-14824 download success
> merge file...
> complete!
> output: /Dir/path/foo.png
```

## Option

```text
Usage of range-get:
  -parallel int
        parallel count (default 5)
  -timeout duration
        time out (default 5m0s)
```

## GoDoc

<!-- markdownlint-disable MD014 -->

```bash
$ godoc -http=:6060
```

<!-- markdownlint-enable MD014 -->

You can access to read the documentation.

See this link:
[http://localhost:6060/pkg/github.com/gopherdojo/dojo3/kadai3-2/daikurosawa/](http://localhost:6060/pkg/github.com/gopherdojo/dojo3/kadai3-2/daikurosawa/)
