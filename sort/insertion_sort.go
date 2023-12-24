package sort

import (
	"algorithm-with-go/utils"
	"cmp"
)

func InsertionSort[S ~[]E, E cmp.Ordered](x S) (newX S) {
	newX = utils.CopySlice(x)

	for i := range newX {
		for j := i; j > 0; j-- {
			if cmp.Less(newX[j], newX[j-1]) {
				Swap(newX, j, j-1)
			} else {
				break
			}
		}
	}

	return
}

func ShellSort[S ~[]E, E cmp.Ordered](x S) (newX S) {
	newX = utils.CopySlice(x)

	h := 1
	for h < len(newX)/3 {
		h = h*3 + 1
	}

	for h >= 1 {
		for i := h; i < len(newX); i++ {
			for j := i; j >= h; j -= h {
				if cmp.Less(newX[j], newX[j-h]) {
					Swap(newX, j, j-h)
				} else {
					break
				}
			}
		}

		h = h / 3
	}

	return
}
