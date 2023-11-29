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
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)
	var result int = 0

	length := int(math.Log10(float64(N)))

	for i := N - 9*(length+1); i < N; i++ {
		sum := i
		arr := strings.Split(strconv.Itoa(i), "")
		for _, v := range arr {
			value, _ := strconv.Atoi(v)
			sum += value
		}

		if N == sum {
			result = i
			break
		}
	}

	fmt.Fprintln(w, result)
}
