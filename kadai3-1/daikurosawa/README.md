# go-type-game

The Golang cli type game!

## Build

<!-- markdownlint-disable MD014 -->

```bash
$ go build main.go
```

<!-- markdownlint-enable MD014 -->

## Usage

```bash
$ ./main testdata/word.txt
> C++
C++ # type
Success!
> JavaScript
foo
Failed
> Go
:
:
Time out!
Result: 1
```

### Word list

Can use a one row one word text file.

```txt
JavaScript
Python
Java
C++
```

## Option

```text
Usage of type-game:
  -limit duration
        time limit (default 30s)
```

### Example

<!-- markdownlint-disable MD014 -->

```bash
$ ./main -limit 5s testdata/word.txt
```

<!-- markdownlint-enable MD014 -->

## GoDoc

<!-- markdownlint-disable MD014 -->

```bash
$ godoc -http=:6060
```

<!-- markdownlint-enable MD014 -->

You can access to read the documentation.

See this link:
[http://localhost:6060/pkg/github.com/gopherdojo/dojo3/kadai3-1/daikurosawa/](http://localhost:6060/pkg/github.com/gopherdojo/dojo3/kadai3-1/daikurosawa/)
