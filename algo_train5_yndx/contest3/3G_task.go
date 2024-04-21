package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)
	points := make(map[Point]struct{})
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		points[Point{x, y}] = struct{}{}
	}

	lack := Solve3G(n, points)
	fmt.Println(len(lack))
	for _, p := range lack {
		fmt.Println(p.x, p.y)
	}
}

type Point struct {
	x, y int
}

type Addition struct {
	coords [][]Point
}

func getAddition(p1, p2 Point) Addition {
	x, y := p1.x, p1.y
	dx, dy := p2.x-x, p2.y-y
	return Addition{[][]Point{
		{Point{x + dy, y - dx}, Point{x + dx + dy, y + dy - dx}},
		{Point{x - dy, y + dx}, Point{x + dx - dy, y + dy + dx}},
	}}
}

func findPoints(pointsToFind []Point, points map[Point]struct{}) []Point {
	found := make([]Point, 0, 2)
	for _, point := range pointsToFind {
		_, exist := points[point]
		if exist {
			found = append(found, point)
		}
	}
	return found
}

func (addition Addition) checkPoints(points map[Point]struct{}) []Point {
	dir1 := findPoints(addition.coords[0], points)
	dir2 := findPoints(addition.coords[1], points)
	if len(dir1) > len(dir2) {
		return dir1
	}
	return dir2
}

func (addition Addition) getLackPoints(foundPoints []Point) []Point {
	if len(foundPoints) == 2 {
		return []Point{}
	} else if len(foundPoints) == 0 {
		return addition.coords[0]
	}

	pDir, pNum := -1, -1
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if addition.coords[i][j] == foundPoints[0] {
				pDir, pNum = i, 1-j
			}
		}
	}
	return []Point{addition.coords[pDir][pNum]}
}

func Solve3G(n int, points map[Point]struct{}) []Point {
	pointsList := make([]Point, 0, n)
	for p := range points {
		pointsList = append(pointsList, p)
	}

	if n == 1 {
		p := pointsList[0]
		return []Point{{p.x + 1, p.y}, {p.x, p.y + 1}, {p.x + 1, p.y + 1}}
	}

	max := -1
	var lack []Point
	for i := 0; i < n; i++ {
		p1 := pointsList[i]
		for j := i + 1; j < n; j++ {
			p2 := pointsList[j]
			addition := getAddition(p1, p2)
			additionPoints := addition.checkPoints(points)
			if len(additionPoints) > max {
				max = len(additionPoints)
				lack = addition.getLackPoints(additionPoints)
			}
			if max == 2 {
				return lack
			}
		}
	}

	return lack
}
