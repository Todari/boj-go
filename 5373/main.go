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
	s := bufio.NewScanner(os.Stdin)
	defer w.Flush()

	// t:w, d:y, f:r, b:o, l:g, r:b
	var result []string

	// var n int
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	for i := 0; i < n; i++ {
		// var x int

		rubiks := [][][]string{
			{{"w", "w", "w"}, {"w", "w", "w"}, {"w", "w", "w"}},
			{{"y", "y", "y"}, {"y", "y", "y"}, {"y", "y", "y"}},
			{{"r", "r", "r"}, {"r", "r", "r"}, {"r", "r", "r"}},
			{{"o", "o", "o"}, {"o", "o", "o"}, {"o", "o", "o"}},
			{{"g", "g", "g"}, {"g", "g", "g"}, {"g", "g", "g"}},
			{{"b", "b", "b"}, {"b", "b", "b"}, {"b", "b", "b"}},
		}

		var input string
		// var rotate []string
		s.Scan()
		// x, _ := strconv.Atoi(s.Text())
		s.Scan()
		input = s.Text()
		rotate := strings.Split(input, " ")

		for _, rot := range rotate {
			switch rot {
			case "R+":
				rubiks = rp(rubiks)
			case "R-":
				rubiks = rm(rubiks)
			case "L+":
				rubiks = lp(rubiks)
			case "L-":
				rubiks = lm(rubiks)
			case "F+":
				rubiks = fp(rubiks)
			case "F-":
				rubiks = fm(rubiks)
			case "B+":
				rubiks = bp(rubiks)
			case "B-":
				rubiks = bm(rubiks)
			case "U+":
				rubiks = up(rubiks)
			case "U-":
				rubiks = um(rubiks)
			case "D+":
				rubiks = dp(rubiks)
			case "D-":
				rubiks = dm(rubiks)
			}

			// fmt.Fprintln(w, rubiks)
		}

		for i := 0; i < 3; i++ {
			query := strings.Join(rubiks[0][i], "")
			result = append(result, query)
		}
		// fmt.Fprintln(w, rubiks)
	}
	for _, v := range result {
		fmt.Fprintln(w, v)
	}
}

func rp(c [][][]string) [][][]string {
	tmp := []string{}
	for _, v := range c[0] {
		tmp = append(tmp, v[2])
	}
	for i := 0; i < 3; i++ {
		c[0][i][2] = c[2][i][2]
		c[2][i][2] = c[1][i][2]
		c[1][i][2] = c[3][i][2]
		c[3][i][2] = tmp[i]
	}
	tmp2 := make([][]string, 3)
	for i := range c[5] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[5][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[5][i][j] = tmp2[2-j][i]
		}
	}

	return c
	// 20 10 00
	// 21 11 01
	// 22 21 02
}
func rm(c [][][]string) [][][]string {
	tmp := []string{}
	for _, v := range c[0] {
		tmp = append(tmp, v[2])
	}
	for i := 0; i < 3; i++ {
		c[0][i][2] = c[3][i][2]
		c[3][i][2] = c[1][i][2]
		c[1][i][2] = c[2][i][2]
		c[2][i][2] = tmp[i]
	}
	tmp2 := make([][]string, 3)
	for i := range c[5] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[5][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[5][i][j] = tmp2[j][2-i]
		}
	}
	//02 12 22
	//01 11 21
	//00 10 20
	return c
}

func lp(c [][][]string) [][][]string {
	tmp := []string{}
	for _, v := range c[0] {
		tmp = append(tmp, v[0])
	}
	for i := 0; i < 3; i++ {
		c[0][i][0] = c[3][i][0]
		c[3][i][0] = c[1][i][0]
		c[1][i][0] = c[2][i][0]
		c[2][i][0] = tmp[i]
	}
	tmp2 := make([][]string, 3)
	for i := range c[4] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[4][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[4][i][j] = tmp2[2-j][i]
		}
	}
	return c
}
func lm(c [][][]string) [][][]string {
	tmp := []string{}
	for _, v := range c[0] {
		tmp = append(tmp, v[0])
	}
	for i := 0; i < 3; i++ {
		c[0][i][0] = c[2][i][0]
		c[2][i][0] = c[1][i][0]
		c[1][i][0] = c[3][i][0]
		c[3][i][0] = tmp[i]
	}
	tmp2 := make([][]string, 3)
	for i := range c[4] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[4][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[4][i][j] = tmp2[j][2-i]
		}
	}
	return c
}

