package main

import (
	"fmt"
	"sort"
)

func main() {

	var numInputs int
	fmt.Scan(&numInputs)
	slice := make([]int, numInputs, 2*numInputs)
	for i := 0; i < numInputs; i++ {
		fmt.Scan(&slice[i])
	}

	// Calculate the mean
	sum := 0
	for _, num := range slice {
		sum += num
	}
	fmt.Printf("%.1f\n", float64(sum)/float64(numInputs))

	// Calculating the median
	sort.Ints(slice)
	if numInputs%2 != 0 {
		fmt.Println(slice[numInputs/2])
	} else {
		fmt.Println(0.5 * float32(slice[numInputs/2]+slice[numInputs/2-1]))
	}

	// Calculating the mode
	mostCommon := slice[numInputs-1]
	currentNum := slice[numInputs-1]
	currentFrequency := 1
	maxFrequency := 1
	for i := numInputs - 2; i >= 0; i-- {
		if slice[i] == currentNum {
			currentFrequency++
		} else {
			currentFrequency = 1
			currentNum = slice[i]
		}
		if currentFrequency >= maxFrequency {
			maxFrequency = currentFrequency
			mostCommon = slice[i]
		}
	}
	fmt.Println(mostCommon)
}
