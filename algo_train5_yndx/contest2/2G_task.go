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
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReaderSize(f, 1024*1024)
	strN, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(strN))
	var tests [][]int
	for i := 0; i < n; i++ {
		reader.ReadString('\n')
		var test []int
		strNums, _ := reader.ReadString('\n')
		nums := strings.Split(strings.TrimSpace(strNums), " ")
		for _, el := range nums {
			num, _ := strconv.Atoi(el)
			test = append(test, num)
		}
		lengths := subArrays(test)
		tests = append(tests, lengths)
	}

	for _, test := range tests {
		fmt.Println(len(test))
		for _, el := range test {
			fmt.Print(el, " ")
		}
		fmt.Println()
	}
}

func subArrays(array []int) []int {
	var lengths []int
	count, min := 0, math.MaxInt
	for i := 0; i < len(array); i++ {
		count++
		if array[i] < min {
			min = array[i]
		}
		if count > min {
			lengths = append(lengths, count-1)
			count = 1
			min = array[i]
		}
	}
	lengths = append(lengths, count)
	return lengths
}
