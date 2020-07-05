package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the aVeryBigSum function below.
func aVeryBigSum(ar []int64) int64 {
	var hasil int64
	for _, v := range ar {
		hasil += v
	}
	return hasil
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int64

	for i := 0; i < int(arCount); i++ {
		arItem, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		ar = append(ar, arItem)
	}

	result := aVeryBigSum(ar)

	fmt.Printf("%d\n", result)
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
