package arrays

// 7 , 10, 4, 3, 20, 15
func FindKthSmallestElement(input []int, k int) int {
	window := make([]int, k)
	window[0] = input[0]
	wi, wj := 0, 0

	for i := 0; i < len(input); i++ {
		j := len(window) - 1

		for window[j] > input[i] && j > 0 {
			window[j] = window[j-1]
			j--
		}

		window[j] = input[i]
	}

	return window[k-1]
}
