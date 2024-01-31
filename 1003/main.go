package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var T int
	fmt.Fscanln(r, &T)
	dp := make([][]int, 41)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	setDp(dp)
	fmt.Println("씨발")
	for i := 0; i < T; i++ {
		var question int
		fmt.Fscanln(r, &question)
		// if dp[question][0] == 0 {
		// 	fmt.Fprintln(w, dp)
		// }
		fmt.Fprintln(w, dp[question][0], dp[question][1])
	}
}

func setDp(dp [][]int) {
	// if n == 0 {
	dp[0][0] = 1
	dp[0][1] = 0
	// } else if n == 1 {
	dp[1][0] = 0
	dp[1][1] = 1
	// } else {
	for i := 2; i < 41; i++ {
		dp[i][0] = dp[i-1][0] + dp[i-2][0]
		dp[i][1] = dp[i-1][1] + dp[i-2][1]
	}
	// }
}
