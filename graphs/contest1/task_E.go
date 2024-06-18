package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var vCnt, fuel int
	fmt.Scanln(&vCnt, &fuel)
	edges := make([][]int, vCnt, vCnt)
	for i := range edges {
		edges[i] = make([]int, 0)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < vCnt-1 && scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), " ")
		v1, _ := strconv.Atoi(line[0])
		v2, _ := strconv.Atoi(line[1])
		v1, v2 = v1-1, v2-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	furthestVertex, _ := findMaxPath(0, vCnt, edges)
	_, maxPathLength := findMaxPath(furthestVertex, vCnt, edges)
	if maxPathLength >= fuel {
		fmt.Println(fuel + 1)
	} else {
		fmt.Println(math.Min(float64(vCnt), float64(maxPathLength+1+(fuel-maxPathLength)/2)))
	}
}

func findMaxPath(s, vCnt int, edges [][]int) (int, int) {
	distance := make([]int, vCnt, vCnt)
	visited := make([]bool, vCnt, vCnt)
	DFS(s, visited, distance, edges)
	furthestVertex, maxDistance := s, 0
	for i := 0; i < vCnt; i++ {
		if distance[i] > maxDistance {
			furthestVertex, maxDistance = i, distance[i]
		}
	}
	return furthestVertex, maxDistance
}

func DFS(v int, visited []bool, d []int, edges [][]int) {
	visited[v] = true
	for _, u := range edges[v] {
		if !visited[u] {
			d[u] = d[v] + 1
			DFS(u, visited, d, edges)
		}
	}
}
