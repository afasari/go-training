package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "The quick brown fox jumps over the lazy dog"
	b := searchword(a, "test")
	fmt.Println(b)
}

func searchword(s string, kw string) bool {
	//for i:=0; i < len(s); i++{
	//
	//}

	if strings.Contains(s, kw) {
		return true
	}
	return false
}
