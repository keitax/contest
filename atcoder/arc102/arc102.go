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

	var ans int64
	if k % 2 == 0 {
		ans = int64(math.Pow(float64(n / k + 1), 3))
	} else {
		ans =
	}
}
