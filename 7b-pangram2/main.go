package main

import (
	"fmt"
	"unicode"
)

func main() {
	for _, tc := range []string{
		"the quick brown fox jumps over the lazy dog",
		"the quick brown fox jumped over the lazy dog",
	} {
		if isPangram(tc) {
			fmt.Println("pangram")
		} else {
			fmt.Println("not pangram")
		}
	}
}

func isPangram(s string) bool {
	mm := mm()
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) {
			t := unicode.ToLower(r)
			if _, ok := mm[t]; ok {
				mm[t] = true
			} else {
				return false
			}
		}
	}
	for _, v := range mm {
		if !v {
			return false
		}
	}
	return true
}

func mm() map[rune]bool {
	m := make(map[rune]bool, 28)
	for i := 'a'; i <= 'z'; i++ {
		m[i] = false
	}
	return m
}
