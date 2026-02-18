package array

func ShellSort(vector []int) []int {
	size := len(vector)

	if size <= 1 {
		return vector
	}

	interval := size / 2
	end := size - 1

	for interval > 0 {
		for i := interval; i <= end; i++ {
			temp := vector[i]
			j := i

			for j >= interval && vector[j-interval] > temp {
				vector[j] = vector[j-interval]
				j -= interval
			}

			vector[j] = temp
		}

		interval = interval / 2
	}

	return vector
}
