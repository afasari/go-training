package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func compareTriplets(a []int32, b []int32) []int32 {
	var solutiona int32
	var solutionb int32
	for i, _ := range a {
		if a[i] > b[i] {
			solutiona += 1
		}
		if a[i] < b[i] {
			solutionb += 1
		}
	}
	solution := []int32{solutiona, solutionb}
	return solution
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < 3; i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	bTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var b []int32

	for i := 0; i < 3; i++ {
		bItemTemp, err := strconv.ParseInt(bTemp[i], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	result := compareTriplets(a, b)

	for i, resultItem := range result {
		fmt.Printf("%d", resultItem)

		if i != len(result)-1 {
			fmt.Printf(" ")
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
