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
	arr, result, status := make([]int, N), make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i + 1
		fmt.Fscan(r, &result[i])
	}

	temp := 0
	for i, v := range result {
		if v == i+1 {
			status[i] = 0
		} else {
			if temp+1 == v {
				temp = v
				status[i-1] = 2
				status[i] = 2
			} else if temp-1 == v {
				temp = v
				status[i-1] = 3
				status[i] = 3
			} else {
				temp = v
				status[i] = 1
			}
		}

		temp = v
	}

	fmt.Fprintln(w, arr)
	fmt.Fprintln(w, result)
	fmt.Fprintln(w, status)
}
