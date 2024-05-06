package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, t int
	fmt.Fscanln(in, &n, &t)
	letters := make(map[string]int)
	inputLetters, _ := in.ReadString('\n')
	for _, letter := range strings.Split(strings.TrimSpace(inputLetters), " ") {
		_, exist := letters[letter]
		if !exist {
			letters[letter] = 0
		}
		letters[letter]++
	}
	for i := 0; i < t; i++ {
		var password string
		fmt.Fscanln(in, &password)
		lettersCopy := make(map[string]int)
		for k, v := range letters {
			lettersCopy[k] = v
		}
		for _, l := range password {
			lettersCopy[string(l)]--
		}
		good := true
		for _, cnt := range lettersCopy {
			if cnt != 0 {
				good = false
				break
			}
		}
		if good {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
