package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	strN, _ := in.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(strN))
	numbersStr, _ := in.ReadString('\n')
	numbersSplitted := strings.Split(strings.TrimSpace(numbersStr), " ")

	numbers := make(map[int]int)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(numbersSplitted[i])
		_, exist := numbers[num]
		if !exist {
			numbers[num] = 0
		}
		numbers[num]++
	}

	fmt.Println(solve3C(n, numbers))
}

func solve3C(n int, numbers map[int]int) int {
	if len(numbers) == 1 {
		return 0
	}

	minToDelete := n
	keys := make([]int, 0)
	for k, v := range numbers {
		keys = append(keys, k)
		minToDelete = int(math.Min(float64(minToDelete), float64(n-v)))
	}
	sort.Ints(keys)

	for i := 0; i < len(keys)-1; i++ {
		if keys[i+1]-keys[i] == 1 {
			minToDelete = int(math.Min(float64(minToDelete), float64(n-numbers[keys[i]]-numbers[keys[i+1]])))
		}
	}
	return minToDelete
}
