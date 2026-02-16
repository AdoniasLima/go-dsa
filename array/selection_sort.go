package array

func SelectionSort(vector []int) []int {
	size := len(vector)

	if size <= 1 {
		return vector
	}

	end := size - 1

	for i := 0; i < end; i++ {
		smallerIndex := i

		for j := i + 1; j <= end; j++ {
			if vector[j] < vector[smallerIndex] {
				smallerIndex = j
			}
		}

		if smallerIndex > i {
			temp := vector[smallerIndex]
			vector[smallerIndex] = vector[i]
			vector[i] = temp
		}
	}

	return vector
}
