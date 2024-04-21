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
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file.Write([]byte{'\n'})
	defer file.Close()

	in := bufio.NewReader(file)
	var sets [3]map[string]struct{}
	for i := 0; i < 3; i++ {
		sets[i] = make(map[string]struct{})
		_, _ = in.ReadString('\n')
		inputNums, _ := in.ReadString('\n')
		nums := strings.Split(strings.TrimSpace(inputNums), " ")
		for _, num := range nums {
			_, exist := sets[i][num]
			if !exist {
				sets[i][num] = struct{}{}
			}
		}
	}

	numberOfSets := make(map[string]int)
	for _, set := range sets {
		for num := range set {
			_, exist := numberOfSets[num]
			if !exist {
				numberOfSets[num] = 0
			}
			numberOfSets[num]++
		}
	}

	var resultSet []int
	for num, cnt := range numberOfSets {
		if cnt >= 2 {
			intNum, _ := strconv.Atoi(num)
			resultSet = append(resultSet, intNum)
		}
	}
	sort.Ints(resultSet)

	for _, i := range resultSet {
		fmt.Print(i, " ")
	}
}
