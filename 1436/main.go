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

	var N int
	fmt.Fscanln(r, &N)

	var arr []int
	for i := 666; len(arr) < 10001; i++ {
		str := strconv.Itoa(i)

		for j := 0; j < len(str)-2; j++ {
			if string(str[j]) == "6" && string(str[j+1]) == "6" && string(str[j+2]) == "6" {
				doomNumber, _ := strconv.Atoi(str)
				arr = append(arr, doomNumber)
				break
			}
		}
	}

	fmt.Fprintln(w, arr[N-1])

}
