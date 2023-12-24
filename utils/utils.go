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

func CopySlice[S ~[]E, E cmp.Ordered](x S) (newX S) {
	newX = make(S, len(x))
	copy(newX, x)
	return
}
