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

	dp := make([]int, 101)
	dp[1] = 1
	dp[2] = 1
	dp[3] = 1
	dp[4] = 2
	dp[5] = 2
	for i := 6; i < 101; i++ {
		dp[i] = dp[i-1] + dp[i-5]
	}

	for i := 0; i < T; i++ {
		var N int
		fmt.Fscanln(r, &N)
		fmt.Fprintln(w, dp[N])
	}
}
