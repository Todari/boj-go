// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	w := bufio.NewWriter(os.Stdout)
// 	r := bufio.NewReader(os.Stdin)
// 	defer w.Flush()

// 	var N, M int
// 	fmt.Fscanln(r, &N)
// 	fmt.Fscanln(r, &M)

// 	network := make(map[int][]int)
// 	status := make([]bool, N+1)
// 	status[1] = true

// 	for i := 0; i < M; i++ {
// 		var com1, com2 int
// 		fmt.Fscan(r, &com1)
// 		fmt.Fscan(r, &com2)
// 		fmt.Fscanln(r)
// 		network[com1] = append(network[com1], com2)
// 		network[com2] = append(network[com2], com1)
// 	}

// 	for computer, computers := range network {
// 		if status[computer] {
// 			for _, v := range computers {
// 				status[v] = true
// 			}
// 		}
// 	}
// 	fmt.Fprintln(w, network)
// 	fmt.Fprintln(w, status)
// 	result := 0
// 	for _, v := range status {
// 		if v {
// 			result += 1
// 		}
// 	}
// 	fmt.Fprintln(w, result)
// }

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {

	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N, M int
	fmt.Fscanln(r, &N)
	fmt.Fscanln(r, &M)

	graph := make(map[int][]int)
	visited := make(map[int]bool)

	for i := 0; i < M; i++ {

		var com1, com2 int
		fmt.Fscan(r, &com1)
		fmt.Fscan(r, &com2)
		fmt.Fscanln(r)

		graph[com1] = append(graph[com1], com2)
		graph[com2] = append(graph[com2], com1)
	}

	result := 0
	dfs(1, graph, visited, &result)
	// fmt.Fprintln(w, bfs(1, graph))

	fmt.Println(result - 1)
}

func dfs(node int, graph map[int][]int, visited map[int]bool, result *int) {
	visited[node] = true
	*result++

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(neighbor, graph, visited, result)
		}
	}
}

func bfs(start int, graph map[int][]int) int {
	visited := make(map[int]bool)
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true
	count := 0

	for queue.Len() > 0 {
		current := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		count++

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
			}
		}
	}

	return count - 1 // 시작 노드는 제외
}
