package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type coord struct {
// 	x int
// 	y int
// }

func main() {
	w := bufio.NewWriter(os.Stdout)
	s := bufio.NewScanner(os.Stdin)

	defer w.Flush()

	s.Scan()
	R, _ := strconv.Atoi(strings.Split(s.Text(), " ")[0])
	C, _ := strconv.Atoi(strings.Split(s.Text(), " ")[1])

	happiness := [][]int{}

	for i := 0; i < R; i++ {
		s.Scan()
		row := strings.Split(s.Text(), " ")
		rowInt := []int{}
		for _, v := range row {
			vInt, _ := strconv.Atoi(v)
			rowInt = append(rowInt, vInt)
		}
		happiness = append(happiness, rowInt)
	}
	fmt.Println(happiness)

	totalRoute := []string{}

	for i := 0; i < 3; i++ {
		x, y := 0, 0
		visited := [][]bool{{false, false, false}, {false, false, false}, {false, false, false}}

		route := []string{}
		available := [][]string{{}}

		for {
			fmt.Println("=====================================================")
			fmt.Println(x, y)
			if len(available[len(available)-1]) == 0 {
				if y != 0 && !visited[y-1][x] {
					available[len(available)-1] = append(available[len(available)-1], "U")
				}
				if x != C-1 && !visited[y][x+1] {
					available[len(available)-1] = append(available[len(available)-1], "R")
				}
				if y != R-1 && !visited[y+1][x] {
					available[len(available)-1] = append(available[len(available)-1], "D")
				}
				if x != 0 && !visited[y][x-1] {
					available[len(available)-1] = append(available[len(available)-1], "L")
				}
				fmt.Println(available)
				if len(available[len(available)-1]) == 0 {
					available = available[:len(available)-1]
					route = route[:len(route)-1]
					fmt.Println("available 삭제")
					fmt.Println(available)

				}
			}
			visited[y][x] = true
			fmt.Println(route)
			fmt.Println(visited)

			for _, v := range available[len(available)-1] {
				switch v {
				case "U":
					y = y - 1
					route = append(route, "U")
				case "R":
					x = x + 1
					route = append(route, "R")
				case "D":
					y = y + 1
					route = append(route, "D")
				case "L":
					x = x - 1
					route = append(route, "L")
				}
				if len(route) == len(available) {
					break
				}
			}

			fmt.Println(x, y)
			fmt.Println(route)
			fmt.Println(totalRoute)
			fmt.Println(available)

			if x == C-1 && y == R-1 {
				totalRoute = append(totalRoute, strings.Join(route, ""))
				for i, v := range available[len(available)-1] {
					if route[len(route)-1] == v {
						available[len(available)-1] = append(available[len(available)-1][:i], available[len(available)-1][i+1:]...)
						if len(available[len(available)-1]) == 0 {
							available = available[:len(available)-1]
							route = route[:len(route)-1]
						}
					}
				}
				visited[y][x] = false
				switch route[len(route)-1] {
				case "U":
					y = y + 1
				case "R":
					x = x - 1
				case "D":
					y = y - 1
				case "L":
					x = x - 1
				}
				route = route[:len(route)-1]
			} else {
				arr := []string{}
				available = append(available, arr)
			}

		}
		if len(route) == 0 {
			fmt.Println("!!!")
			break
		}
	}

	fmt.Println(totalRoute)
}
