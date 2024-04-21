package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var w, h, n int
	fmt.Scanln(&w, &h, &n)
	var tiles []Coo
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		var x, y int
		fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		tiles = append(tiles, Coo{x, y})
	}

	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i].x < tiles[j].x
	})

	prefixes, suffixes := make([]Range, n, n), make([]Range, n, n)
	prefixes[0], suffixes[n-1] = Range{tiles[0].y, tiles[0].y}, Range{tiles[n-1].y, tiles[n-1].y}
	for i := 0; i < n-1; i++ {
		prefixes[i+1] = Range{
			int(math.Min(float64(prefixes[i].min), float64(tiles[i+1].y))),
			int(math.Max(float64(prefixes[i].max), float64(tiles[i+1].y))),
		}
	}
	for i := n - 1; i > 0; i-- {
		suffixes[i-1] = Range{
			int(math.Min(float64(suffixes[i].min), float64(tiles[i-1].y))),
			int(math.Max(float64(suffixes[i].max), float64(tiles[i-1].y))),
		}
	}

	min := binS(1, int(math.Min(float64(w), float64(h))), func(lineW int) bool {
		l, r := 0, 0
		for r < n {
			if tiles[r].x-tiles[l].x+1 > lineW {
				min, max := suffixes[r].min, suffixes[r].max
				if l > 0 {
					min = int(math.Min(float64(min), float64(prefixes[l-1].min)))
					max = int(math.Max(float64(max), float64(prefixes[l-1].max)))
				}
				if max-min+1 <= lineW {
					return true
				}
				l++
			} else if r == n-1 {
				return l == 0 || prefixes[l-1].max-prefixes[l-1].min+1 <= lineW
			} else {
				r++
			}
		}
		return false
	})

	fmt.Println(min)
}

type Range struct {
	min, max int
}

type Coo struct {
	x, y int
}

func binS(low, high int, check func(int) bool) int {
	var mid int

	for low < high {
		mid = (low + high) / 2
		if check(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
