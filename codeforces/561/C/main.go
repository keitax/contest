package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscanf(in, "%d\n", &n)
	vs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d", &vs[i])
	}
	fmt.Fscanln(in)

	cnt := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if i == j {
				continue
			}
			if isLegend(vs[i], vs[j]) {
				cnt++
			}
		}
	}
	fmt.Fprintln(out, cnt)
}

func isLegend(x, y int) bool {
	ax, ay := abs(x), abs(y)
	axmy, axpy := abs(x-y), abs(x+y)
	return min(axmy, axpy) <= min(ax, ay) && max(ax, ay) <= max(axmy, axpy)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
