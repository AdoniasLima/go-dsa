package array

func QuickSort(vector []int) []int {
	size := len(vector)

	if size <= 1 {
		return vector
	}

	pivot := vector[size/2]

	var left, equal, right []int

	for i := range size {
		if vector[i] == pivot {
			equal = append(equal, vector[i])
		} else if vector[i] < pivot {
			left = append(left, vector[i])
		} else {
			right = append(right, vector[i])
		}
	}

	left = append(QuickSort(left), equal...)
	right = QuickSort(right)

	return append(left, right...)
}
