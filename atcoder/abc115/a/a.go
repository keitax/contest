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

	var d int
	_, err := fmt.Fscanf(in, "%d\n", &d)
	if err != nil {
		panic(err)
	}

	switch d {
	case 25:
		fmt.Fprintln(out, "Christmas")
	case 24:
		fmt.Fprintln(out, "Christmas Eve")
	case 23:
		fmt.Fprintln(out, "Christmas Eve Eve")
	case 22:
		fmt.Fprintln(out, "Christmas Eve Eve Eve")
	}
}
