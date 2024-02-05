package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N, M int
	fmt.Fscanf(r, "%d %d\n", &N, &M)

	var maze [][]bool
	for i := 0; i < N; i++ {
		maze = append(maze, make([]bool, M))
	}

	for i := 0; i < N; i++ {
		var input string
		fmt.Fscan(r, &input)
		for j, v := range strings.Split(input, "") {
			if v == "1" {
				maze[i][j] = true
			}
		}
	}

	// for _, v := range bfs(maze) {
	// 	fmt.Fprintln(w, v)
	// }

	fmt.Fprintln(w, bfs(maze)[N-1][M-1])
}

func bfs(maze [][]bool) [][]int {
	counter := 1
	var visited [][]int
	for i := 0; i < len(maze); i++ {
		visited = append(visited, make([]int, len(maze[0])))
	}

	visited[0][0] = counter
	q := new(list.List)
	q.PushBack([]int{0, 0})

	for q.Len() != 0 {
		current := q.Front().Value.([]int)
		q.Remove(q.Front())
		// if len(nextNode(maze, current[0], current[1], visited)) != 0 {
		// 	counter += 1
		// }
		for _, v := range nextNode(maze, current[0], current[1], visited) {
			q.PushBack(v)
			if visited[v[0]][v[1]] == 0 {
				visited[v[0]][v[1]] = visited[current[0]][current[1]] + 1
			}
		}
	}
	return visited
}

func nextNode(maze [][]bool, N, M int, visited [][]int) [][]int {
	var result [][]int
	if N != len(maze)-1 && maze[N+1][M] && visited[N+1][M] == 0 {
		result = append(result, []int{N + 1, M})
	}
	if M != len(maze[0])-1 && maze[N][M+1] && visited[N][M+1] == 0 {
		result = append(result, []int{N, M + 1})
	}
	if N != 0 && maze[N-1][M] && visited[N-1][M] == 0 {
		result = append(result, []int{N - 1, M})
	}
	if M != 0 && maze[N][M-1] && visited[N][M-1] == 0 {
		result = append(result, []int{N, M - 1})
	}

	return result
}
