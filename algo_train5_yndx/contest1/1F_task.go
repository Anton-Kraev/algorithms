package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Main1F() string {
	var nums []int
	in := bufio.NewReader(os.Stdin)
	nStr, _ := in.ReadString('\n')
	numsStr, _ := in.ReadString('\n')
	n, _ := strconv.Atoi(strings.Trim(nStr, "\n\r\t "))
	for _, str := range strings.Fields(numsStr) {
		i, _ := strconv.Atoi(str)
		nums = append(nums, i)
	}
	return Solve1F(n, nums)
}

func Solve1F(n int, nums []int) string {
	var sb strings.Builder

	i := 0
	for ; i < n-1 && nums[i]%2 == 0; i++ {
		sb.WriteString("+")
	}
	for ; i < n-1; i++ {
		if nums[i+1]%2 == 0 {
			sb.WriteString("+")
			break
		}
		sb.WriteString("x")
	}
	if i <= n-2 {
		sb.WriteString(strings.Repeat("x", n-i-2))
	}

	return sb.String()
}

// чет + неч = неч
// неч + неч = чет
//
// чет -> ищем нечетное и складываем
// неч -> складываем с четными (+ след чет   )
//
//
