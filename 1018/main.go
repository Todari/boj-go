package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int
	var M int
	var chess [][]string

	fmt.Fscanf(r, "%d %d\n", &N, &M)

	for i := 0; i < N; i++ {
		var col string
		fmt.Fscanln(r, &col)
		chess = append(chess, strings.Split(col, ""))
	}

	for n := 0; n+7 < N; n++ {
		for m := 0; m+7 < M; m++ {
		}
	}

	var min int = 64
	for n := 0; n+7 < N; n++ {
		for m := 0; m+7 < M; m++ {

			var part [][]string
			for _, v := range chess[n : n+8] {
				part = append(part, v[m:m+8])
			}

			var counter1 int = 0
			var counter2 int = 0
			for i := 0; i < 8; i++ {
				for j := 0; j < 8; j++ {
					if (i+j)%2 == 0 {
						if part[i][j] == "W" {
							counter1 += 1
						} else {
							counter2 += 1
						}
					} else {
						if part[i][j] == "B" {
							counter1 += 1
						} else {
							counter2 += 1
						}
					}
				}
			}
			counter := int(math.Min(float64(counter1), float64(counter2)))

			min = int(math.Min(float64(min), float64(counter)))
		}
	}

	fmt.Fprintln(w, min)
}
