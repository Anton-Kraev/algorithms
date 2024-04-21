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
	str, _ := in.ReadString('\n')
	pricesStr, _ := in.ReadString('\n')
	firstStr := strings.Split(strings.TrimSpace(str), " ")
	days, _ := strconv.Atoi(firstStr[0])
	diff, _ := strconv.Atoi(firstStr[1])
	var prices []int
	for _, str := range strings.Fields(pricesStr) {
		i, _ := strconv.Atoi(str)
		prices = append(prices, i)
	}
	fmt.Println(Solve1B(days, diff, prices))
}

func Solve1B(days, diff int, prices []int) int {
	max := 0
	for i := 0; i < days; i++ {
		for j := i; j < days && j <= i+diff; j++ {
			max = int(math.Max(float64(max), float64(prices[j]-prices[i])))
		}
	}
	return max
}
