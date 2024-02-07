package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M int
	var arr []int

	s.Scan()
	N, _ = strconv.Atoi(strings.Split(s.Text(), " ")[0])
	M, _ = strconv.Atoi(strings.Split(s.Text(), " ")[1])
	s.Scan()
	inputArr := strings.Split(s.Text(), " ")

	for _, v := range inputArr {
		value, _ := strconv.Atoi(v)
		arr = append(arr, value)
	}

	var max int

	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			for k := j + 1; k < N; k++ {
				sum := arr[i] + arr[j] + arr[k]

				if sum <= M {
					max = int(math.Max(float64(max), float64(sum)))
				}
			}
		}
	}

	fmt.Fprintln(w, max)
}
