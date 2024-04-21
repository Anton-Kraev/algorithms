package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Main1C() int {
	var lineCnt int
	var lines []int
	fmt.Scanln(&lineCnt)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < lineCnt && scanner.Scan(); i++ {
		text := scanner.Text()
		value, _ := strconv.Atoi(text)
		lines = append(lines, value)
	}
	return Solve1C(lineCnt, lines)
}

func Solve1C(lineCnt int, lines []int) int {
	cnt := 0
	for _, l := range lines {
		cnt += l / 4
		l = l % 4
		if l == 2 || l == 3 {
			cnt += 2
		} else {
			cnt += l
		}
	}
	return cnt
}
