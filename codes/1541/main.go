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

	var input string
	var products []string
	var numbers []int
	fmt.Fscanln(r, &input)
	inputs := strings.Split(input, "")
	_, err := strconv.Atoi(inputs[0])
	if err == nil {
		inputs = append([]string{"+"}, inputs...)
	}

	tmp := 0
	for _, v := range inputs {
		if v == "+" || v == "-" {
			products = append(products, v)
			numbers = append(numbers, tmp)
			tmp = 0
		}
		vInt, _ := strconv.Atoi(v)
		tmp = tmp*10 + vInt
	}
	numbers = append(numbers[1:], tmp)
	result := 0

	for i, v := range numbers {
		if products[i] == "+" {
			result += v
		} else {
			for j := i; j < len(products); j++ {
				result -= numbers[j]
			}
			break
		}
	}

	//fmt.Fprintln(w, inputs)
	//fmt.Fprintln(w, products)
	//fmt.Fprintln(w, numbers)
	fmt.Fprintln(w, result)
}
