package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.Open("input.txt")
	reader := bufio.NewReaderSize(in, 1000000)
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
	strArr, _ := reader.ReadString('\n')
	var numbers []int
	for _, s := range strings.Split(strings.TrimSpace(strArr), " ") {
		i, _ := strconv.Atoi(s)
		numbers = append(numbers, i)
	}
	sort.Ints(numbers)
	queriesCntStr, _ := reader.ReadString('\n')
	queriesCnt, _ := strconv.Atoi(strings.TrimSpace(queriesCntStr))
	var results []int
	for i := 0; i < queriesCnt; i++ {
		var start, final int
		line, _ := reader.ReadString('\n')
		fmt.Sscanf(strings.TrimSpace(line), "%d %d", &start, &final)
		iStart := binSearch(numbers, func(val int) bool {
			return val < start
		}, Last)
		iFinal := binSearch(numbers, func(val int) bool {
			return val > final
		}, First)
		if iStart == n {
			iStart = -1
		}
		results = append(results, iFinal-iStart-1)
	}
	for _, res := range results {
		fmt.Print(res, " ")
	}
}

const (
	First = true
	Last  = false
)

func binSearch[T any](array []T, check func(T) bool, searchType bool) int {
	low := 0
	high := len(array) - 1
	var mid int

	for low < high {
		if searchType == First {
			mid = (low + high) / 2
			if check(array[mid]) {
				high = mid
			} else {
				low = mid + 1
			}
		} else if searchType == Last {
			mid = (low + high + 1) / 2
			if check(array[mid]) {
				low = mid
			} else {
				high = mid - 1
			}
		}
	}

	if check(array[low]) {
		return low
	}
	return len(array)
}
