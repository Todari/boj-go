package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N, M int
	fmt.Fscan(r, &N)
	fmt.Fscan(r, &M)
	fmt.Fscanln(r)

	notHeard := make([]string, N)
	notSeen := make([]string, M)
	var notHeardAndSeen []string

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &notHeard[i])
	}
	for i := 0; i < M; i++ {
		fmt.Fscanln(r, &notSeen[i])
	}

	sort.Strings(notSeen)
	sort.Strings(notHeard)
	if N < M {
		for _, v := range notHeard {
			if binarySearch(notSeen, v) {
				notHeardAndSeen = append(notHeardAndSeen, v)
			}
		}
	} else {
		for _, v := range notSeen {
			if binarySearch(notHeard, v) {
				notHeardAndSeen = append(notHeardAndSeen, v)
			}
		}
	}

	fmt.Fprintln(w, len(notHeardAndSeen))
	for _, v := range notHeardAndSeen {
		fmt.Fprintln(w, v)
	}
}

func binarySearch(target []string, value string) bool {
	start := 0
	end := len(target) - 1

	for start <= end {
		mid := (start + end) / 2
		if target[mid] < value {
			start = mid + 1
		} else if target[mid] > value {
			end = mid - 1
		} else {
			return true
		}
	}
	return false
}
