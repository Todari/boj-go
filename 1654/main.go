package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// func getNumber(slice []int, l int) int {
// 	counter := 0
// 	for _, v := range slice {
// 		counter += v / l
// 	}
// 	return counter
// }

// func getLength(slice []int, number, start, end int) int {
// 	mid := (start + end) / 2

// }

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var K, N int
	fmt.Fscanf(r, "%d %d\n", &K, &N)

	var lan []int
	for i := 0; i < K; i++ {
		var l int
		fmt.Fscanln(r, &l)
		lan = append(lan, l)
	}
	sort.Slice(lan, func(i, j int) bool {
		return lan[i] > lan[j]
	})
	// arr := make([]int, lan[0])
	// for i := range arr {
	// 	arr[i] = i + 1
	// }

	fmt.Fprintln(w, upper_bound(N, lan)-1)

	// result := upper_bound(N, lan)
	// if result == 1 {
	// 	fmt.Fprintln(w, 1)
	// } else {
	// 	fmt.Fprintln(w, result-1)
	// }
}

func getAmount(lan []int, length int) int {
	amount := 0
	if length == 0 {
		return 2 ^ 31 - 1
	}
	for _, v := range lan {
		amount += v / length
	}
	return amount
}

func upper_bound(target int, lan []int) int {
	start := 0
	end := lan[0] + 1

	for start < end {
		middle := (start + end) / 2
		// fmt.Println(start, middle, end, getAmount(lan, middle))
		if target <= getAmount(lan, middle) {
			start = middle + 1
		} else {
			end = middle
		}
	}
	return start
}
