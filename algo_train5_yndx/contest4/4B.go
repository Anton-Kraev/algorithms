package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k uint64
	fmt.Scanln(&n)
	cbrt := uint64(math.Cbrt(float64(n)))
	k = binSearchShips(0, cbrt*2, func(size uint64) bool { return checkSize(size, n) })
	fmt.Println(k)
}

func binSearchShips(low, high uint64, check func(uint64) bool) uint64 {
	var mid uint64

	for low < high {
		mid = (low + high + 1) / 2
		if check(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return low
}

func checkSize(size uint64, max uint64) bool {
	return (size*(size+1)*(size+5)-6)/6 <= max
}
