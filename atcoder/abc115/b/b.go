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
	p := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d\n", &p[i])
	}

	sum := 0
	max := 0
	for i := 0; i < n; i++ {
		if max < p[i] {
			max = p[i]
		}
		sum += p[i]
	}

	ans := sum - max/2

	fmt.Fprintln(out, ans)
}
