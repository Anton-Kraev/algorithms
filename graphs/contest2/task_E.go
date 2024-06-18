package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var vCnt int
	fmt.Scanln(&vCnt)

	scanner := bufio.NewScanner(os.Stdin)
	var cities []CityPoint
	for i := 0; i < vCnt && scanner.Scan(); i++ {
		var x, y float64
		fmt.Sscanf(scanner.Text(), "%f %f", &x, &y)
		cities = append(cities, CityPoint{x, y})
	}

	graph := make([][]float64, vCnt, vCnt)
	for i := 0; i < vCnt; i++ {
		graph[i] = make([]float64, vCnt, vCnt)
	}
	for i := 0; i < vCnt; i++ {
		for j := i; j < vCnt; j++ {
			dist := cities[i].calcDistance(cities[j])
			graph[i][j], graph[j][i] = dist, dist
		}
	}

	res := solveE(vCnt, graph)
	fmt.Println(res)
}

const inf = 1e8

type CityPoint struct {
	x, y float64
}

func (c CityPoint) calcDistance(otherC CityPoint) float64 {
	return math.Sqrt(math.Pow(c.x-otherC.x, 2) + math.Pow(c.y-otherC.y, 2))
}

func solveE(vCnt int, graph [][]float64) float64 {
	used := make([]bool, vCnt, vCnt)
	minE := make([]float64, vCnt, vCnt)
	selE := make([]int, vCnt, vCnt)
	for i := 0; i < vCnt; i++ {
		minE[i], selE[i] = inf, -1
	}
	minE[0] = 0

	minWeight := 0.0
	for i := 0; i < vCnt; i++ {
		v := -1
		for j := 0; j < vCnt; j++ {
			if !used[j] && (v == -1 || minE[j] < minE[v]) {
				v = j
			}
		}
		used[v] = true
		if selE[v] != -1 {
			minWeight += graph[v][selE[v]]
		}
		for to := 0; to < vCnt; to++ {
			if graph[v][to] < minE[to] {
				minE[to] = graph[v][to]
				selE[to] = v
			}
		}
	}

	return minWeight
}
