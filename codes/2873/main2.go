// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // type coord struct {
// // 	x int
// // 	y int
// // }

// func main() {
// 	w := bufio.NewWriter(os.Stdout)
// 	s := bufio.NewScanner(os.Stdin)

// 	defer w.Flush()

// 	s.Scan()
// 	R, _ := strconv.Atoi(strings.Split(s.Text(), " ")[0])
// 	C, _ := strconv.Atoi(strings.Split(s.Text(), " ")[1])

// 	happiness := [][]int{}

// 	for i := 0; i < R; i++ {
// 		s.Scan()
// 		row := strings.Split(s.Text(), " ")
// 		rowInt := []int{}
// 		for _, v := range row {
// 			vInt, _ := strconv.Atoi(v)
// 			rowInt = append(rowInt, vInt)
// 		}
// 		happiness = append(happiness, rowInt)
// 	}
// 	// fmt.Println(happiness)

// 	totalRoute := []string{}
// 	totalHappiness := []int{}
// 	// failRoute := [][]string{}
// 	isComplete := false

// 	// for !isComplete{
// 	x, y := 0, 0
// 	currentHappiness := happiness[y][x]
// 	// visitedX := make([]bool, C)
// 	// for i := 0; i < R; i++ {
// 	// 	visitedX[i] = append(visitedX, false)
// 	// }
// 	visited := make([][]bool, R)
// 	for i := 0; i < R; i++ {
// 		visited[i] = make([]bool, C)
// 		for j := 0; j < C; j++ {
// 			visited[i][j] = false
// 		}
// 		// visited = append(visited, visitedX)
// 	}

// 	route := []string{}
// 	available := [][]string{{}}
// 	// fmt.Println(visited)

// 	for {
// 		fmt.Println("=====================================================")
// 		fmt.Println(x, y)
// 		if x == C-1 && y == R-1 {
// 			totalRoute = append(totalRoute, strings.Join(route, ""))
// 			totalHappiness = append(totalHappiness, currentHappiness)
// 			arr := []string{}
// 			available = append(available, arr)
// 		} else {
// 			if len(route) == len(available) && len(route) != 0 {
// 				arr := []string{}
// 				available = append(available, arr)
// 			}
// 			if len(available[len(route)]) == 0 {
// 				if y != 0 && !visited[y-1][x] {
// 					available[len(route)] = append(available[len(route)], "U")
// 				}
// 				if x != C-1 && !visited[y][x+1] {
// 					available[len(route)] = append(available[len(route)], "R")
// 				}
// 				if y != R-1 && !visited[y+1][x] {
// 					available[len(route)] = append(available[len(route)], "D")
// 				}
// 				if x != 0 && !visited[y][x-1] {
// 					available[len(route)] = append(available[len(route)], "L")
// 				}
// 				visited[y][x] = true
// 			}
// 		}
// 		fmt.Println(route)
// 		fmt.Println(available)
// 		fmt.Println(totalRoute)

// 		for len(available[len(available)-1]) == 0 {

// 			if len(route) == 0 {
// 				isComplete = true
// 				break
// 			}

// 			visited[y][x] = false
// 			last := route[len(route)-1]
// 			currentHappiness -= happiness[y][x]
// 			switch last {
// 			case "U":
// 				y = y + 1
// 			case "R":
// 				x = x - 1
// 			case "D":
// 				y = y - 1
// 			case "L":
// 				x = x + 1
// 			}
// 			route = route[:len(route)-1]
// 			available = available[:len(available)-1]

// 			for i, v := range available[len(available)-1] {
// 				if last == v {
// 					available[len(available)-1] = append(available[len(available)-1][:i], available[len(available)-1][i+1:]...)
// 				}
// 			}

// 		}

// 		if isComplete {
// 			break
// 		}

// 		for _, v := range available[len(available)-1] {
// 			switch v {
// 			case "U":
// 				route = append(route, "U")
// 				y = y - 1
// 			case "R":
// 				x = x + 1
// 				route = append(route, "R")
// 			case "D":
// 				y = y + 1
// 				route = append(route, "D")
// 			case "L":
// 				x = x - 1
// 				route = append(route, "L")
// 			}
// 			currentHappiness += happiness[y][x]
// 			if len(route) == len(available) {
// 				break
// 			}
// 		}

// 		// fmt.Println(x, y)
// 		// fmt.Println(route)
// 		// fmt.Println(available)
// 		// fmt.Println(totalRoute)

// 	}
// 	// }

// 	// fmt.Println(totalRoute)
// 	// fmt.Println(totalHappiness)

// 	maxHappiness := 0
// 	maxIndex := 0

// 	for i, v := range totalHappiness {
// 		if v > maxHappiness {
// 			maxHappiness = v
// 			maxIndex = i
// 		}
// 	}

// 	fmt.Fprintln(w, totalRoute[maxIndex])
// }
