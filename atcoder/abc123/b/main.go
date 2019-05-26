package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()

	cook := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscanf(in, "%d\n", &cook[i])
	}
	sort.Ints(cook)

	fmt.Fprintf(out, "%v\n", cook)
}

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
