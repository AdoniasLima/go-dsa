package array

func MergeSort(vector []int) []int {
	size := len(vector)

	if size <= 1 {
		return vector
	}

	half := size / 2
	left := make([]int, half)
	right := make([]int, size-half)

	copy(left, vector[:half])
	copy(right, vector[half:])

	MergeSort(left)
	MergeSort(right)

	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			vector[k] = left[i]
			i++
		} else {
			vector[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		vector[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		vector[k] = right[j]
		j++
		k++
	}

	return vector
}
