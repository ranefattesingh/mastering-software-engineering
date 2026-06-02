package main

import "fmt"

func MaxSubArraySum(arr []int) int {
	maximum := arr[0]
	currentMaximum := arr[0]
	for i := 1; i < len(arr); i++ {
		currentMaximum = max(currentMaximum+arr[i], arr[i])
		maximum = max(maximum, currentMaximum)
	}

	return maximum
}

func main() {
	sum := MaxSubArraySum([]int{2, -3, 8, 7, -1, 2, 3})
	fmt.Println(sum)
}
