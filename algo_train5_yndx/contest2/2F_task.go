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
	var minDeg, maxDeg, slowdown int
	var sectors []int

	in := bufio.NewReader(os.Stdin)
	line1, _ := in.ReadString('\n')
	line2, _ := in.ReadString('\n')
	line3, _ := in.ReadString('\n')

	sectorsCnt, _ := strconv.Atoi(strings.TrimSpace(line1))
	fmt.Sscanf(strings.TrimSpace(line3), "%d %d %d", &minDeg, &maxDeg, &slowdown)
	line := strings.Split(strings.TrimSpace(line2), " ")
	for i := 0; i < sectorsCnt; i++ {
		sector, _ := strconv.Atoi(line[i])
		sectors = append(sectors, sector)
	}

	max := Solve1F(sectorsCnt, minDeg, maxDeg, slowdown, sectors)
	fmt.Println(max)
}

func Solve1F(sectorsCnt, minDeg, maxDeg, slowdown int, sectors []int) int {
	minSpins := (minDeg - 1) / slowdown
	maxSpins := (maxDeg - 1) / slowdown
	start := minSpins % sectorsCnt
	final := maxSpins % sectorsCnt
	max := math.MinInt

	if maxSpins-minSpins >= sectorsCnt-1 {
		return int(findMax(sectors))
	}
	if start <= final {
		max = int(math.Max(
			findMax(sectors[start:final+1]),
			findMax(sectors[sectorsCnt-final:int(math.Min(float64(sectorsCnt), float64(sectorsCnt-start+1)))]),
		),
		)
	} else {
		max = int(math.Max(
			math.Max(
				findMax(sectors[start:]),
				findMax(sectors[:final+1]),
			),
			math.Max(
				findMax(sectors[:sectorsCnt-start+1]),
				findMax(sectors[sectorsCnt-final:]),
			),
		))
	}

	return max
}

func findMax(array []int) float64 {
	max := math.Inf(-1)
	for i := 0; i < len(array); i++ {
		max = math.Max(max, float64(array[i]))
	}
	return max
}
