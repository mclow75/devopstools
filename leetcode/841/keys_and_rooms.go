package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	cnt := 0
	dfs(rooms, visited, 0, &cnt)

	return cnt == len(rooms)
}

func dfs(rooms [][]int, visited []bool, room int, cnt *int) {
	visited[room] = true
	*cnt++
	for _, key := range rooms[room] {
		if !visited[key] {
			dfs(rooms, visited, key, cnt)
		}
	}
}

func main() {
	var ex1 = [][]int{{1}, {2}, {3}}
	fmt.Println(canVisitAllRooms(ex1))
}
