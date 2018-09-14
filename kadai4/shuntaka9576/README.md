# Gopher道場課題 #4

## description
1/1 ~ 1/3 のみ大吉になります。  
それ以外は、(吉,中吉,小吉,末吉,凶,大凶)が返却されます。

### sever start
```bash
$ go run main.go
```

### usage
```bash
$curl localhost:8080/omikuji
{"result":"大凶"}
```

### test
```bash
make cov
```
