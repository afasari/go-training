package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	// start
	var min, max, sum int64
	max = int64(arr[0])
	min = int64(arr[0])
	for _, value := range arr {
		if int64(value) > max {
			max = int64(value)
		}
		if int64(value) < min {
			min = int64(value)
		}
		sum += int64(value)
	}

	// fmt.Printf("max = %d\n", max)
	// fmt.Printf("min = %d\n", min)
	fmt.Printf("%d %d", sum-max, sum-min)
	// fmt.Printf("%d %d", min, max)
	// end
}


func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
