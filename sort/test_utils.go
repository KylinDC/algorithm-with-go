package sort

import (
	"algorithm-with-go/utils"
	"math/rand"
)

func generateBestCase(count int) []int {
	arr := make([]int, count)
	for i := 0; i < count; i++ {
		arr[i] = i + 1
	}
	return arr
}

func generateWorstCase(count int) []int {
	arr := make([]int, count)
	for i := 0; i < count; i++ {
		arr[i] = count - i
	}
	return arr
}

func generateAverageCase(count int) []int {
	arr := make([]int, count)
	for i := 0; i < count; i++ {
		arr[i] = rand.Intn(count)
	}
	return arr
}

type TestData struct {
	input       []int
	expected    []int
	description string
}

func GenerateTests(count int) []TestData {
	bestCase := generateBestCase(count)
	worstCase := generateWorstCase(count)
	averageCase := generateAverageCase(count)

	tests := []TestData{
		{input: bestCase, expected: utils.ToSorted(bestCase), description: "best case"},
		{input: worstCase, expected: utils.ToSorted(worstCase), description: "worst case"},
		{input: averageCase, expected: utils.ToSorted(averageCase), description: "average case"},
	}
	return tests
}
