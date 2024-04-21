package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Main1E() string {
	var startProfit, foundersCnt int64
	var daysCnt int
	fmt.Scan(&startProfit, &foundersCnt, &daysCnt)
	return Solve1E(startProfit, foundersCnt, daysCnt)
}

func Solve1E(startProfit, foundersCnt int64, daysCnt int) string {
	startProfit *= 10
	for i := 0; i < 10; i++ {
		if startProfit%foundersCnt == 0 {
			return strconv.FormatInt(startProfit, 10) + strings.Repeat("0", daysCnt-1)
		}
		startProfit++
	}
	return "-1"
}
