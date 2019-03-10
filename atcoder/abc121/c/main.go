package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	var items [][2]int
	var n, m int
	fmt.Fscanf(stdin, "%d %d\n", &n, &m)
	for i := 0; i < n; i++ {
		var item [2]int
		fmt.Fscanf(stdin, "%d %d\n", &item[0], &item[1])
		items = append(items, item)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	ans := 0
	k := 0
	for c := m; c > 0; c-- {
		ans += items[k][0]
		items[k][1]--
		if items[k][1] < 1 {
			k++
		}
	}

	fmt.Fprintf(stdout, "%d\n", ans)
}
