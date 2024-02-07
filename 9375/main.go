package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var T int
	fmt.Fscanln(r, &T)

	for i := 0; i < T; i++ {
		var n int
		fmt.Fscanln(r, &n)
		typeIndex := make(map[string]int)
		var clothing []int
		counter := 0
		for j := 0; j < n; j++ {
			var clothes, types string
			fmt.Fscan(r, &clothes)
			fmt.Fscan(r, &types)
			fmt.Fscanln(r)
			idx, ok := typeIndex[types]
			if ok {
				clothing[idx] += 1
			} else {
				typeIndex[types] = counter
				clothing = append(clothing, 1)
				// clothing[counter] = append(clothing[counter], clothes)
				counter += 1
			}
		}
		for clothings := 1; clothings < len(clothing)+1; clothings++ {
			for i := 0; i < clothings; i++ {

			}
		}

		comb := make([]int, len(clothing))
		for k := 0; k < len(clothing); k++ {
			comb[k] = k
		}

		result := 0
		for _, v := range combination(comb) {
			result1 := 1
			for _, vv := range v {
				result1 *= clothing[vv]
			}
			result += result1
		}
		// fmt.Fprintln(w, clothing)
		// fmt.Fprintln(w, combination(comb))
		fmt.Fprintln(w, result)

	}

	// fmt.Fprintln(w, combination([]int{0, 1, 2, 3, 4, 5}))

}

func combination(arr []int) [][]int {
	result := [][]int{}
	n := len(arr)

	for i := 1; i < (1 << uint(n)); i++ {
		var subset []int
		for j := 0; j < n; j++ {
			if i&(1<<uint(j)) > 0 {
				subset = append(subset, arr[j])
			}
		}
		if len(subset) > 0 {
			result = append(result, subset)
		}
	}
	return result

	// result1 := [][]int{}
	// result2 := [][]int{}
	// for i := range arr {
	// 	result1 = append(result1, arr[i:i+1])
	// 	result2 = append(result2, arr[i:i+1])
	// }

	// for len(result1[0]) <= len(arr)-1 {
	// 	for _, v := range result1 {
	// 		var tmp []int
	// 		tmp = append(tmp, arr...)

	// 		for idx, value := range tmp {
	// 			if value == v[len(v)-1] {
	// 				tmp = tmp[idx+1:]
	// 				break
	// 			}
	// 		}

	// 		for _, vv := range tmp {
	// 			var tmp []int
	// 			tmp = append(tmp, v...)
	// 			tmp = append(tmp, vv)
	// 			result1 = append(result1, tmp)
	// 		}
	// 		result1 = result1[1:]
	// 	}
	// 	result2 = append(result2, result1...)
	// }

	// return result2
}
