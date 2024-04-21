package main

import (
	"fmt"
)

func main() {
	var str1, str2 string
	fmt.Scanln(&str1)
	fmt.Scanln(&str2)
	fmt.Println(isAnagram(str1, str2))
}

func isAnagram(str1, str2 string) string {
	letters := make(map[rune]int)
	for _, c := range str1 {
		_, exist := letters[c]
		if !exist {
			letters[c] = 0
		}
		letters[c]++
	}
	for _, c := range str2 {
		_, exist := letters[c]
		if !exist {
			letters[c] = 0
		}
		letters[c]--
	}

	for _, cnt := range letters {
		if cnt != 0 {
			return "NO"
		}
	}
	return "YES"
}
