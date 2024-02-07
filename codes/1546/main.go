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
	s := bufio.NewScanner(os.Stdin)
	defer w.Flush()

	s.Scan()
	s.Scan()

	var arr []int
	inputArr := strings.Split(s.Text(), " ")

	for _, v := range inputArr {
		value, _ := strconv.Atoi(v)
		arr = append(arr, value)

		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
	}
	sum := 0.0

	for _, v := range arr {
		sum += float64(v) / float64(arr[0]) * 100
	}

	result := float64(sum) / float64(len(arr))
	fmt.Fprintln(w, result)
}
