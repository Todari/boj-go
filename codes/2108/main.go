package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func getArithmeticMean(s []int) int {
	var sum int = 0
	for _, v := range s {
		sum += v
	}
	return int(math.Round(float64(sum) / float64(len(s))))
}

func getMedian(s []int) int {
	slice := s
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice[(len(s)-1)/2]
}

func getMode(s []int) int {
	var m = map[int]int{}
	for _, v := range s {
		m[v] += 1
	}
	max := 0
	var result []int
	for k, v := range m {
		if v > max {
			result = append(result[0:0], k)
			max = v
		} else if v == max {
			result = append(result, k)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	if len(result) == 1 {
		return result[0]
	}
	return result[1]
}

func getRange(s []int) int {
	slice := s
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice[len(slice)-1] - slice[0]
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)

	var slice []int
	for i := 0; i < N; i++ {
		var v int
		fmt.Fscanln(r, &v)
		slice = append(slice, v)
	}

	fmt.Fprintln(w, getArithmeticMean(slice))
	fmt.Fprintln(w, getMedian(slice))
	fmt.Fprintln(w, getMode(slice))
	fmt.Fprintln(w, getRange(slice))
}
