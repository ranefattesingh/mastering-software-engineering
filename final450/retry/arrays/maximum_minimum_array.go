package arrays

func MaximumAndMimimumElement(arr []int) (maximum, minimum int) {
	index := 0

	// check if array contains odd number of elements.
	if len(arr)%2 == 1 {
		minimum = arr[0]
		maximum = arr[0]
		index = 1
	} else {
		if arr[0] < arr[1] {
			maximum = arr[1]
			minimum = arr[0]
		} else {
			maximum = arr[0]
			minimum = arr[1]
		}
		index = 2
	}

	for ; index < len(arr); index += 2 {
		if arr[index] < arr[index+1] {
			minimum = min(maximum, arr[index])
			maximum = max(maximum, arr[index+1])
		} else {
			minimum = min(maximum, arr[index+1])
			maximum = max(maximum, arr[index])
		}
	}

	return maximum, minimum
}