func fp(c [][][]string) [][][]string {
	tmp := []string{}
	tmp = append(tmp, c[0][2]...)
	for i := 0; i < 3; i++ {
		c[0][2][i] = c[4][2-i][2]
		c[4][2-i][2] = c[1][0][2-i]
		c[1][0][2-i] = c[5][i][0]
		c[5][i][0] = tmp[i]
	}
	tmp2 := make([][]string, 3)
	for i := range c[2] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[2][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[2][i][j] = tmp2[2-j][i]
		}
	}
	return c
}

func fm(c [][][]string) [][][]string {
	tmp := []string{}
	tmp = append(tmp, c[0][2]...)
	for i := 0; i < 3; i++ {
		c[0][2][i] = c[5][i][0]
		c[5][i][0] = c[1][0][2-i]
		c[1][0][2-i] = c[4][2-i][2]
		c[4][2-i][2] = tmp[i]
	}
	tmp2 := make([][]string, 3)
	for i := range c[2] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[2][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[2][i][j] = tmp2[j][2-i]
		}
	}
	return c
}

func bp(c [][][]string) [][][]string {
	tmp := []string{}
	tmp = append(tmp, c[0][0]...)
	for i := 0; i < 3; i++ {
		c[0][0][i] = c[5][i][2]
		c[5][i][2] = c[1][2][2-i]
		c[1][2][2-i] = c[4][2-i][0]
		c[4][2-i][0] = tmp[i]
	}
	tmp2 := make([][]string, 3)
	for i := range c[3] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[3][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[3][i][j] = tmp2[2-j][i]
		}
	}
	return c
}

func bm(c [][][]string) [][][]string {
	tmp := []string{}
	tmp = append(tmp, c[0][0]...)
	for i := 0; i < 3; i++ {
		c[0][0][i] = c[4][2-i][0]
		c[4][2-i][0] = c[1][2][2-i]
		c[1][2][2-i] = c[5][i][2]
		c[5][i][2] = tmp[i]
	}
	tmp2 := make([][]string, 3)
	for i := range c[3] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[3][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[3][i][j] = tmp2[j][2-i]
		}
	}
	return c
}

func up(c [][][]string) [][][]string {
	tmp := []string{}
	tmp = append(tmp, c[2][0]...)
	for i := 0; i < 3; i++ {
		c[2][0][i] = c[5][0][i]
		c[5][0][i] = c[3][2][2-i]
		c[3][2][2-i] = c[4][0][i]
		c[4][0][i] = tmp[i]
	}
	tmp2 := make([][]string, 3)
	for i := range c[0] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[0][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[0][i][j] = tmp2[2-j][i]
		}
	}
	return c
}

func um(c [][][]string) [][][]string {
	// tmp := c[2][0]
	tmp := []string{}
	tmp = append(tmp, c[2][0]...)
	for i := 0; i < 3; i++ {
		c[2][0][i] = c[4][0][i]
		c[4][0][i] = c[3][2][2-i]
		c[3][2][2-i] = c[5][0][i]
		c[5][0][i] = tmp[i]
		// fmt.Println(tmp)
	}
	tmp2 := make([][]string, 3)
	for i := range c[0] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[0][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[0][i][j] = tmp2[j][2-i]
		}
	}
	// c[0] = ccw(c[0])
	return c
}

func dp(c [][][]string) [][][]string {
	tmp := []string{}
	tmp = append(tmp, c[2][2]...)
	// tmp := c[2][2]
	for i := 0; i < 3; i++ {
		c[2][2][i] = c[4][2][i]
		c[4][2][i] = c[3][0][2-i]
		c[3][0][2-i] = c[5][2][i]
		c[5][2][i] = tmp[i]
	}
	tmp2 := make([][]string, 3)
	for i := range c[1] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[1][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[1][i][j] = tmp2[2-j][i]
		}
	}
	return c
}

func dm(c [][][]string) [][][]string {
	tmp := []string{}
	tmp = append(tmp, c[2][2]...)
	for i := 0; i < 3; i++ {
		c[2][2][i] = c[5][2][i]
		c[5][2][i] = c[3][0][2-i]
		c[3][0][2-i] = c[4][2][i]
		c[4][2][i] = tmp[i]
	}
	tmp2 := make([][]string, 3)
	for i := range c[1] {
		for j := 0; j < 3; j++ {
			tmp2[i] = append(tmp2[i], c[1][i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[1][i][j] = tmp2[j][2-i]
		}
	}
	return c
}

// func cw(c [][]string) [][]string {
// 	tmp2 := c
// 	var result [][]string
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			result[i][j] = tmp2[2-j][i]
// 		}
// 	}
// 	return result
// }

// func ccw(c [][]string) [][]string {
// 	tmp2 := c
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			c[i][j] = tmp2[j][2-i]
// 		}
// 	}
// 	return c
// }
