package sort

import (
	"algorithm-with-go/utils"
	"cmp"
)

func InsertionSort[S ~[]E, E cmp.Ordered](x S) (newX S) {
	newX = utils.CopySlice(x)

	for i := range newX {
		for j := i; j > 0; j-- {
			if newX[j] < newX[j-1] {
				utils.Swap(newX, j, j-1)
			} else {
				break
			}
		}
	}

	return
}
