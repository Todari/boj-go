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
	dp[2] = 2
	for i := 3; i < len(dp); i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 10007
	}

	// fmt.Fprintln(w, dp)

	fmt.Fprintln(w, dp[n])
}

// 1	1		1
// 2 2		1 + 1
// 3 3		1 + 2
// 4 5		1 + 3 + 1
// 5 8		1 + 4 + 3
// 6 13	1 + 5 + 6 + 1
// 7 21	1 + 6 + 10 + 4
// 8	34	1 + 7 + 15 + 10 + 1
// 9	55	1 + 8 + 21 + 20 + 5
// 10 89 1 + 9 + 28 + 35 + 15 + 1
