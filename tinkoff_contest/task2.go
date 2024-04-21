package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	var image [][]uint64
	for i := 0; i < m; i++ {
		image = append(image, make([]uint64, n, n))
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		row := strings.Split(scanner.Text(), " ")
		for j, el := range row {
			num, _ := strconv.ParseUint(el, 10, 64)
			image[j][n-i-1] = num
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(image[i][j], " ")
		}
		fmt.Println()
	}
}
