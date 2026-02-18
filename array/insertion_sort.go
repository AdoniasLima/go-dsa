package array

func InsertionSort(vector []int) []int {
	size := len(vector)

	if size <= 1 {
		return vector
	}

	end := size - 1

	for i := 1; i <= end; i++ {
		marked := vector[i]

		j := i - 1

		for j >= 0 && marked < vector[j] {
			vector[j+1] = vector[j]
			j--
		}

		vector[j+1] = marked
	}

	return vector
}
