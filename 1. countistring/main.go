package main

import (
"fmt"
"strconv"
)

func main() {
	sol := solution(3, 2)
	fmt.Printf("%d\n", sol)
}

func solution(a, b int) int {
	c := a * b
	d := int64(c)
	e := strconv.FormatInt(d, 2)
	var count int
	for i := 0; i < len(e); i++ {
		if string(e[i]) == "1" {
			count++
		}
	}
	return count
}
