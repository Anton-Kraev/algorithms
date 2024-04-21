package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const epsi = 1.0e-8

func main() {
	var maxDist, n int
	fmt.Scanln(&maxDist, &n)

	var players []Circle
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		var x, y, speed float64
		fmt.Sscanf(scanner.Text(), "%f %f %f", &x, &y, &speed)
		players = append(players, Circle{x, y, speed})
	}
	players = append(players, Circle{0, 0, float64(maxDist)})

	var res *Circle = nil
	floatBinSearch(0, 2000, func(t float64) bool {
		intersect := false
		for _, c := range players[:len(players)-1] {
			dist := math.Sqrt(c.x*c.x + c.y*c.y)
			if dist-epsi < float64(maxDist)+c.r*t {
				intersect = true
			} else if c.r*t+epsi > dist+float64(maxDist) {
				fmt.Println(t)
				return false
			}
		}
		newRes := checkTime(t, float64(maxDist), players)
		if newRes == nil {
			return !intersect
		}
		res = newRes
		return true
	})

	fmt.Printf("%.5f\n", res.r)
	fmt.Printf("%.5f %.5f", res.x, res.y)
}

type Circle struct {
	x, y, r float64
}

type point struct {
	x, y float64
}

func floatBinSearch(l, r float64, check func(float64) bool) float64 {
	var mid float64
	for l+epsi < r {
		mid = 0.5 * (l + r + epsi)
		if check(mid) {
			l = mid
		} else {
			r = mid - epsi
		}
	}
	return mid
}

func checkTime(time, R float64, players []Circle) *Circle {
	for i := 0; i < len(players)-1; i++ {
		x1, y1, r1 := players[i].x, players[i].y, players[i].r*time
		var a, b, c float64 = 0, 1, 0
		points := intersection(x1, y1, r1, a, b, c)

		for _, p := range points {
			x, y := p.x, p.y

			if x < -R-epsi || x > R+epsi {
				continue
			}

			covers := false
			for k := 0; k < len(players)-1; k++ {
				if k == i {
					continue
				}
				if (x-players[k].x)*(x-players[k].x)+(y-players[k].y)*(y-players[k].y)-(players[k].r*time)*(players[k].r*time) < -epsi {
					covers = true
					break
				}
			}

			if !covers {
				return &Circle{x, y, time}
			}
		}
	}

	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(players); j++ {
			circle1, circle2 := players[i], players[j]
			x1, y1, r1, x2, y2, r2 := circle1.x, circle1.y, circle1.r*time, circle2.x, circle2.y, circle2.r
			if circle2.x != 0 || circle2.y != 0 {
				r2 *= time
			}
			a, b, c := circleToLine(x1, y1, r1, x2, y2, r2)
			points := intersection(x1, y1, r1, a, b, c)

			for _, p := range points {
				x, y := p.x, p.y

				if y < -epsi || x*x+y*y-R*R > epsi {
					continue
				}

				covers := false
				for k := 0; k < len(players)-1; k++ {
					if k == i || k == j {
						continue
					}
					if (x-players[k].x)*(x-players[k].x)+(y-players[k].y)*(y-players[k].y)-(players[k].r*time)*(players[k].r*time) < -epsi {
						covers = true
						break
					}
				}

				if !covers {
					return &Circle{x, y, time}
				}
			}
		}
	}

	return nil
}

func intersection(x0, y0, r, la, lb, lc float64) []point {
	if lb != 0 {
		tmp := (lc + y0*lb) / lb
		a := 1 + la*la/lb/lb
		b := 2*la*tmp/lb - 2*x0
		c := x0*x0 + tmp*tmp - r*r
		D := b*b - 4*a*c

		if D < -epsi {
			return []point{}
		} else if D > epsi {
			x1 := (-b + math.Sqrt(D)) / (2 * a)
			x2 := (-b - math.Sqrt(D)) / (2 * a)
			y1 := (-la*x1 - lc) / lb
			y2 := (-la*x2 - lc) / lb
			return []point{{x1, y1}, {x2, y2}}
		} else {
			x := -b / 2 / a
			y := (-la*x - lc) / lb
			return []point{{x, y}}
		}
	} else {
		x := -lc / la
		a := 1.0
		b := -2 * y0
		c := y0*y0 + (x-x0)*(x-x0) - r*r
		D := b*b - 4*a*c

		if D < -epsi {
			return []point{}
		} else if D > epsi {
			y1 := (-b + math.Sqrt(D)) / (2 * a)
			y2 := (-b - math.Sqrt(D)) / (2 * a)
			return []point{{x, y1}, {x, y2}}
		} else {
			y := -b / 2 / a
			return []point{{x, y}}
		}
	}
}

func circleToLine(x1, y1, r1, x2, y2, r2 float64) (float64, float64, float64) {
	a := 2 * (x1 - x2)
	b := 2 * (y1 - y2)
	c := x2*x2 - x1*x1 + y2*y2 - y1*y1 + r1*r1 - r2*r2
	return a, b, c
}
