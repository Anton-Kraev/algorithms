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
	var n int
	fmt.Scanln(&n)
	var points []Point
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		points = append(points, Point{x, y})
	}
	fmt.Println(Solve1A(points))
}

type Point struct {
	x, y int
}

func Solve1A(points []Point) (int, int, int, int) {
	xMin := math.MaxInt32
	xMax := math.MinInt32
	yMin := math.MaxInt32
	yMax := math.MinInt32

	for _, p := range points {
		xMin = int(math.Min(float64(xMin), float64(p.x)))
		xMax = int(math.Max(float64(xMax), float64(p.x)))
		yMin = int(math.Min(float64(yMin), float64(p.y)))
		yMax = int(math.Max(float64(yMax), float64(p.y)))
	}

	return xMin, yMin, xMax, yMax
}
