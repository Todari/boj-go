package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	counter := 0
	for i := 1; i < n+1; i++ {
		if n%i == 0 {
			counter += 1
		}
	}
	return counter == 2
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var arr []string

	s.Split(bufio.ScanLines)
	s.Scan()
	s.Scan()
	S := s.Text()
	arr = strings.Split(S, " ")

	var counter int = 0

	for _, v := range arr {
		value, _ := strconv.Atoi(v)
		if isPrime(value) {
			counter += 1
		}
	}

	fmt.Fprintln(w, counter)
}
