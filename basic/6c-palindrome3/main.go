package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(input string) bool {
	input = strings.TrimSpace(input)
	chars := []string{"]", "^", "\\\\", "[", ".", "(", ")", "-", "?", " "}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	input = re.ReplaceAllString(input, "")

	fmt.Println(input)
	for i := 0; i < len(input)/2; i++ {
		if strings.EqualFold(string(input[i]), string(input[len(input)-i-1])) {
			return true
		}
		//if input[i] != input[len(input)-i-1] {
		//	return false
		//}
	}
	return false
}

func main() {
	fmt.Println(isPalindrome("anna"))
	fmt.Println(isPalindrome("not a palindrome"))
	fmt.Println(isPalindrome("Was it a car or a cat I saw?"))
	fmt.Println(isPalindrome("Never odd or even"))
}
