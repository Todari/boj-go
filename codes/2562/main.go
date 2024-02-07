package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n float64
	var arr []float64
	var max float64 = 0
	index := 0
	for i := 0; i < 9; i++ {
		fmt.Fscanln(r, &n)
		arr = append(arr, n)
	}

	for i, v := range arr {
		if v == math.Max(max, v) {
			index = i + 1
			max = v
		}
	}

	fmt.Fprintln(w, max)
	fmt.Fprintln(w, index)
}
