package main

import (
	"fmt"
	"math"
)

func main() {
	var n uint64
	fmt.Scanln(&n)
	sqrt := uint64(math.Sqrt(float64(n)))
	line := binSearchLine(1, sqrt*2, func(line uint64) bool { return checkLine(line, n) })
	numerator, denominator := getFraction(line, n)
	fmt.Printf("%d/%d", numerator, denominator)
}

func binSearchLine(low, high uint64, check func(uint64) bool) uint64 {
	var mid uint64

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

func checkLine(line, n uint64) bool {
	return line*(line+1) >= n*2
}

func getFraction(line, n uint64) (uint64, uint64) {
	order := n - line*(line-1)/2
	if line%2 == 1 {
		return order, line - order + 1
	}
	return line - order + 1, order
}
