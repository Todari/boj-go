package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N int

	fmt.Fscanln(r, &N)

	// var s sort.IntSlice
	// var s []int
	counts := make([]int, 10001)

	for i := 0; i < N; i++ {
		var x int
		fmt.Fscanln(r, &x)
		counts[x] += 1
		// s = append(s, x)
	}
	// fmt.Fprintln(w, counts)

	for i := 0; i < 10001; i++ {
		for j := 0; j < counts[i]; j++ {
			fmt.Fprintln(w, i)
		}
	}
}
