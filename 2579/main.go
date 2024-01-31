package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)
	stairs, dp := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &stairs[i])
	}

	for i := 0; i < N; i++ {
		if i == 0 {
			dp[0] = stairs[0]
		} else if i == 1 {
			dp[1] = stairs[0] + stairs[1]
		} else if i == 2 {
			dp[2] = int(math.Max(
				float64(stairs[i-1]+stairs[i]), float64(stairs[i-2]+stairs[i]),
			))
		} else {
			dp[i] = int(math.Max(
				float64(stairs[i]+stairs[i-1]+dp[i-3]), float64(stairs[i]+dp[i-2]),
			))
		}
	}

	fmt.Fprintln(w, dp[N-1])
}
