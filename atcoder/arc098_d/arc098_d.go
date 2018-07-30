// https://arc098.contest.atcoder.jp/tasks/arc098_b
package main

import (
	"bufio"
	"os"
	"fmt"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

// しゃくとり法
// https://paiza.hatenablog.com/entry/2015/01/21/【累積和、しゃくとり法】初級者でも解るアルゴ
func main() {
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &a[i])
	}
	fmt.Fscanf(reader, "\n")

	ba := []string{}
	for i := 0; i < len(a); i++ {
		ba = append(ba, fmt.Sprintf("%b", a[i]))
	}

	cnt := 0
	base := make([]int, 20)
	for _, bai := range ba {
		os := len(base) - len(bai)
		for j, baij := range bai {
			base[j+os] += int(baij - '0')
		}
		flg := false
		for _, b := range base {
			if b > 1 {
				flg = true
				break
			}
		}
		if flg {
			// cancel
			for j, baij := range bai {
				base[j+os] -= int(baij - '0')
			}
		} else {
			cnt++
		}
	}

	//base := "00000000000000000000"
	fmt.Fprintf(writer, "%#v\n", ba)
}
