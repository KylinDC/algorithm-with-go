package sort

import (
	"reflect"
	"testing"
)

func BenchmarkSelectionSort(b *testing.B) {
	tests := GenerateTests(1000)

	for _, test := range tests {
		b.Run(test.description, func(b *testing.B) {
			actual := SelectionSorting(test.input)

			if !reflect.DeepEqual(actual, test.expected) {
				b.Errorf("order don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}
