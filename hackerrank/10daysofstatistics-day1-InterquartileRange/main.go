// Day 1 Challenge: Calculating the Interquartile range
package main

import (
	"fmt"
	"sort"
)

func main() {
	var numVals int
	fmt.Scan(&numVals)
	values := make([]int, numVals)
	for i := 0; i < numVals; i++ {
		fmt.Scan(&values[i])
	}

	frequencies := make([]int, numVals)
	for i := 0; i < numVals; i++ {
		fmt.Scan(&frequencies[i])
	}

	totalNumbers := 0
	for _, num := range frequencies {
		totalNumbers += num
	}

	var arr []int
	/*
		values[i] needs to be replicated frequencies[i] times so if values[i] = 3 and frequencies[i] = 5
		the final array would contain [3 3 3 3 3]
	*/

	for i := 0; i < len(values); i++ {
		for j := 0; j < frequencies[i]; j++ {
			arr = append(arr, values[i])
		}
	}
	sort.Ints(arr)
	firstHalf := make([]int, totalNumbers/2)
	secondHalf := make([]int, totalNumbers/2)

	medianIndex := totalNumbers / 2

	if totalNumbers%2 != 0 {
		for i := 0; i < medianIndex; i++ {
			firstHalf[i] = arr[i]
			secondHalf[i] = arr[totalNumbers-i-1]
		}

	} else {
		for i := 0; i < medianIndex; i++ {
			firstHalf[i] = arr[i]
			secondHalf[i] = arr[totalNumbers-i-1]
		}
	}

	quartileIndex := medianIndex / 2
	var q1Value float64
	var q3Value float64
	if medianIndex%2 != 0 {
		q1Value = float64(firstHalf[quartileIndex])
		q3Value = float64(secondHalf[quartileIndex])
	} else {
		q1Value = float64(firstHalf[quartileIndex]+firstHalf[quartileIndex-1]) / float64(2)
		q3Value = float64(secondHalf[quartileIndex]+secondHalf[quartileIndex-1]) / float64(2)
	}
	fmt.Printf("%.1f\n", q3Value-q1Value)
}
