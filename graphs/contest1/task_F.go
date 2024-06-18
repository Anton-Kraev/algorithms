package main

import (
	"fmt"
)

func main() {
	var start, final string
	fmt.Scanln(&start)
	fmt.Scanln(&final)

	matrix := make([][]bool, 64, 64)
	for i := range matrix {
		matrix[i] = make([]bool, 64, 64)
	}
	dx, dy := []int{-2, -2, -1, -1, 1, 1, 2, 2}, []int{-1, 1, -2, 2, -2, 2, -1, 1}
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			from := r*8 + c
			for k := 0; k < 8; k++ {
				if onBoard(r+dy[k], c+dx[k]) {
					to := (r+dy[k])*8 + (c + dx[k])
					matrix[from][to] = true
				}
			}
		}
	}

	res := solve6(cellToIdx(start), cellToIdx(final), matrix)
	for _, cell := range res {
		fmt.Println(cell)
	}
}

func onBoard(r, c int) bool {
	return r >= 0 && r <= 7 && c >= 0 && c <= 7
}

func cellToIdx(cell string) int {
	return int(56-cell[1])*8 + int(cell[0]-97)
}

func idxToCell(idx int) string {
	return string([]byte{byte(idx%8 + 97), byte(56 - idx/8)})
}

func solve6(start, final int, matrix [][]bool) []string {
	dist := make([]int, 64, 64)
	for i := range dist {
		dist[i] = 64
	}
	dist[start] = 0
	from := make([]int, 64, 64)
	for i := range from {
		from[i] = -1
	}
	queue := make(chan int, 64)
	queue <- start

	for len(queue) > 0 {
		v := <-queue
		for u, hasEdge := range matrix[v] {
			if hasEdge && dist[u] > dist[v]+1 {
				from[u] = v
				dist[u] = dist[v] + 1
				queue <- u
			}
		}
	}

	if dist[final] == 64 {
		return []string{}
	}

	idxPath := []int{final}
	for v := from[final]; v != -1; v = from[v] {
		idxPath = append(idxPath, v)
	}
	for i := 0; i < len(idxPath)/2; i++ {
		idxPath[i], idxPath[len(idxPath)-i-1] = idxPath[len(idxPath)-i-1], idxPath[i]
	}
	var cellPath []string
	for _, idx := range idxPath {
		cellPath = append(cellPath, idxToCell(idx))
	}

	return cellPath
}
