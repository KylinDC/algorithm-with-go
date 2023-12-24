package sort

import "cmp"

func Swap[S ~[]E, E cmp.Ordered](x S, i1 int, i2 int) {
	x[i1], x[i2] = x[i2], x[i1]
}
