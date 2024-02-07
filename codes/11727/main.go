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

	dp := make([]int, 10001)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 3
	for i := 3; i < len(dp); i++ {
		dp[i] = (dp[i-1] + 2*dp[i-2]) % 10007
	}

	fmt.Fprintln(w, dp[n])
}
