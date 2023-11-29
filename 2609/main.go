package main

import (
	"bufio"
	"fmt"
	"os"
)

func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func Lcm(x, y int) int {
	gcd := Gcd(x, y)
	return gcd * (x / gcd) * (y / gcd)
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var x, y int
	fmt.Fscanf(r, "%d %d\n", &x, &y)

	fmt.Fprintln(w, Gcd(x, y))
	fmt.Fprintln(w, Lcm(x, y))
}
