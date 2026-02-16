package array

func BubbleSort(vector []int) []int {
	size := len(vector)

	if size <= 1 {
		return vector
	}

	end := size - 1

	for end > 0 {
		for index := 0; index < end; index++ {
			if vector[index] > vector[index+1] {
				temp := vector[index+1]
				vector[index+1] = vector[index]
				vector[index] = temp
			}
		}
		end--
	}

	return vector
}
