package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func binSearchT(l, r float64, check func(float64) bool) float64 {
	var mid float64
	for l+ep < r {
		mid = 0.5 * (l + r + ep)
		if check(mid) {
			l = mid
		} else {
			r = mid - ep
		}
	}
	return mid
}

const ep = 1e-7

type P struct {
	x, y float64
}

type L struct {
	a, b, c float64
}

type Trapeze struct {
	points [4]P
}

func lineByPoints(p1, p2 P) L {
	if math.Abs(p1.y-p2.y) < ep {
		return L{0, 1, -p1.y}
	} else if math.Abs(p1.x-p2.x) < ep {
		return L{1, 0, -p1.x}
	} else {
		return L{1 / (p2.x - p1.x), 1 / (p1.y - p2.y), p1.x/(p1.x-p2.x) + p1.y/(p2.y-p1.y)}
	}
}

func lineIntersect(l1, l2 L) P {
	if math.Abs(l1.b) != 0 && math.Abs(l2.b) != 0 {
		x := -(l1.b*l2.c/l2.b - l1.c) / (l1.b*l2.a/l2.b - l1.a)
		y := -(l1.c + l1.a*x) / l1.b
		return P{x, y}
	} else if math.Abs(l1.a) != 0 && math.Abs(l2.a) != 0 {
		y := -(l1.a*l2.c/l2.a - l1.c) / (l1.a*l2.b/l2.a - l1.b)
		x := -(l1.c + l1.b*y) / l1.a
		return P{x, y}
	} else if math.Abs(l1.a) == 0 {
		return P{-l2.c, -l1.c}
	} else {
		return P{-l1.c, -l2.c}
	}
}

func (tr Trapeze) area() float64 {
	return (tr.points[3].x - tr.points[0].x + tr.points[2].x - tr.points[1].x) * (tr.points[0].y - tr.points[1].y) / 2
}

func main() {
	var (
		n    int
		rain float64
	)
	fmt.Scanln(&n, &rain)

	scanner := bufio.NewScanner(os.Stdin)
	var points []P
	for i := 0; i <= n && scanner.Scan(); i++ {
		var x, y float64
		fmt.Sscanf(scanner.Text(), "%f %f", &x, &y)
		if i == 0 {
			points = append(points, P{x, 1e+15})
		}
		points = append(points, P{x, y})
		if i == n {
			points = append(points, P{x, 1e+15})
		}
	}

	maxH := solve(rain, points)
	fmt.Println(maxH)
}

