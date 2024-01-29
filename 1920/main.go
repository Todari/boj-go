package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func binarySearch(target []int, value int, start, end int) bool {
	for start <= end {
		middle := (start + end) / 2
		if value == target[middle] {
			return true
		} else if value > target[middle] {
			start = middle + 1
		} else if value < target[middle] {
			end = middle - 1
		}
	}
	return false
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)
	target := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &target[i])
	}
	fmt.Fscanln(r)

	var M int
	fmt.Fscanln(r, &M)
	question := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(r, &question[i])
	}
	sort.Ints(target)

	// fmt.Println(target)

	for _, v := range question {

		if binarySearch(target, v, 0, len(target)-1) {
			fmt.Fprintln(w, 1)
		} else {
			fmt.Fprintln(w, 0)
		}
	}
}
