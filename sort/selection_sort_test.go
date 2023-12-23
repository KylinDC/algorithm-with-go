package sort

import (
	"algorithm-with-go/utils"
	"math/rand"
	"reflect"
	"testing"
)

func generateBestCase(count int) []int {
	arr := make([]int, count)
	for i := 0; i < 10; i++ {
		arr[i] = i + 1
	}
	return arr
}

func generateWorstCase(count int) []int {
	arr := make([]int, count)
	for i := 0; i < 10; i++ {
		arr[i] = 10 - i
	}
	return arr
}

func generateAverageCase(count int) []int {
	arr := make([]int, count)
	for i := 0; i < count; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

type TestData struct {
	input       []int
	expected    []int
	description string
}

func generateTests() []TestData {
	bestCase := generateBestCase(1000)
	worstCase := generateWorstCase(1000)
	averageCase := generateAverageCase(1000)

	tests := []TestData{
		{input: bestCase, expected: utils.ToSorted(bestCase), description: "best case"},
		{input: worstCase, expected: utils.ToSorted(worstCase), description: "worst case"},
		{input: averageCase, expected: utils.ToSorted(averageCase), description: "average case"},
	}
	return tests
}

func BenchmarkSelectionSort(b *testing.B) {
	tests := generateTests()

	for _, test := range tests {
		b.Run(test.description, func(b *testing.B) {
			actual := SelectionSorting(test.input)

			if !reflect.DeepEqual(actual, test.expected) {
				b.Errorf("order don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}
