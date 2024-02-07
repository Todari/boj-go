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

	var T int
	fmt.Fscanln(r, &T)
	dp := make([]int, 11)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 4
	for i := 3; i < 11; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	for i := 0; i < T; i++ {
		var n int
		fmt.Fscanln(r, &n)
		fmt.Fprintln(w, dp[n-1])
	}
}
