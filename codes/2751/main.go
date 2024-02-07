package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int

	fmt.Fscanln(r, &N)

	var arr []int
	for i := 0; i < N; i++ {
		var x int
		fmt.Fscanln(r, &x)
		arr = append(arr, x)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for _, v := range arr {
		fmt.Fprintln(w, v)
	}
}
