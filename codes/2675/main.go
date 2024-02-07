package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var T int
	fmt.Fscanln(r, &T)

	var result []string

	for i := 0; i < T; i++ {
		var R int
		var S string
		fmt.Fscanf(r, "%d %s\n", &R, &S)

		var str string = ""
		strArr := strings.Split(S, "")
		for _, v := range strArr {
			for j := 0; j < R; j++ {
				str += v
			}
		}
		result = append(result, str)
	}

	for _, v := range result {
		fmt.Fprintln(w, v)
	}
}
