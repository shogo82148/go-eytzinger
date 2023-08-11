package eytzinger

import (
	"cmp"
	"math/bits"
)

// Eytzinger rearranges the elements of a slice into a cache friendly binary tree
// as known as [Eytzinger Layout].
// The slice must be sorted in increasing order.
//
// [Eytzinger Layout]: https://en.wikipedia.org/wiki/Binary_tree#Arrays
func Eytzinger[S ~[]E, E any](a S) S {
	b := make(S, len(a))
	ctx := &eytzingerContext[S, E]{a: a, b: b}
	ctx.eytzinger(0, 0)
	return b
}

type eytzingerContext[S ~[]E, E any] struct {
	a S // input slice
	b S // output slice
}

// based on https://algorithmica.org/en/eytzinger
func (ctx *eytzingerContext[S, E]) eytzinger(i, k int) int {
	if k < len(ctx.a) {
		i = ctx.eytzinger(i, 2*k+1)
		ctx.b[k] = ctx.a[i]
		i++
		i = ctx.eytzinger(i, 2*k+2)
	}
	return i
}

// IsEytzinger reports whether a slice is in Eytzinger Layout.
func IsEytzinger[S ~[]E, E cmp.Ordered](a S) bool {
	return isEytzinger(a, 0)
}

func isEytzinger[S ~[]E, E cmp.Ordered](a S, k int) bool {
	if k >= len(a) {
		return true
	}
	i, j := 2*k+1, 2*k+2
	if i < len(a) && cmp.Less(a[k], a[i]) {
		return false
	}
	if j < len(a) && cmp.Less(a[j], a[k]) {
		return false
	}
	return isEytzinger(a, i) && isEytzinger(a, j)
}

// Search searches for target in a Eytzinger Layout slice.
func Search[S ~[]E, E cmp.Ordered](x S, target E) (int, bool) {
	k := 1
	for k <= len(x) {
		if cmp.Less(x[k-1], target) {
			k = 2*k + 1
		} else {
			k = 2 * k
		}
	}
	k >>= bits.TrailingZeros(^uint(k)) + 1
	if k == 0 {
		return len(x), false
	}
	k--
	return k, k < len(x) && (x[k] == target || isNaN(target) && isNaN(x[k]))
}

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T cmp.Ordered](x T) bool {
	return x != x
}
