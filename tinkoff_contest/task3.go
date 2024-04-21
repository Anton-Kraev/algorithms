package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var dirs []string
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		dir := scanner.Text()
		dirs = append(dirs, dir)
	}
	sort.Strings(dirs)

	for _, dir := range dirs {
		nestingLevel := strings.Count(dir, "/")
		spaces := strings.Repeat(" ", nestingLevel*2)
		fmt.Println(spaces + strings.Split(dir, "/")[nestingLevel])
	}
}
