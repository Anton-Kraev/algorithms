package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const eps = 1e-7

type Point struct {
	x, y float64
}

type Line struct {
	a, b, c float64
}

type Trapezoid struct {
	points [4]Point
}

func lineByTwoPoints(p1, p2 Point) Line {
	if math.Abs(p1.y-p2.y) < eps {
		return Line{0, 1, -p1.y}
	} else if math.Abs(p1.x-p2.x) < eps {
		return Line{1, 0, -p1.x}
	} else {
		return Line{1 / (p2.x - p1.x), 1 / (p1.y - p2.y), p1.x/(p1.x-p2.x) + p1.y/(p2.y-p1.y)}
	}
}

func lineIntersection(l1, l2 Line) Point {
	if math.Abs(l1.b) != 0 && math.Abs(l2.b) != 0 {
		x := -(l1.b*l2.c/l2.b - l1.c) / (l1.b*l2.a/l2.b - l1.a)
		y := -(l1.c + l1.a*x) / l1.b
		return Point{x, y}
	} else if math.Abs(l1.a) != 0 && math.Abs(l2.a) != 0 {
		y := -(l1.a*l2.c/l2.a - l1.c) / (l1.a*l2.b/l2.a - l1.b)
		x := -(l1.c + l1.b*y) / l1.a
		return Point{x, y}
	} else if math.Abs(l1.a) == 0 {
		return Point{-l2.c, -l1.c}
	} else {
		return Point{-l1.c, -l2.c}
	}
}

func trapezoidArea(tr Trapezoid) float64 {
	return (tr.points[3].x - tr.points[0].x + tr.points[2].x - tr.points[1].x) * (tr.points[0].y - tr.points[1].y) / 2
}

func main() {
	var (
		n    int
		rain float64
	)
	fmt.Scanln(&n, &rain)

	scanner := bufio.NewScanner(os.Stdin)
	var points []Point
	for i := 0; i <= n && scanner.Scan(); i++ {
		var x, y float64
		fmt.Sscanf(scanner.Text(), "%f %f", &x, &y)
		if i == 0 {
			points = append(points, Point{x, 1e+15})
		}
		points = append(points, Point{x, y})
		if i == n {
			points = append(points, Point{x, 1e+15})
		}
	}

	maxH, _ := findMaxH(1, n+1, rain*(points[n+2].x-points[0].x), points)
	fmt.Println(maxH)
}

func findMaxH(l, r int, rain float64, points []Point) (float64, float64) {
	if l > r || l <= 0 || r >= len(points)-1 {
		return 0.0, rain
	}
	if l == r {
		lp, p, rp := points[l-1], points[l], points[l+1]
		ll, rl := lineByTwoPoints(lp, p), lineByTwoPoints(rp, p)

		var maxArea float64
		maxH := binSearchH(0.0, math.Min(lp.y, rp.y)-p.y, func(h float64) bool {
			tl := lineByTwoPoints(Point{0, p.y + h}, Point{1, p.y + h})
			trapezoid := Trapezoid{[4]Point{lineIntersection(ll, tl), p, p, lineIntersection(rl, tl)}}
			area := trapezoidArea(trapezoid)
			maxArea = area
			return rain-area > 0
		})

		return maxH, rain - maxArea
	}

	maxY, maxIdx := points[l].y, l
	for i := l + 1; i <= r; i++ {
		if points[i].y > maxY {
			maxY = points[i].y
			maxIdx = i
		}
	}

	lPart, rPart := 0.0, 0.0
	if l == maxIdx {
		rPart = 1
	} else if r == maxIdx {
		lPart = 1
	} else {
		lPoint, rPoint := points[l], points[r]
		for i := l - 1; i >= 0; i-- {
			if points[i].y > maxY {
				lPoint = points[i]
				if i-1 < 0 || points[i-1].y < points[i].y {
					break
				}
			}
		}
		for i := r + 1; i < len(points); i++ {
			if points[i].y > maxY {
				rPoint = points[i]
				if i+1 >= len(points) || points[i+1].y < points[i].y {
					break
				}
			}
		}
		fmt.Println(lPoint.y, rPoint.y, lPoint.x, rPoint.x)
		lPart = (points[maxIdx].x - lPoint.x) / (rPoint.x - lPoint.x)
		rPart = (rPoint.x - points[maxIdx].x) / (rPoint.x - lPoint.x)
	}

	var lMaxH, lRainLeft, rMaxH, rRainLeft float64
	if lPart > eps && rPart < eps {
		lMaxH, lRainLeft = findMaxH(l, maxIdx-1, rain, points)
	} else if rPart > eps && lPart < eps {
		rMaxH, rRainLeft = findMaxH(maxIdx+1, r, rain, points)
	} else if rPart > eps && lPart > eps {
		rMaxH, rRainLeft = findMaxH(maxIdx+1, r, rain*rPart, points)
		lMaxH, lRainLeft = findMaxH(l, maxIdx-1, rain*lPart, points)
		if lRainLeft > eps && rRainLeft < eps {
			rMaxH, rRainLeft = findMaxH(maxIdx+1, r, rain*rPart+lRainLeft, points)
			lRainLeft = 0
		} else if rRainLeft > eps && lRainLeft < eps {
			lMaxH, lRainLeft = findMaxH(l, maxIdx-1, rain*lPart+rRainLeft, points)
			rRainLeft = 0
		}
	}

	if lRainLeft+rRainLeft < eps {
		return math.Max(lMaxH, rMaxH), 0
	}

	var higherL, higherR float64
	bl := lineByTwoPoints(points[maxIdx], points[maxIdx])
	var ll, rl Line
	for i := maxIdx - 1; i >= 0; i-- {
		if points[i].y > maxY {
			higherL = points[i].y
			ll = lineByTwoPoints(points[i], points[i+1])
			break
		}
	}
	for i := maxIdx + 1; i < len(points); i++ {
		if points[i].y > maxY {
			higherR = points[i].y
			rl = lineByTwoPoints(points[i], points[i-1])
			break
		}
	}
	plb, prb := lineIntersection(ll, bl), lineIntersection(rl, bl)

	var maxArea float64
	maxH := binSearchH(0.0, math.Min(higherL, higherR)-maxY, func(h float64) bool {
		tl := lineByTwoPoints(Point{0, maxY + h}, Point{1, maxY + h})
		trapezoid := Trapezoid{[4]Point{lineIntersection(ll, tl), plb, prb, lineIntersection(rl, tl)}}
		area := trapezoidArea(trapezoid)
		maxArea = area
		return lRainLeft+rRainLeft-area > 0
	})

	return maxH + math.Max(lMaxH, rMaxH), lRainLeft + rRainLeft - maxArea
}

func binSearchH(l, r float64, check func(float64) bool) float64 {
	var mid float64
	for l+eps < r {
		mid = 0.5 * (l + r + eps)
		if check(mid) {
			l = mid
		} else {
			r = mid - eps
		}
	}
	return mid
}
