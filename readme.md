# Encode/Decode string to invisible text in Go

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://go.dev/)


## Install package

```shell
$ go get github.com/lequocbinh04/strinvsbl
```

## How to use
Simple demo in example/main.go:
````go
strinvsbl.Init()
x := strinvsbl.Encode("Can you see me?")
fmt.Println("Encode: '" + x + "'")
fmt.Println("Decode: '" + strinvsbl.Decode(x) + "'")
````

Output:
```shell
Encode: '‌󠁡󠁡󠁡󠁣󠁮󠁡󠁡󠁡󠁤󠁱󠁡󠁡󠁡󠁥󠁣󠁡󠁡󠁡󠁢󠁦󠁡󠁡󠁡󠁥󠁮󠁡󠁡󠁡󠁥󠁤󠁡󠁡󠁡󠁥󠁪󠁡󠁡󠁡󠁢󠁦󠁡󠁡󠁡󠁥󠁨󠁡󠁡󠁡󠁤󠁵󠁡󠁡󠁡󠁤󠁵󠁡󠁡󠁡󠁢󠁦󠁡󠁡󠁡󠁥󠁢󠁡󠁡󠁡󠁤󠁵󠁡󠁡󠁡󠁣󠁪'
Decode: 'Can you see me?'
```
