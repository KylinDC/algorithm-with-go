package sort

import (
	"algorithm-with-go/utils"
	"cmp"
)

func SelectionSorting[S ~[]E, E cmp.Ordered](x S) (newX S) {
	newX = utils.CopySlice(x)

	for i := range newX {
		minI := i
		for j := i; j < len(x); j++ {
			if cmp.Less(newX[j], newX[minI]) {
				minI = j
			}
		}
		Swap(newX, i, minI)
	}

	return
}
