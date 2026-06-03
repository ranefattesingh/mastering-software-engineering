package main

import (
	"container/heap"
	"errors"
)

var (
	ErrArrayEmpty = errors.New("array is empty")
	ErrKValue     = errors.New("k should be greater than 0 and less than length of array")
)

// Given an integer array arr[] and an integer k, find and return the kth smallest element in the given array.
// Note: The kth smallest element is determined based on the sorted order of the array.

// Examples :

// Input: arr[] = [10, 5, 4, 3, 48, 6, 2, 33, 53, 10], k = 4
// Output: 5
// Explanation: 4th smallest element in the given array is 5.
// Input: arr[] = [7, 10, 4, 3, 20, 15], k = 3
// Output: 7
// Explanation: 3rd smallest element in the given array is 7.

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(e any) {
	*pq = append(*pq, e.(int))
}

func (pq *PriorityQueue) Pop() any {
	prev := *pq
	e := prev[len(prev)-1]
	*pq = prev[0 : len(prev)-1]

	return e
}

func FindKthSmallestElement(arr []int, k int) (int, error) {
	if len(arr) <= 0 {
		return -1, ErrArrayEmpty
	}

	if k > 0 && k > len(arr) {
		return -1, ErrKValue
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i := 0; i < len(arr); i++ {
		heap.Push(&pq, arr[i])
		if pq.Len() > k {
			heap.Pop(&pq)
		}
	}

	return pq[k], nil
}

func main() {
	FindKthSmallestElement([]int{10, 5, 4, 3, 48, 6, 2, 33, 53, 10}, 4)
}
