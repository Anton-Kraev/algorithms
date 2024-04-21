package main

import (
	"bufio"
	"fmt"
	"os"
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
	str1, _ := in.ReadString('\n')
	dictWords := strings.Split(strings.TrimSpace(str1), " ")
	str2, _ := in.ReadString('\n')
	text := strings.Split(strings.TrimSpace(str2), " ")

	dict := make(map[string]struct{})
	for _, word := range dictWords {
		dict[word] = struct{}{}
	}

	fmt.Println(strings.Join(replaceAll(dict, text), " "))
}

func replaceAll(dict map[string]struct{}, text []string) []string {
	for wordPos, word := range text {
		for i := 0; i < len(word); i++ {
			wordPart := word[:i+1]
			_, inDict := dict[wordPart]
			if inDict {
				text[wordPos] = wordPart
				break
			}
		}
	}
	return text
}
