package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	marksCntStr, _ := reader.ReadString('\n')
	marksCnt, _ := strconv.Atoi(strings.TrimSpace(marksCntStr))

	marksStr, _ := reader.ReadString('\n')
	var marks []int
	for _, numStr := range strings.Split(strings.TrimSpace(marksStr), " ") {
		m, _ := strconv.Atoi(numStr)
		marks = append(marks, m)
	}

	res := solve1(marksCnt, marks)
	fmt.Println(res)
}

func solve1(marksCnt int, marks []int) int {
	var badMarks, goodMarks int
	max := -1

	for i := 0; i < marksCnt && i < 7; i++ {
		badMarks, goodMarks = processMark(marks[i], badMarks, goodMarks, 1)
	}
	if badMarks == 0 && marksCnt >= 7 {
		max = goodMarks
	}

	for i := 7; i < marksCnt; i++ {
		badMarks, goodMarks = processMark(marks[i], badMarks, goodMarks, 1)
		badMarks, goodMarks = processMark(marks[i-7], badMarks, goodMarks, -1)
		if goodMarks > max && badMarks == 0 {
			max = goodMarks
		}
	}

	return max
}

func processMark(mark, badMarks, goodMarks, changeCount int) (int, int) {
	if mark == 5 {
		goodMarks += changeCount
	} else if mark == 2 || mark == 3 {
		badMarks += changeCount
	}
	return badMarks, goodMarks
}
