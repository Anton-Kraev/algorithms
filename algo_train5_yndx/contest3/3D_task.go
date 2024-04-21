package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	strs := strings.Split(strings.TrimSpace(str), " ")
	cnt, _ := strconv.Atoi(strs[0])
	window, _ := strconv.Atoi(strs[1])

	numbersStr, _ := in.ReadString('\n')
	numbersSplitted := strings.Split(strings.TrimSpace(numbersStr), " ")
	var numbers []int
	for i := 0; i < cnt; i++ {
		num, _ := strconv.Atoi(numbersSplitted[i])
		numbers = append(numbers, num)
	}

	fmt.Println(solve3D(window, numbers))
}

func solve3D(window int, numbers []int) string {
	lastPos := make(map[int]int)

	for pos, number := range numbers {
		prevPos, met := lastPos[number]
		if met && pos-prevPos <= window {
			return "YES"
		}
		lastPos[number] = pos
	}

	return "NO"
}
