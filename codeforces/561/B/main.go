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

	var k int
	fmt.Fscanf(in, "%d\n", &k)

	var n, m int
	for n = 5; float64(n) <= math.Sqrt(float64(k)); n++ {
		if k%n == 0 {
			m = k / n
			break
		}
	}
	if m == 0 {
		fmt.Fprintln(out, -1)
		return
	}
	for i := 0; i < n; i++ {
		fmt.Fprint(out, vowel(m, i))
	}
	fmt.Fprintln(out)
}

func vowel(len int, offset int) string {
	w := "aiueo"
	r := make([]byte, len)
	for i := 0; i < len; i++ {
		r[i] = w[(i+offset)%5]
	}
	return string(r)
}
