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
	r:= bufio.NewReader(os.Stdin)
	w:= bufio.NewWriter(os.Stdout)
	defer w.Flush()

	line, _ := r.ReadString('\n')
	parts := strings.Fields(line)

	N,_ := strconv.Atoi(parts[0])
	M,_ := strconv.Atoi(parts[1])
	
	line, _ = r.ReadString('\n')
	parts = strings.Fields(line)

	trees := make([]int, N)
	for i, v := range parts {
		trees[i], _ = strconv.Atoi(v)
	}

	sort.Ints(trees)

	var left = 0
	var right = trees[N - 1]
	result:=0
	
	for left <= right {
		mid := (left + right) / 2
		remain := 0;
		for _, tree := range trees {
			if tree > mid {
				remain += tree - mid
			}
		}
		if (remain >= M) {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Fprintln(w, result)
}