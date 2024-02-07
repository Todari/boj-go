package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N, M int
	fmt.Fscan(r, &N)
	fmt.Fscan(r, &M)
	fmt.Fscanln(r)

	dogam := make([]string, N+1)
	dogamIndex := make(map[string]int)
	for i := 0; i < N; i++ {
		var name string
		fmt.Fscanln(r, &name)
		dogam[i+1] = name
		dogamIndex[name] = i + 1
	}
	for i := 0; i < M; i++ {
		var question string
		fmt.Fscanln(r, &question)
		number, err := strconv.Atoi(question)
		if err != nil {
			fmt.Fprintln(w, dogamIndex[question])
		} else {
			fmt.Fprintln(w, dogam[number])
		}
	}
}
