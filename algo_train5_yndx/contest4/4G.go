package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	var figure [][]rune
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		line := scanner.Text()
		figure = append(figure, []rune(line))
	}

	ps := make([][]int, n+1, n+1)
	for i := 0; i < n+1; i++ {
		ps[i] = make([]int, m+1, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ps[i+1][j+1] = ps[i+1][j] + ps[i][j+1] - ps[i][j]
			if figure[i][j] == '#' {
				ps[i+1][j+1]++
			}
		}
	}

	maxK := findMaxPlus(1, int(math.Ceil(math.Min(float64(n), float64(m))))/3, func(k int) bool {
		plusArea := k * k * 5
		for i := 3 * k; i < n+1; i++ {
			for j := 3 * k; j < m+1; j++ {
				if ps[i-k][j]+ps[i][j-k]+ps[i-2*k][j-k]+ps[i-k][j-2*k]+ps[i-3*k][j-2*k]+ps[i-2*k][j-3*k]-
					ps[i-3*k][j-k]-ps[i-k][j-3*k]-ps[i-2*k][j]-ps[i][j-2*k]-ps[i-2*k][j-2*k]-ps[i-k][j-k] == plusArea {
					return true
				}
			}
		}
		return false
	})
	fmt.Println(maxK)
}

func findMaxPlus(low, high int, check func(int) bool) int {
	var mid int

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
