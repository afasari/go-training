package main

import (
	"fmt"
	"sort"
)

// This is the first day 1 challenge
// first line is number of values
// second line are the actual numbers

func main() {
	var numValues int
	fmt.Scan(&numValues)

	arr := make([]int, numValues)
	for i := 0; i < numValues; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr)

	// Calculating Median
	// _ _ _ _ _ _ _ _ _ _
	firstHalf := make([]int, numValues/2)
	secondHalf := make([]int, numValues/2)
	medianIndex := numValues / 2
	var medianValue int

	if numValues%2 != 0 {
		medianValue = arr[medianIndex]
		for i := 0; i < medianIndex; i++ {
			firstHalf[i] = arr[i]
			secondHalf[i] = arr[numValues-i-1]
		}

	} else {
		medianValue = (arr[medianIndex] + arr[medianIndex-1]) / 2
		for i := 0; i < medianIndex; i++ {
			firstHalf[i] = arr[i]
			secondHalf[i] = arr[numValues-i-1]
		}
	}

	// Calculating Q1 and Q3

	quartileIndex := medianIndex / 2
	var q1Value int
	var q3Value int
	if medianIndex%2 != 0 {
		q1Value = firstHalf[quartileIndex]
		q3Value = secondHalf[quartileIndex]
	} else {
		q1Value = (firstHalf[quartileIndex] + firstHalf[quartileIndex-1]) / 2
		q3Value = (secondHalf[quartileIndex] + secondHalf[quartileIndex-1]) / 2
	}

	//STDOUT

	fmt.Println(q1Value)
	fmt.Println(medianValue)
	fmt.Println(q3Value)
}

/*
   9
   3 5 7 8 12 13 14 18 21 34 51
   _  _  (_)  _  (_)  _  _  _  _
   9/2 = 4
*/
