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

	var n, b int
	fmt.Fscanf(r, "%d %d\n", &n, &b)
	var word string
	var X int
	fmt.Fscanf(r, "%s %d\n", &word, &X)
	number := make([][]int, 2*b-2)
	for i := range number {
		// arr:= make([]int, n)
		number[i] = make([]int, n)
	}

	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &words[i])
	}
}
