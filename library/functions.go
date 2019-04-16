package library

import "sort"

func NextPerm(xs []int) bool {
	l := len(xs) - 2
	for 0 <= l && xs[l] >= xs[l+1] {
		l--
	}
	if l < 0 {
		return false
	}
	r := l + 1
	for i := l + 1; i < len(xs); i++ {
		if xs[i] > xs[l] && xs[i] < xs[r] {
			r = i
		}
	}
	xs[l], xs[r] = xs[r], xs[l]
	sort.Ints(xs[l+1:])
	return true
}
