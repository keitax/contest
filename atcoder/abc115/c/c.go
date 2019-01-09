package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscanf(in, "%d %d\n", &n, &k)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d\n", &h[i])
	}
	sort.Ints(h)

	min := math.MaxInt32
	for i := 0; i <= n-k; i++ {
		diff := h[i+k-1] - h[i]
		if min > diff {
			min = diff
		}
	}

	fmt.Fprintln(out, min)
}
