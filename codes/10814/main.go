package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Info struct {
	age   int
	name  string
	order int
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)

	// var arr []Info
	arr := make([]Info, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &arr[i].age)
		fmt.Fscan(r, &arr[i].name)
		arr[i].order = i
		fmt.Fscanln(r)
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].age == arr[j].age {
			return arr[i].order < arr[j].order
		}
		return arr[i].age < arr[j].age
	})
	// fmt.Println(arr)

	for _, v := range arr {
		fmt.Fprintln(w, v.age, v.name)
	}
}
