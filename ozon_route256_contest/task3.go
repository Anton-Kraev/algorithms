package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		var str string
		fmt.Fscanln(in, &str)
		firstChar := rune(str[0])
		goodStr := true
		for j, s := range str[1:] {
			if s != firstChar && rune(str[j]) != firstChar {
				goodStr = false
				break
			}
		}
		if rune(str[len(str)-1]) == firstChar && goodStr {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