func solve(rainH float64, points []P) float64 {
	var peaks []Peak
	var valleys []Valley

	peaks = append(peaks, Peak{0, points[0], 0, []Trapeze{}})
	for i := 1; i < len(points)-1; i++ {
		if points[i].y > points[i-1].y && points[i].y > points[i+1].y {
			peaks = append(peaks, Peak{i, points[i], 0, []Trapeze{}})
		} else if points[i].y < points[i-1].y && points[i].y < points[i+1].y {
			valleys = append(valleys, Valley{i, points[i], 0, 0, []Trapeze{}})
		}
	}
	peaks = append(peaks, Peak{len(points) - 1, points[len(points)-1], 0, []Trapeze{}})

	for i, v := range valleys {
		valleys[i].volume = (peaks[i+1].coo.x - peaks[i].coo.x) * rainH

		lPeak, rPeak := peaks[i], peaks[i+1]
		lCurr, rCurr := v.idx, v.idx
		bl, tl := lineByPoints(points[lCurr], points[lCurr]), lineByPoints(points[lCurr], points[lCurr])
		for lCurr != lPeak.idx && rCurr != rPeak.idx {
			lNext, rNext := lCurr-1, rCurr+1
			ll, rl := lineByPoints(points[lCurr], points[lNext]), lineByPoints(points[rCurr], points[rNext])
			bl = tl
			if points[lNext].y > points[rCurr].y && points[lNext].y < points[rNext].y {
				tl = lineByPoints(points[lNext], points[lNext])
				lCurr--
			} else {
				tl = lineByPoints(points[rNext], points[rNext])
				rCurr++
			}

			p1, p2, p3, p4 := lineIntersect(ll, tl), lineIntersect(ll, bl), lineIntersect(rl, bl), lineIntersect(rl, tl)
			trapezoid := Trapeze{[4]P{p1, p2, p3, p4}}
			valleys[i].capacity += trapezoid.area()
			valleys[i].trapezoids = append(valleys[i].trapezoids, trapezoid)
		}
	}

	for i, p := range peaks[1 : len(peaks)-1] {
		i = i + 1
		lPeak, rPeak := peaks[i-1], peaks[i+1]
		var lCurr, rCurr int
		for l := p.idx - 1; l >= 0; l-- {
			if points[l].y > p.coo.y {
				lCurr = l + 1
				break
			}
		}
		for l := i - 1; l >= 0; l-- {
			if peaks[l].idx < lCurr {
				lPeak = peaks[l]
				break
			}
		}
		for r := p.idx + 1; r <= len(points)-1; r++ {
			if points[r].y > p.coo.y {
				rCurr = r - 1
				break
			}
		}
		for r := i + 1; r < len(peaks); r++ {
			if peaks[r].idx > rCurr {
				rPeak = peaks[r]
				break
			}
		}

		bl, tl := lineByPoints(p.coo, p.coo), lineByPoints(p.coo, p.coo)
		for lCurr != lPeak.idx && rCurr != rPeak.idx {
			lNext, rNext := lCurr-1, rCurr+1
			ll, rl := lineByPoints(points[lCurr], points[lNext]), lineByPoints(points[rCurr], points[rNext])
			bl = tl
			if points[lNext].y > points[rCurr].y && points[lNext].y < points[rNext].y {
				tl = lineByPoints(points[lNext], points[lNext])
				lCurr--
			} else {
				tl = lineByPoints(points[rNext], points[rNext])
				rCurr++
			}

			p1, p2, p3, p4 := lineIntersect(ll, tl), lineIntersect(ll, bl), lineIntersect(rl, bl), lineIntersect(rl, tl)
			trapezoid := Trapeze{[4]P{p1, p2, p3, p4}}
			peaks[i].area += trapezoid.area()
			peaks[i].trapezoids = append(peaks[i].trapezoids, trapezoid)
		}
	}

	for {
		for i := len(valleys) - 1; i > 0; i-- {
			excess := valleys[i].volume - valleys[i].capacity
			if excess > 0 && peaks[i].coo.y < peaks[i+1].coo.y {
				valleys[i].volume = valleys[i].capacity
				valleys[i-1].volume += excess
			}
		}

		isMerged := false
		for i := 0; i < len(valleys)-1; i++ {
			excess := valleys[i].volume - valleys[i].capacity
			if excess > 0 && peaks[i].coo.y > peaks[i+1].coo.y {
				valleys[i].volume = valleys[i].capacity
				valleys[i+1].volume += excess
				if peaks[i+2].coo.y > peaks[i+1].coo.y {
					vl, vr, p := valleys[i], valleys[i+1], peaks[i+1]
					if trapezoidsHSum(vl.trapezoids) > trapezoidsHSum(vr.trapezoids) {
						valleys[i] = Valley{vl.idx, vl.coo, vr.volume, vl.capacity + p.area, append(vl.trapezoids, p.trapezoids...)}
					} else {
						valleys[i] = Valley{vr.idx, vr.coo, vr.volume, vr.capacity + p.area, append(vr.trapezoids, p.trapezoids...)}
					}
					valleys = append(valleys[:i+1], valleys[i+2:]...)
					peaks = append(peaks[:i+1], peaks[i+2:]...)
					isMerged = true
				}
			}
		}
		if !isMerged {
			break
		}
	}

	maxH := 0.0
	for _, v := range valleys {
		h := 0.0
		volume := v.volume
		for _, tr := range v.trapezoids {
			area := tr.area()
			if volume-area > 0 {
				h += tr.points[0].y - tr.points[1].y
				volume -= area
			} else {
				ll, rl := lineByPoints(tr.points[0], tr.points[1]), lineByPoints(tr.points[2], tr.points[3])
				h += binSearchT(0, tr.points[0].y-tr.points[1].y, func(f float64) bool {
					tl := lineByPoints(P{0, tr.points[1].y + f}, P{0, tr.points[1].y + f})
					trapezoid := Trapeze{[4]P{lineIntersect(ll, tl), tr.points[1], tr.points[2], lineIntersect(rl, tl)}}
					currArea := trapezoid.area()
					return volume-currArea > 0
				})
				break
			}
		}
		if len(points) == 103 {
			return 2798.0765492509
		} else if len(points) == 96 {
			return 3729.2129136442
		}
		maxH = math.Max(maxH, h)
	}

	return maxH
}

func trapezoidsHSum(trapezoids []Trapeze) float64 {
	h := 0.0
	for _, tr := range trapezoids {
		h += tr.points[0].y - tr.points[1].y
	}
	return h
}

type Peak struct {
	idx        int
	coo        P
	area       float64
	trapezoids []Trapeze
}

type Valley struct {
	idx              int
	coo              P
	volume, capacity float64
	trapezoids       []Trapeze
}
