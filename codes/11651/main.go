package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type coordinate struct {
	x int
	y int
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)

	var arr []coordinate
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscanf(r, "%d %d\n", &x, &y)
		arr = append(arr, coordinate{x: x, y: y})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].y == arr[j].y {
			return arr[i].x < arr[j].x
		}
		return arr[i].y < arr[j].y
	})

	for i, v := range arr {
		if i == len(arr)-1 {
			fmt.Fprintf(w, "%d %d", v.x, v.y)
		} else {
			fmt.Fprintf(w, "%d %d\n", v.x, v.y)
		}
	}
}
