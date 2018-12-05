package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
	"bytes"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(time string) string {
	/*
	 * Write your code here.
	 */
	trimmedTime := time[:8] // HH:MM:SS

	splittedTime := strings.Split(trimmedTime, ":")
	standardHour, _ := strconv.Atoi(splittedTime[0])

	meridiem := string(time[8:]) //AM:PM
	var militaryTime bytes.Buffer
	switch meridiem {
	case "AM":
		if standardHour == 12 {
			militaryTime.WriteString("00" + ":" + splittedTime[1] + ":" + splittedTime[2])
			return militaryTime.String()
		} else {
			// Nothing changed
			return trimmedTime
		}
	case "PM":
		if standardHour == 12 {
			return trimmedTime
		} else {
			militaryHour := standardHour + 12
			militaryTime.WriteString(strconv.Itoa(militaryHour) + ":" + splittedTime[1] + ":" + splittedTime[2])
			return militaryTime.String()
		}
	default:
		return ""
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	s := readLine(reader)

	result := timeConversion(s)

	fmt.Printf( "%s\n", result)
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
