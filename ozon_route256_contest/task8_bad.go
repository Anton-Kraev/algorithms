package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		var (
			n, m, k int
		)

		fmt.Fscanln(in, &n, &m)
		fmt.Fscanln(in, &k)

		resources := make([][][2]int, k)
		for j := 0; j < k; j++ {
			var count int
			fmt.Fscanln(in, &count)
			locations := make([][2]int, count)
			for l := 0; l < count; l++ {
				var x, y int
				fmt.Fscanln(in, &x, &y)
				locations[l] = [2]int{x, y}
			}
			resources[j] = locations
		}

		minArea := n * m
		for x1 := 0; x1 < n; x1++ {
			for y1 := 0; y1 < m; y1++ {
				for x2 := x1; x2 < n; x2++ {
					for y2 := y1; y2 < m; y2++ {
						allResources := true
						for _, resource := range resources {
							found := false
							for _, loc := range resource {
								x, y := loc[0], loc[1]
								if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
									found = true
									break
								}
							}
							if !found {
								allResources = false
								break
							}
						}
						if allResources {
							area := (x2 - x1 + 1) * (y2 - y1 + 1)
							if area < minArea {
								minArea = area
							}
						}
					}
				}
			}
		}

		fmt.Fprintln(out, minArea)
	}
}
