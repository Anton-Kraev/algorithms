package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var m, n int
	fmt.Scanln(&m, &n)
	var figure [][]rune
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < m && scanner.Scan(); i++ {
		line := scanner.Text()
		figure = append(figure, []rune(line))
	}

	ans, fig := Solve1J(m, n, figure)

	fmt.Println(ans)
	for i := 0; i < len(fig); i++ {
		fmt.Println(strings.TrimSpace(string(fig[i])))
	}
}

type Rectangle struct {
	x0, y0, xf, yf int
}

func Solve1J(m, n int, figure [][]rune) (string, [][]rune) {
	var rects []Rectangle

	for i := 0; i < m; i++ {
		start, final := -1, -1
		for j := 0; j < n; j++ {
			if figure[i][j] == '#' {
				if start == -1 {
					start = j
					final = j - 1
				}
				final++
			}
			if (figure[i][j] == '.' || j == n-1) && start != -1 {
				isNew := true
				for k, rect := range rects {
					if rect.yf == i-1 && rect.x0 == start && rect.xf == final {
						rects[k].yf++
						isNew = false
						break
					}
				}
				if isNew {
					rects = append(rects, Rectangle{start, i, final, i})
				}
				start, final = -1, -1
			}
		}
	}

	switch len(rects) {
	case 1:
		if rects[0].x0 == rects[0].xf && rects[0].y0 == rects[0].yf {
			return "NO", [][]rune{}
		}
		var rectA, rectB Rectangle
		if rects[0].xf > rects[0].x0 {
			rectA = Rectangle{rects[0].x0, rects[0].y0, rects[0].x0, rects[0].yf}
			rectB = Rectangle{rects[0].x0 + 1, rects[0].y0, rects[0].xf, rects[0].yf}
		} else {
			rectA = Rectangle{rects[0].x0, rects[0].y0, rects[0].xf, rects[0].y0}
			rectB = Rectangle{rects[0].x0, rects[0].y0 + 1, rects[0].xf, rects[0].yf}
		}
		rects = []Rectangle{rectA, rectB}
		return "YES", newFigure(rects, figure)
	case 2:
		return "YES", newFigure(rects, figure)
	case 3:
		var min, max Rectangle
		if rects[0].x0 < rects[2].x0 {
			min, max = rects[0], rects[2]
		} else {
			min, max = rects[2], rects[0]
		}

		if min.x0 == rects[1].x0 && max.xf == rects[1].xf && min.xf == max.x0-1 {
			rects[0].yf = rects[1].yf
			rects[2].y0 = rects[1].y0
			rects = []Rectangle{rects[0], rects[2]}
			return "YES", newFigure(rects, figure)
		} else if min.x0 == max.x0 && min.xf == max.xf {
			rects[0].yf = rects[2].yf
			if rects[1].x0 == rects[2].x0 {
				rects[1].x0 = rects[2].xf + 1
			} else if rects[1].xf == rects[2].xf {
				rects[1].xf = rects[2].x0 - 1
			}
			rects = []Rectangle{rects[0], rects[1]}
			return "YES", newFigure(rects, figure)
		}
		return "NO", [][]rune{}
	default:
		return "NO", [][]rune{}
	}
}

func newFigure(rects []Rectangle, figure [][]rune) [][]rune {
	for i := rects[0].x0; i <= rects[0].xf; i++ {
		for j := rects[0].y0; j <= rects[0].yf; j++ {
			figure[j][i] = 'a'
		}
	}
	for i := rects[1].x0; i <= rects[1].xf; i++ {
		for j := rects[1].y0; j <= rects[1].yf; j++ {
			figure[j][i] = 'b'
		}
	}
	return figure
}
