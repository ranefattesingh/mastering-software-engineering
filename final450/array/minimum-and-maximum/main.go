// Given an array of integers arr[], the task is to find the maximum and minimum elements in the array using the minimum number of comparisons.

// Examples:

// Input: arr[] = [3, 5, 4, 1, 9]
// Output: [1, 9]
// Explanation: The minimum element is 1, and the maximum element is 9.

// Input: arr[] = [22, 14, 8, 17, 35, 3]
// Output: [3, 35]
// Explanation: The minimum element is 3, and the maximum element is 35.
package main

import "errors"

var ErrInvalidArraySize = errors.New("invalid array size")

func FindMinimumAndMaximum(arr []int, size int) (minNum, maxNum int, err error) {
	if size <= 0 {
		return -1, -1, ErrInvalidArraySize
	}

	i := 0

	if size%2 == 1 {
		maxNum = arr[0]
		minNum = arr[0]
		i = 1
	} else {
		if arr[0] < arr[1] {
			minNum = arr[0]
			maxNum = arr[1]
		} else {
			minNum = arr[1]
			maxNum = arr[0]
		}

		i = 2
	}

	for ; i < size-1; i += 2 {
		if arr[i] < arr[i+1] {
			minNum = min(minNum, arr[i])
			maxNum = max(maxNum, arr[i+1])
		} else {
			minNum = min(minNum, arr[i+1])
			maxNum = max(maxNum, arr[i])
		}
	}

	return minNum, maxNum, nil
}
