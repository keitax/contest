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

	var n, x int64
	fmt.Fscanf(in, "%d %d\n", &n, &x)

	fmt.Fprintln(out, f(n, x))
}

func f(n, x int64) int64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if x <= 1+a(n-1) {
		return f(n-1, x-1)
	}
	return 1 + p(n-1) + f(n-1, x-2-a(n-1))
}

func a(n int64) int64 {
	if n == 0 {
		return 1
	}
	return 2*a(n-1) + 3
}

func p(n int64) int64 {
	if n == 0 {
		return 1
	}
	return 2*p(n-1) + 1
}
