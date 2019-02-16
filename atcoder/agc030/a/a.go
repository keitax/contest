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

	var a, b, c int
	fmt.Fscanf(in, "%d %d %d\n", &a, &b, &c)

	ans := b + min(c, a+b+1)
	fmt.Fprintln(out, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
