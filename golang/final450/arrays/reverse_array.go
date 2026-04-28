package arrays

import (
	"fmt"
	"os"
)

func ReverseArray(input []int) []int {
	if len(input) <= 0 {
		fmt.Println("input is too small.")
		os.Exit(1)
	}

	if len(input) == 1 {
		return input
	}

	for i := 0; i < len(input)/2; i++ {
		input[i], input[len(input)-i-1] = input[len(input)-i-1], input[i]
	}

	return input
}
