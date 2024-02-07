package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N, M int
	fmt.Fscanf(r, "%d %d\n", &N, &M)

	var result []string

	dfs(result, N, M, w)
}

func dfs(result []string, number int, length int, w *bufio.Writer) {
	if len(result) == length {
		fmt.Fprintln(w, strings.Join(result, " "))
	}

	for i := 1; i < number+1; i++ {
		if !contains(strconv.Itoa(i), result) {
			if len(result) == 0 || result[len(result)-1] < strconv.Itoa(i) {
				result = append(result, strconv.Itoa(i))
				dfs(result, number, length, w)
				result = result[:len(result)-1]
			}

		}
	}
}

func contains(elem string, arr []string) bool {
	for _, v := range arr {
		if elem == v {
			return true
		}
	}
	return false
}
