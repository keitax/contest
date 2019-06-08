// http://codeforces.com/contest/1166/problem/C

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	fmt.Fscanf(in, "%d\n", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d", &a[i])
	}
	fmt.Fprintln(out, Solve(n, a))
}

func Solve(n int, a []int) int64 {
	b := make([]int, n)
	for i := int(0); i < n; i++ {
		if a[i] < 0 {
			b[i] = -a[i]
		} else {
			b[i] = a[i]
		}
	}
	sort.Ints(b)

	cnt := int64(0)
	for l := int(0); l < n; l++ {
		r := BinSearch(b, l, b[l]*2)
		cnt += int64(r - l)
	}

	return cnt
}

func BinSearch(a []int, l, n int) int {
	l = l - 1
	r := len(a)
	for r-l > 1 {
		m := l + (r-l)/2
		if a[m] <= n {
			l = m
		} else {
			r = m
		}
	}
	return l
}
