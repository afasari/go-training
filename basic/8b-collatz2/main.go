package main

import (
	"fmt"
	"os"
	"strconv"
	//"time"
)

func linear_collatz(init_num int) {
	num := 0
	count := 0
	highest_number := 0
	highest_count := 0
	for i := 2; i < init_num; i++ {
		num = i
		count = 0
		for {
			if num == 1 {
				break
			}

			num1 := 0
			if num%2 == 0 {
				num1 = num / 2
			} else {
				num1 = ((3 * num) + 1)
			}
			num = num1
			count = count + 1
		}
		if count > highest_count {
			highest_count = count
			highest_number = i
		}
	}
	fmt.Println(highest_count)
	fmt.Println(highest_number)
}

func cached_linear_collatz(init_num int) {
	m := make(map[int]int)
	num := 0
	count := 0
	highest_number := 0
	highest_count := 0
	for i := 2; i < init_num; i++ {
		num = i
		count = 0
		for {
			if num == 1 {
				m[i] = count
				break
			}

			num1 := 0
			if m[num] != 0 {
				count = count + m[num]
				m[i] = count
				break
			} else {
				if num%2 == 0 {
					num1 = num / 2
				} else {
					num1 = ((3 * num) + 1)
				}
			}
			num = num1
			count = count + 1
		}
		if count > highest_count {
			highest_count = count
			highest_number = i
		}
	}
	fmt.Println(highest_count)
	fmt.Println(highest_number)
}

func main() {
	init_num := 1
	method := ""
	if len(os.Args) == 3 {
		var arg_num, _ = strconv.Atoi(os.Args[1])
		init_num = arg_num
		method = os.Args[2]
	} else if len(os.Args) == 2 {
		var arg_num, _ = strconv.Atoi(os.Args[1])
		init_num = arg_num
	} else {
		fmt.Println("Invalid amount of arguments")
	}

	if method == "" || method == "linear" {
		linear_collatz(init_num)
	} else if method == "cached" {
		cached_linear_collatz(init_num)
	}
}
