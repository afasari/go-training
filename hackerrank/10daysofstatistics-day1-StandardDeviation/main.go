// Calculating the standard deviation
// GIVEN: a list of numbers
// CALCULATE: Standard Deviation of the given numbers

package main

import (
	"fmt"
	"math"
)

func main() {
	var numVals int
	fmt.Scan(&numVals)
	arr := make([]float64, numVals)

	for i := 0; i < numVals; i++ {
		fmt.Scan(&arr[i])
	}

	// Calculating the mean

	sum := float64(0)

	for _, num := range arr {
		sum += num
	}

	mean := float64(sum) / float64(numVals)

	for i := 0; i < numVals; i++ {
		arr[i] = (arr[i] - mean) * (arr[i] - mean)
	}
	sumOfDifferences := float64(0)

	for _, num := range arr {
		sumOfDifferences += num
	}

	fmt.Printf("%0.1f", math.Sqrt(sumOfDifferences/float64(numVals)))
}