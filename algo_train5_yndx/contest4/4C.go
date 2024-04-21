package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.Open("input.txt")
	reader := bufio.NewReaderSize(in, 100000*2*4)
	line1, _ := reader.ReadString('\n')
	line1splitted := strings.Split(strings.TrimSpace(line1), " ")
	rgtN, _ := strconv.Atoi(line1splitted[0])
	sortieN, _ := strconv.Atoi(line1splitted[1])
	strArr, _ := reader.ReadString('\n')
	var rgts []int
	for _, s := range strings.Split(strings.TrimSpace(strArr), " ") {
		i, _ := strconv.Atoi(s)
		rgts = append(rgts, i)
	}
	prefixSums := make([]int64, rgtN+1, rgtN+1)
	for i, rgt := range rgts {
		prefixSums[i+1] = prefixSums[i] + int64(rgt)
	}

	for i := 0; i < sortieN; i++ {
		var rgtCnt int
		var rgtStrength int64
		line, _ := reader.ReadString('\n')
		fmt.Sscanf(strings.TrimSpace(line), "%d %d", &rgtCnt, &rgtStrength)
		firstIdx := findFirst(0, rgtN-rgtCnt, func(idx int) int64 {
			return prefixSums[idx+rgtCnt] - prefixSums[idx] - rgtStrength
		})
		if firstIdx == -1 {
			fmt.Println(-1)
		} else {
			fmt.Println(firstIdx + 1)
		}
	}
}

func findFirst(l, r int, check func(int) int64) int {
	var mid int

	for l <= r {
		mid = (l + r) / 2
		compare := check(mid)
		if compare == 0 {
			return mid
		}
		if compare > 0 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
