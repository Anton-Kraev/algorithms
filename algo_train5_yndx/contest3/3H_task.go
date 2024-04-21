package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n int
	fmt.Scanln(&n)
	initialSticks := make([]Stick, 0, n)
	resultingSticks := make([]Stick, 0, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2*n && scanner.Scan(); i++ {
		var xs, ys, xf, yf int
		fmt.Sscanf(scanner.Text(), "%d %d %d %d", &xs, &ys, &xf, &yf)
		if i < n {
			initialSticks = append(initialSticks, Stick{xs, ys, xf, yf})
		} else {
			resultingSticks = append(resultingSticks, Stick{xs, ys, xf, yf})
		}
	}
	fmt.Println(solve3G(n, initialSticks, resultingSticks))
}

type Stick struct {
	xs, ys, xf, yf int
}

type Translation struct {
	dx, dy int
}

func solve3G(n int, initialSticks, resultingSticks []Stick) int {
	translations := make(map[Translation]int)
	for _, stickFrom := range initialSticks {
		for _, stickTo := range resultingSticks {
			trans := Translation{math.MaxInt, math.MaxInt}
			if stickFrom.xs-stickTo.xs == stickFrom.xf-stickTo.xf && stickFrom.ys-stickTo.ys == stickFrom.yf-stickTo.yf {
				trans = Translation{stickFrom.xs - stickTo.xs, stickFrom.ys - stickTo.ys}
			} else if stickFrom.xs-stickTo.xf == stickFrom.xf-stickTo.xs && stickFrom.ys-stickTo.yf == stickFrom.yf-stickTo.ys {
				trans = Translation{stickFrom.xs - stickTo.xf, stickFrom.ys - stickTo.yf}
			}
			if trans.dx != math.MaxInt && trans.dy != math.MaxInt {
				_, exist := translations[trans]
				if !exist {
					translations[trans] = 0
				}
				translations[trans]++
			}
		}
	}

	max := 0
	for _, v := range translations {
		max = int(math.Max(float64(max), float64(v)))
	}
	return n - max
}
