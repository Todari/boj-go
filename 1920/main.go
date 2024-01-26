package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func binarySearch(target []string, value string, start, end int) bool {
	if start > end {
		return false
	}
	mid := (start + end) / 2
	v, _ := strconv.Atoi(value)
	m, _ := strconv.Atoi(target[mid])

	if v == m {
		return true
	} else if v > m {
		return binarySearch(target, value, mid+1, end)
	} else {
		return binarySearch(target, value, start, mid-1)
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	s.Scan()
	N, _ := strconv.Atoi(s.Text())
	s.Scan()
	target := strings.Split(s.Text(), " ")
	s.Scan()
	M, _ := strconv.Atoi(s.Text())
	s.Scan()

	question := strings.Split(s.Text(), " ")

	sort.Slice(target, func(i, j int) bool {
		prev, _ := strconv.Atoi(target[i])
		next, _ := strconv.Atoi(target[j])
		return prev < next
	})

	// fmt.Println(N, M)
	if N != 0 && M != 0 {

		for _, v := range question {

			if binarySearch(target, v, 0, len(target)-1) {
				fmt.Fprintln(w, 1)
			} else {
				fmt.Fprintln(w, 0)
			}
		}
	}
}
