package main

import (
	"fmt"
)

func main() {
	var numInputs int
	fmt.Scan(&numInputs)
	nums := make([]int, numInputs)
	weights := make([]int, numInputs)

	for i := 0; i < numInputs; i++ {
		fmt.Scan(&nums[i])
	}

	for i := 0; i < numInputs; i++ {
		fmt.Scan(&weights[i])
	}

	numerator := 0

	for i := 0; i < numInputs; i++ {
		numerator += nums[i] * weights[i]
	}

	denominator := 0

	for _, num := range weights {
		denominator += num
	}
	fmt.Printf("%.1f\n", float64(numerator)/float64(denominator))
}