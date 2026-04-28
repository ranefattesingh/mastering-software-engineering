package arrays

import (
	"fmt"
	"os"
)

// function takes in array input and returns array output
// where index 0 is minumum and index 1 is maximum.
func FindMinMax(input []int) []int {
	if len(input) <= 0 {
		fmt.Println("input is too small.")
		os.Exit(1)
	}

	if len(input) == 1 {
		return []int{input[0], input[0]}
	}

	min := input[0]
	max := input[0]

	for i := 1; i < len(input); i++ {
		if min > input[i] {
			min = input[i]
		} else if max < input[i] {
			max = input[i]
		}
	}

	return []int{min, max}
}
