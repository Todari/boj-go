package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	var str string
	result := 0
	fmt.Fscanln(r, &n)
	fmt.Fscanln(r, &str)

	strArr := strings.Split(str, "")
	for _, v := range strArr {
		number, _ := strconv.Atoi(v)
		result += number
	}

	fmt.Fprintln(w, result)
}
