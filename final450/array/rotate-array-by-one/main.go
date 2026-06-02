package main

import "fmt"

func RotateArrayByOne(arr []int) {
	if len(arr) == 0 {
		fmt.Println("Err: input array is nil or empty")
		return
	}

	lastElement := arr[len(arr)-1]
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}

	arr[0] = lastElement
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	RotateArrayByOne(arr)
	fmt.Println(arr)
}
