package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var M, N int
	fmt.Fscanf(r, "%d %d\n", &M, &N)

	var eratos = map[int]bool{1: false}
	for i := 2; i < N+1; i++ {
		eratos[i] = true
	}

	for i := 2; i*i <= N; i += 1 {
		if eratos[i] {
			for j := i * i; j <= N; j += i {
				eratos[j] = false
			}
		}
	}
	for i := M; i < N+1; i++ {
		if eratos[i] {
			fmt.Fprintln(w, i)
		}
	}
}
