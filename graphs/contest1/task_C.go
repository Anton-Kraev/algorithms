package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	v1, v2 int
}

func main() {
	var vCnt, eCnt int
	fmt.Scanln(&vCnt, &eCnt)

	matrix := make([][]bool, vCnt, vCnt)
	for i := range matrix {
		matrix[i] = make([]bool, vCnt, vCnt)
	}
	edges := make(map[Edge]struct{}, eCnt)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < eCnt && scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), " ")
		v1, _ := strconv.Atoi(line[0])
		v2, _ := strconv.Atoi(line[1])
		v1, v2 = v1-1, v2-1
		if v1 > v2 {
			v1, v2 = v2, v1
		}
		matrix[v1][v2] = true
		matrix[v2][v1] = true
		edges[Edge{v1, v2}] = struct{}{}
	}

	res := solve3(vCnt, edges, matrix)
	fmt.Println(res)
}

func solve3(vCnt int, edges map[Edge]struct{}, matrix [][]bool) int {
	cnt := 0
	for edge, _ := range edges {
		v1, v2 := edge.v1, edge.v2
		for i := 0; i < vCnt; i++ {
			if i != v1 && i != v2 && matrix[i][v1] && matrix[i][v2] {
				cnt++
			}
		}
	}
	return cnt / 3
}
