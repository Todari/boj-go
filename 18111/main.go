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

	var N, M, B int
	fmt.Fscan(r, &N)
	fmt.Fscan(r, &M)
	fmt.Fscan(r, &B)
	fmt.Fscanln(r)

	heights := make([]int, N*M)
	minTime := 2147483647
	height := 0

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Fscan(r, &heights[i*M+j])
		}
		fmt.Fscanln(r)
	}

	sum := 0
	for _, v := range heights {
		sum += v
	}
	sort.Ints(heights)
	// fmt.Println(heights)

	for i := heights[0]; i < heights[len(heights)-1]+1; i++ {
		time := 0
		bag := B
		result := i * N * M
		// fmt.Println("i", i)
		if result > sum && bag < result-sum {
			continue
		}

		for _, v := range heights {
			need := i - v
			if need > 0 {
				time += need
			} else {
				time -= need * 2
			}
		}
		// fmt.Println(time, minTime)
		if time <= minTime {
			minTime = time
			height = i
		}
	}

	fmt.Fprintln(w, minTime, height)

}
