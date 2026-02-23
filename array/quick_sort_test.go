package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type QuickSortInputTests struct {
	input    []int
	expected []int
}

func TestQuickSort(t *testing.T) {
	assert := assert.New(t)

	tests := []QuickSortInputTests{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{0, 1}, []int{0, 1}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{5, 1}, []int{1, 5}},
		{[]int{100, 40, 9, 20, 11}, []int{9, 11, 20, 40, 100}},
	}

	for _, currentTest := range tests {
		result := QuickSort(currentTest.input)
		assert.Equal(currentTest.expected, result, "Failed: expected %d but got %d", currentTest.expected, result)
	}
}
