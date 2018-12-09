package main

import "fmt"

func main() {
	//fmt.Println("test")

	a,b := 1,2
	fmt.Println(a,b)
	a,b = b,a
	fmt.Println(a, b)
	//n, err := fmt.Println(a, b)
	//fmt.Println(n)
	//fmt.Println(err)

}
