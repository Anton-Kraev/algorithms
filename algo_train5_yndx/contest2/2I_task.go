package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var ships []Coord
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), " ")
		row, _ := strconv.Atoi(line[0])
		col, _ := strconv.Atoi(line[1])
		ships = append(ships, Coord{row, col})
	}
	res := Solve1I(n, ships)
	fmt.Println(res)
}

type Coord struct {
	r, c int
}

func Solve1I(n int, ships []Coord) int {
	sort.Slice(ships, func(i int, j int) bool {
		return ships[i].r < ships[j].r
	})

	minR, minC := 0, math.MaxInt
	for i, ship := range ships {
		minR += int(math.Abs(float64(i + 1 - ship.r)))

		sum := 0
		for j := 0; j < n; j++ {
			sum += int(math.Abs(float64(i + 1 - ships[j].c)))
		}
		minC = int(math.Min(float64(minC), float64(sum)))
	}

	return minR + minC
}
