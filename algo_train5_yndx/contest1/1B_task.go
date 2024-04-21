package main

import (
	"fmt"
)

func Main1B() int {
	var prevScore, currScore string
	var flag int
	fmt.Scan(&prevScore)
	fmt.Scan(&currScore)
	fmt.Scan(&flag)
	return Solve1B(prevScore, currScore, flag)
}

func Solve1B(prevScore string, currScore string, flag int) int {
	var prev1, prev2, curr1, curr2 int
	fmt.Sscanf(prevScore, "%d:%d", &prev1, &prev2)
	fmt.Sscanf(currScore, "%d:%d", &curr1, &curr2)

	diff := prev2 + curr2 - prev1 - curr1
	if diff < 0 {
		return 0
	}

	if flag == 2 && prev1 <= curr2 || flag == 1 && curr1+diff <= prev2 {
		return diff + 1
	}
	return diff
}
