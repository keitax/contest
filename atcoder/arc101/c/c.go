package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscanf(in, "%d %d\n", &n, &k)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d", &x[i])
	}
	fmt.Fscanln(in)

	minS := math.MaxInt32
	for t := 0; t <= n-k; t++ {
		s := 0
		if x[t+k-1] <= 0 {
			s = abs(x[t])
		} else if x[t] >= 0 {
			s = abs(x[t+k-1])
		} else if abs(x[t]) > abs(x[t+k-1]) {
			s = abs(x[t]) + 2*abs(x[t+k-1])
		} else {
			s = 2*abs(x[t]) + abs(x[t+k-1])
		}
		if s < minS {
			minS = s
		}
	}

	fmt.Fprintln(out, minS)
}

func abs(x int) int {
	return int(math.Sqrt(float64(x * x)))
}
