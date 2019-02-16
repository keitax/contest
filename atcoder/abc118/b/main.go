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

	var n, m int
	fmt.Fscanf(in, "%d %d\n", &n, &m)
	c := make([][]bool, n)
	for i := 0; i < n; i++ {
		c[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		var ki int
		fmt.Fscanf(in, "%d", &ki)
		for j := 0; j < ki; j++ {
			var a int
			fmt.Fscanf(in, "%d", &a)
			c[i][a-1] = true
		}
		fmt.Fscanf(in, "\n")
	}

	v := make([]bool, m)
	for j := 0; j < m; j++ {
		v[j] = c[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			v[j] = v[j] && c[i][j]
		}
	}

	cnt := 0
	for j := 0; j < m; j++ {
		if v[j] {
			cnt++
		}
	}
	fmt.Fprintln(out, cnt)
}
