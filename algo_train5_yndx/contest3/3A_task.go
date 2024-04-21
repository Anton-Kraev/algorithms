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
	in := bufio.NewReader(os.Stdin)
	strN, _ := in.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(strN))

	freqs := make(map[string]int)
	for i := 0; i < n; i++ {
		_, _ = in.ReadString('\n')
		inputGroups, _ := in.ReadString('\n')
		groups := strings.Split(strings.TrimSpace(inputGroups), " ")
		for _, group := range groups {
			_, exist := freqs[group]
			if !exist {
				freqs[group] = 0
			}
			freqs[group]++
		}
	}

	var groups []string
	for group, freq := range freqs {
		if freq == n {
			groups = append(groups, group)
		}
	}
	sort.Strings(groups)

	fmt.Println(len(groups))
	fmt.Println(strings.Join(groups, " "))
}
