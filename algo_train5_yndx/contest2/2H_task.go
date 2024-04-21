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
	var n, m int
	fmt.Scanln(&n, &m)
	var characters [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), " ")
		characters = append(characters, make([]int, m, m))
		for j, el := range line {
			strength, _ := strconv.Atoi(el)
			characters[i][j] = strength
		}
	}

	race, class := Solve1H(characters)
	fmt.Println(race, class)
}

func Solve1H(characters [][]int) (int, int) {
	cols := transpose(characters)
	sortedRows := sortRows(characters)
	sortedCols := sortRows(cols)

	maxRow1 := findMaxRow(sortedRows)
	maxCol1 := findMaxRow(sortRows(transpose(removeRow(characters, maxRow1))))
	choice1 := removeElement(append(characters[maxRow1], cols[maxCol1]...), characters[maxRow1][maxCol1])

	maxCol2 := findMaxRow(sortedCols)
	maxRow2 := findMaxRow(sortRows(transpose(removeRow(cols, maxCol2))))
	choice2 := removeElement(append(characters[maxRow2], cols[maxCol2]...), characters[maxRow2][maxCol2])

	res := findMaxRow(sortRows([][]int{choice1, choice2}))
	if res == 0 {
		return maxRow1 + 1, maxCol1 + 1
	}
	return maxRow2 + 1, maxCol2 + 1
}

func findMaxRow(matrix [][]int) int {
	var idxs []int
	for i := 0; i < len(matrix); i++ {
		idxs = append(idxs, i)
	}

	for step := 0; step < len(matrix[0]); step++ {
		var commonRows []int
		max := math.MinInt
		for _, idx := range idxs {
			if matrix[idx][step] > max {
				max = matrix[idx][step]
				commonRows = []int{idx}
			} else if matrix[idx][step] == max {
				commonRows = append(commonRows, idx)
			}
		}
		idxs = commonRows
		if len(idxs) == 1 {
			break
		}
	}

	return idxs[0]
}

func transpose(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	var transposed [][]int
	for c := 0; c < m; c++ {
		col := make([]int, n, n)
		for r := 0; r < n; r++ {
			col[r] = matrix[r][c]
		}
		transposed = append(transposed, col)
	}
	return transposed
}

func sortRows(matrix [][]int) [][]int {
	var sorted [][]int
	for r := 0; r < len(matrix); r++ {
		row := make([]int, len(matrix[0]), len(matrix[0]))
		copy(row, matrix[r])
		sort.Slice(row, func(i, j int) bool {
			return row[i] > row[j]
		})
		sorted = append(sorted, row)
	}
	return sorted
}

func removeRow(matrix [][]int, row int) [][]int {
	var newMatrix [][]int
	for r := 0; r < len(matrix); r++ {
		if r == row {
			continue
		}
		rowCopy := make([]int, len(matrix[0]), len(matrix[0]))
		copy(rowCopy, matrix[r])
		newMatrix = append(newMatrix, rowCopy)
	}
	return newMatrix
}

func removeElement(row []int, target int) []int {
	var newRow []int
	for _, el := range row {
		if el != target {
			newRow = append(newRow, el)
		}
	}
	return newRow
}
