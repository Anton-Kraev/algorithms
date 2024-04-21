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
	var berries []Berry
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), " ")
		up, _ := strconv.Atoi(line[0])
		down, _ := strconv.Atoi(line[1])
		berries = append(berries, Berry{int64(i + 1), int64(up), int64(down)})
	}
	Solve1E(berries)
}

type Berry struct {
	idx, up, down int64
}

func Solve1E(berries []Berry) {
	sort.Slice(berries, func(i, j int) bool {
		if berries[i].up-berries[i].down > 0 && berries[j].up-berries[j].down > 0 {
			return berries[i].down < berries[j].down
		} else if berries[i].up-berries[i].down <= 0 && berries[j].up-berries[j].down <= 0 {
			return berries[i].up > berries[j].up
		}
		return berries[i].up-berries[i].down > berries[j].up-berries[j].down
	})

	var curr int64 = 0
	var max int64 = 0
	for _, berry := range berries {
		curr += berry.up
		max = int64(math.Max(float64(max), float64(curr)))
		curr -= berry.down
	}

	var indices []any
	for _, berry := range berries {
		indices = append(indices, berry.idx)
	}
	fmt.Println(max)
	fmt.Println(indices...)
}
