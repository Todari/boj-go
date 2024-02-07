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

	dp := make([]int, 1000001)
	dp[1] = 0
	dp[2] = 1
	dp[3] = 1

	for i := 4; i < 1000001; i++ {
		min := 1000000
		if i%3 == 0 && min > dp[i/3]+1 {
			min = dp[i/3] + 1
		}
		if i%2 == 0 && min > dp[i/2]+1 {
			min = dp[i/2] + 1
		}
		if min > dp[i-1]+1 {
			min = dp[i-1] + 1
		}
		dp[i] = min
	}

	fmt.Fprintln(w, dp[N])
}
