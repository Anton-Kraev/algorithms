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
	in := bufio.NewReader(os.Stdin)
	_, _ = in.ReadString('\n')
	strLens, _ := in.ReadString('\n')
	var lens []int
	for _, str := range strings.Fields(strLens) {
		i, _ := strconv.Atoi(str)
		lens = append(lens, i)
	}
	fmt.Println(Solve1C(lens))
}

func Solve1C(lens []int) int {
	sum := 0
	for _, i := range lens {
		sum += i
	}

	min := sum
	for _, i := range lens {
		diff := 2*i - sum
		if diff > 0 {
			min = int(math.Min(float64(min), float64(diff)))
		}
	}

	return min
}
