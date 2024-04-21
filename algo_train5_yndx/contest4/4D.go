package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.Open("input.txt")
	reader := bufio.NewReaderSize(in, 100000*4)

	var width, wCnt1, wCnt2 int
	line0, _ := reader.ReadString('\n')
	fmt.Sscanf(strings.TrimSpace(line0), "%d %d %d", &width, &wCnt1, &wCnt2)

	var wLens1, wLens2 []int
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	for _, w := range strings.Split(strings.TrimSpace(line1), " ") {
		wordLen, _ := strconv.Atoi(w)
		wLens1 = append(wLens1, wordLen)
	}
	for _, w := range strings.Split(strings.TrimSpace(line2), " ") {
		wordLen, _ := strconv.Atoi(w)
		wLens2 = append(wLens2, wordLen)
	}

	fmt.Println(findMinSize(width, wLens1, wLens2))
}

func findMinSize(w int, wLens1, wLens2 []int) int {
	l, r := 0, w
	minSize := math.MaxInt
	var mid int

	for l <= r {
		mid = (l + r) / 2
		h1, h2 := calcHeight(wLens1, mid), calcHeight(wLens2, w-mid)
		if h1 < minSize && h2 < minSize && h1 > 0 && h2 > 0 {
			minSize = int(math.Max(float64(h1), float64(h2)))
		}
		if h1 < h2 && h1 > 0 || h2 < 0 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return minSize
}

func calcHeight(wordLens []int, width int) int {
	height := 1
	currWidth := 0
	for _, length := range wordLens {
		if length > width {
			return -1
		}
		if currWidth+length > width {
			height++
			currWidth = 0
		}
		currWidth += length + 1
	}
	return height
}
