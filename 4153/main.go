package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var result []string
	for {
		var input []int
		var A, B, C int
		fmt.Fscanf(r, "%d %d %d\n", &A, &B, &C)
		input = append(input, A, B, C)
		sort.Slice(input, func(i, j int) bool {
			return input[i] < input[j]
		})

		if A+B+C == 0 {

			break
		}

		if math.Pow(float64(input[2]), 2) == math.Pow(float64(input[1]), 2)+math.Pow(float64(input[0]), 2) {
			result = append(result, "right")
		} else {
			result = append(result, "wrong")
		}

	}

	for _, v := range result {
		fmt.Fprintln(w, v)
	}
}
