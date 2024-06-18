package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		edges[v2] = append(edges[v2], v1)
	}

	compNum := 0
	components := make([]int, vCnt, vCnt)
	for v := 0; v < vCnt; v++ {
		if components[v] == 0 {
			compNum++
			dfs(v, compNum, components, edges)
		}
	}

	fmt.Println(compNum)
	for _, c := range components {
		fmt.Print(c, " ")
	}
}

func dfs(v, compNum int, components []int, edges [][]int) {
	components[v] = compNum
	for _, u := range edges[v] {
		if components[u] == 0 {
			dfs(u, compNum, components, edges)
		}
	}
}
