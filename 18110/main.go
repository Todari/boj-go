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

	var n int
	fmt.Fscanln(r, &n)

	if n == 0 {
		fmt.Fprintln(w, 0)
	} else {
		var difficulties []int
		for i := 0; i < n; i++ {
			var d int
			fmt.Fscanln(r, &d)
			difficulties = append(difficulties, d)
		}

		sort.Slice(difficulties, func(i, j int) bool {
			return difficulties[i] < difficulties[j]
		})

		m := int(math.Round(float64(len(difficulties)) * 15 / 100))
		meanDifficulties := difficulties[m : len(difficulties)-m]

		sum := 0
		for _, v := range meanDifficulties {
			sum += v
		}

		result := int(math.Round(float64(sum) / float64(len(meanDifficulties))))
		fmt.Fprintln(w, result)
	}

}
