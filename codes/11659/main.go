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

	var N, M int
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	sum := make([]int, N+1)

	for i := 1; i < N+1; i++ {
		var x int
		fmt.Fscan(r, &x)
		sum[i] = sum[i-1] + x
	}
	fmt.Fscanln(r)

	for k := 0; k < M; k++ {
		var i, j int
		fmt.Fscan(r, &i)
		fmt.Fscan(r, &j)
		fmt.Fscanln(r)
		result := sum[j] - sum[i-1]
		fmt.Fprintln(w, result)
	}

}
