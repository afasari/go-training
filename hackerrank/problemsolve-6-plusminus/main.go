package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func calculateFraction(val float32, len float32) (ret float32){
	return float32(val/len)
}

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	var positive, negative, zero float32 = 0,0,0
	for _, value := range arr {
		if value > 0{
			positive += 1
		} else if value < 0 {
			negative += 1
		} else {
			zero += 1
		}
	}
	len := float32(len(arr))

	fmt.Printf("%.5f\n",calculateFraction(positive, len))
	fmt.Printf("%.5f\n",calculateFraction(negative, len))
	fmt.Printf("%.5f\n",calculateFraction(zero, len))


}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
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
