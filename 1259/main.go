package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var result []string

	for {
		isNo := false
		var S string
		fmt.Fscanln(r, &S)
		if S == "0" {
			break
		}

		arr := strings.Split(S, "")
		for i := 0; i < len(arr)/2; i++ {
			if arr[i] != arr[len(arr)-i-1] {
				result = append(result, "no")
				isNo = true

				break
			}
		}

		if !isNo {
			result = append(result, "yes")
		}
	}

	for _, v := range result {
		fmt.Fprintln(w, v)
	}
}
