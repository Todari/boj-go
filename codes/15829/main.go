package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var L int
	var S string
	// var arr []int

	fmt.Fscanln(r, &L)
	fmt.Fscanln(r, &S)
	arr := []rune(S)

	sum := big.NewInt(0)

	for i, v := range arr {
		value := new(big.Int)
		value = value.Sub(big.NewInt(int64(v)), big.NewInt(96))

		exp := new(big.Int)
		exp = exp.Exp(big.NewInt(31), big.NewInt(int64(i)), big.NewInt(0))

		mul := new(big.Int)
		mul = mul.Mul(value, exp)

		sum = sum.Add(sum, mul)
	}

	result := new(big.Int)
	result = result.Mod(sum, big.NewInt(1234567891))
	fmt.Fprintln(w, result)
}
