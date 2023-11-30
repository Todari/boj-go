package main

import (
	"bufio"
	"fmt"
	"os"
)

// func fact(x *big.Int) *big.Int {
// 	if x.Cmp(big.NewInt(1)) == 0 {
// 		return big.NewInt(1)
// 	}
// 	newX := new(big.Int)
// 	newX = newX.Sub(x, big.NewInt(1))
// 	result := new(big.Int)
// 	result = result.Mul(x, fact(newX))
// 	return result
// }

// func getZeros(x *big.Int, counter int) int {
// 	mod := new(big.Int)
// 	mod = mod.Mod(x, big.NewInt(10))
// 	div := new(big.Int)
// 	if mod.Cmp(big.NewInt(0)) == 0 {
// 		return getZeros(div.Div(x, big.NewInt(10)), counter+1)
// 	} else {
// 		return counter
// 	}
// }

// func main() {
// 	w := bufio.NewWriter(os.Stdout)
// 	r := bufio.NewReader(os.Stdin)
// 	defer w.Flush()

// 	var N big.Int
// 	fmt.Fscanln(r, &N)

// 	counter := getZeros(fact(&N), 0)

// 	fmt.Fprintln(w, counter)
// }

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)

	fmt.Fprintln(w, getZeros(N))
}

func getZeros(x int) int {
	counter := 0
	for i := 1; i < x+1; i++ {
		y := i
		for y > 1 {
			if y%5 == 0 {
				counter += 1
				y = y / 5
				continue
			}
			break
		}
	}

	return counter
}
