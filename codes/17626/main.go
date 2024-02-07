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

	var n int
	fmt.Fscanln(r, &n)

	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		min := i

		for j := 1; j*j <= i; j++ {
			if min > dp[i-j*j] {
				min = dp[i-j*j]
			}
		}
		dp[i] = min + 1
	}

	//fmt.Fprintln(w, dp)
	fmt.Fprintln(w, dp[n])
}
