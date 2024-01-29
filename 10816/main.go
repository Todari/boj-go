package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N, M int
	fmt.Fscanln(r, &N)

	cards := make(map[int]int)
	for i := 0; i < N; i++ {
		var number int
		fmt.Fscan(r, &number)
		cards[number] += 1
	}
	fmt.Fscanln(r)

	fmt.Fscanln(r, &M)
	question := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(r, &question[i])
	}

	var keys []int
	for i := range cards {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	// fmt.Fprintln(w, cards)
	// fmt.Fprintln(w, question)
	var result []string
	for _, v := range question {
		result = append(result, strconv.Itoa(binarySearch(keys, v, cards)))
	}

	fmt.Fprintln(w, strings.Join(result, " "))
}

func binarySearch(keys []int, v int, cards map[int]int) int {
	start := 0
	end := len(keys) - 1

	for start <= end {
		middle := (start + end) / 2
		if v < keys[middle] {
			end = middle - 1
		} else if v > keys[middle] {
			start = middle + 1
		} else {
			return cards[v]
		}
	}
	return 0
}
