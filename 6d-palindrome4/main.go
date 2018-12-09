package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func Reverse(s string) string {
	totalLength := len(s)
	buffer := make([]byte, totalLength)
	for i := 0; i < totalLength; {
		r, size := utf8.DecodeRuneInString(s[i:])
		i += size
		utf8.EncodeRune(buffer[totalLength-i:], r)
	}
	return string(buffer)
}

func isPalindrome(input string, caseSensitive bool) bool {

	if caseSensitive != true {
		input = strings.ToLower(input)
	}

	// to deal with phrases, we need to eliminate the
	// spaces

	input = strings.TrimSpace(input)

	reverse := Reverse(input)

	if input == reverse {
		return true
	} else {
		return false
	}

}

func main() {

	fmt.Println("Enter a word, phrase or number : ")

	consoleReader := bufio.NewReader(os.Stdin)
	answer, _ := consoleReader.ReadString('\n')

	// get rid of the extra newline character from ReadString() function
	answer = strings.TrimSuffix(answer, "\n")

	fmt.Println(answer, "is palindrome[case sensitive] ?", isPalindrome(answer, true))
	fmt.Println(answer, "is palindrome[case INsensitive] ?", isPalindrome(answer, false))

}