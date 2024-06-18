package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	White int8 = iota
	Gray
	Black
)

func main() {
	var vCnt, eCnt int
	fmt.Scanln(&vCnt, &eCnt)
	edges := make([][]int, vCnt, vCnt)
	for i := range edges {
		edges[i] = make([]int, 0)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < eCnt && scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), " ")
		v1, _ := strconv.Atoi(line[0])
		v2, _ := strconv.Atoi(line[1])
		v1, v2 = v1-1, v2-1
		edges[v1] = append(edges[v1], v2)
	}

	cycle := solve2(vCnt, edges)
	if len(cycle) > 0 {
		fmt.Println("YES")
		for _, v := range cycle {
			fmt.Print(v+1, " ")
		}
	} else {
		fmt.Println("NO")
	}
}

func solve2(vCnt int, edges [][]int) []int {
	visited := make([]int8, vCnt, vCnt)
	for i, _ := range visited {
		visited[i] = White
	}
	from := make([]int, vCnt, vCnt)
	for i, _ := range from {
		from[i] = -1
	}

	var cycle []int
	for v := 0; v < vCnt && len(cycle) == 0; v++ {
		if visited[v] == White {
			cycle = dfsPath(v, visited, from, edges)
		}
	}
	return cycle
}

func dfsPath(v int, visited []int8, from []int, edges [][]int) []int {
	visited[v] = Gray
	for _, u := range edges[v] {
		if visited[u] == White {
			from[u] = v
			cycle := dfsPath(u, visited, from, edges)
			if len(cycle) > 0 {
				return cycle
			}
		} else if visited[u] == Gray {
			from[u] = v
			return getCycle(u, from)
		}
	}
	visited[v] = Black
	return []int{}
}

func getCycle(lastV int, from []int) []int {
	cycle := []int{lastV}
	for v := from[lastV]; v != lastV; v = from[v] {
		cycle = append(cycle, v)
	}
	for i := 1; i <= len(cycle)/2; i++ {
		cycle[i], cycle[len(cycle)-i] = cycle[len(cycle)-i], cycle[i]
	}
	return cycle
}
