package main

import (
	"fmt"
	"github.com/lequocbinh04/strinvsbl"
)

func main() {
	strinvsbl.Init()
	x := strinvsbl.Encode("Can you see me?")
	fmt.Println("Encode: '" + x + "'")
	fmt.Println("Decode: '" + strinvsbl.Decode(x) + "'")
}
