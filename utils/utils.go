package utils

import (
	"cmp"
	"slices"
)

func ToSorted[S ~[]E, E cmp.Ordered](x S) (newX S) {
	newX = make([]E, len(x))
	copy(newX, x)

	slices.Sort(newX)
	return
}

func DeepCopy[S ~[]E, E cmp.Ordered](x S) (newX S) {
	newX = make(S, len(x))
	copy(newX, x)
	return
}

func Swap[S ~[]E, E cmp.Ordered](x S, i1 int, i2 int) {
	x[i1], x[i2] = x[i2], x[i1]
}