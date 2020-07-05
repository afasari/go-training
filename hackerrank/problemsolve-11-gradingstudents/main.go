package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the gradingStudents function below.
 */
func gradingStudents(grades []int32) []int32 {
	/*
	 * Write your code here.
	 */
	var arrayNum []int32
	for _, original := range grades {
		if original < 38 {
			original = original
		} else if original%5 >= 3 {
			original = original/5*5 + 5
		}
		arrayNum = append(arrayNum, original)
	}
	return arrayNum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grades []int32

	for gradesItr := 0; gradesItr < int(n); gradesItr++ {
		gradesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		gradesItem := int32(gradesItemTemp)
		grades = append(grades, gradesItem)
	}

	result := gradingStudents(grades)

	for resultItr, resultItem := range result {
		fmt.Printf("%d", resultItem)

		if resultItr != len(result)-1 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
