package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N, M int
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	keychain := make(map[string]string)

	for i := 0; i < N; i++ {
		var site, password string
		fmt.Fscan(r, &site)
		fmt.Fscan(r, &password)
		fmt.Fscanln(r)
		keychain[site] = password
	}

	for i := 0; i < M; i++ {
		var site string
		fmt.Fscanln(r, &site)
		fmt.Fprintln(w, keychain[site])
	}

}
