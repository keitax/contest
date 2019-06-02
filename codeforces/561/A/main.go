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
	sts := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%s\n", &sts[i])
	}

	cnt := map[byte]int{}
	for i := 0; i < n; i++ {
		if _, ok := cnt[sts[i][0]]; !ok {
			cnt[sts[i][0]] = 0
		}
		cnt[sts[i][0]]++
	}

	var ans int
	for _, v := range cnt {
		ans += Combination(v/2, 2) + Combination((v-1)/2+1, 2)
	}

	fmt.Println(ans)
}

func Combination(n, r int) int {
	if n < r {
		return 0
	}
	if r == 0 || n == r {
		return 1
	}
	if r == 1 {
		return n
	}
	return Combination(n-1, r-1) + Combination(n-1, r)
}
