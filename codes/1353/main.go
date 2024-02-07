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

	var S, P int
	fmt.Fscanf(r, "%d %d\n", &S, &P)

	var eratos = map[int]bool{1: false}
	for i := 2; i < P+1; i++ {
		eratos[i] = true
	}

	for i := 2; i*i <= P; i += 1 {
		if eratos[i] {
			for j := i * i; j <= P; j += i {
				eratos[j] = false
			}
		}
	}
	for i := 1; i < P+1; i++ {
		if eratos[i] {
			fmt.Fprintln(w, i)
		}
	}
}
