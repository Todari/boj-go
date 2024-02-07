package main

import (
	"bufio"
	"fmt"
	"os"
)

func fact(x int) int {
	result := 1
	for i := 1; i < x+1; i++ {
		result *= i
	}

	return result
}

func getBinomialCoefficient(n, k int) int {
	return fact(n) / (fact(k) * fact(n-k))
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N, K int
	fmt.Fscanf(r, "%d %d", &N, &K)
	fmt.Fprintln(w, getBinomialCoefficient(N, K))
}
