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
	var cells []Cell
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), " ")
		r, _ := strconv.Atoi(line[0])
		c, _ := strconv.Atoi(line[1])
		cells = append(cells, Cell{r, c})
	}
	fmt.Println(Solve1D(n, cells))
}

type Cell struct {
	r, c int
}

func Solve1D(n int, cells []Cell) int {
	p := n * 4
	deltas := []Cell{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for _, d := range deltas {
				if d.r+cells[i].r == cells[j].r && d.c+cells[i].c == cells[j].c {
					p -= 2
					break
				}
			}
		}
	}
	return p
}
