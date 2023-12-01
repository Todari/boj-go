package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Info struct {
	weight int
	height int
}

var arr []Info
var result []string

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)

	for i := 0; i < N; i++ {
		var weight, height int
		fmt.Fscanf(r, "%d %d\n", &weight, &height)
		arr = append(arr, Info{weight: weight, height: height})
	}

	for _, info := range arr {
		counter := 1
		for _, v := range arr {
			if info.height < v.height && info.weight < v.weight {
				counter += 1
			}
		}

		result = append(result, strconv.Itoa(counter))
	}
	// fmt.Fprintln(w, arr)
	// fmt.Fprintln(w, result)

	output := strings.Join(result, " ")
	fmt.Fprintln(w, output)
}
