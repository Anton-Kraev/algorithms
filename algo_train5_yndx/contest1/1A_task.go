package main

import (
	"fmt"
	"math"
)

func Main1A() int {
	var vStart, vDist, mStart, mDist int
	fmt.Scan(&vStart, &vDist)
	fmt.Scan(&mStart, &mDist)
	return Solve1A(vStart, vDist, mStart, mDist)
}

func Solve1A(vStart int, vDist int, mStart int, mDist int) int {
	vMin := float64(vStart - vDist)
	vMax := float64(vStart + vDist)
	mMin := float64(mStart - mDist)
	mMax := float64(mStart + mDist)

	interval := int(math.Max(vMax, mMax)-math.Min(vMin, mMin)) + 1
	hole := int(math.Max(vMin, mMin) - math.Min(vMax, mMax))
	if hole > 0 {
		return interval - hole + 1
	}
	return interval
}
