package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var maze [][]rune
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		line := scanner.Text()
		maze = append(maze, []rune(line))
	}

	dp := make([][]int, n+1, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 5, 5)
		dp[i][0], dp[i][4] = -1, -1
	}

	for r := 0; r < n; r++ {
		for c := 0; c < 3; c++ {
			max := -1
			for prevc := c; prevc < c+3; prevc++ {
				if dp[r][prevc] > max {
					max = dp[r][prevc]
				}
			}

			if max == -1 || maze[r][c] == 'W' {
				dp[r+1][c+1] = -1
			} else if maze[r][c] == '.' {
				dp[r+1][c+1] = max
			} else if maze[r][c] == 'C' {
				dp[r+1][c+1] = max + 1
			}
		}
	}

	maxMushroomsCnt := 0
	for _, row := range dp {
		for _, el := range row {
			if el > maxMushroomsCnt {
				maxMushroomsCnt = el
			}
		}
	}

	fmt.Println(maxMushroomsCnt)
}
