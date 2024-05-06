package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dsu struct {
	p []int
}

func newDsu(size int) *Dsu {
	p := make([]int, size)
	for i := 0; i < size; i++ {
		p[i] = i
	}
	return &Dsu{p}
}

func (dsu *Dsu) find(n int) int {
	if dsu.p[n] == n {
		return n
	}
	dsu.p[n] = dsu.find(dsu.p[n])
	return dsu.p[n]
}

func (dsu *Dsu) get(n int) int {
	n = dsu.find(n)
	if n == len(dsu.p)-1 {
		return -1
	}
	dsu.p[n] = n + 1
	return n
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscanln(in, &n, &m)

	dsu := newDsu(m + 1)
	var results []int
	maxCardsStr, _ := in.ReadString('\n')
	for _, maxCardStr := range strings.Split(strings.TrimSpace(maxCardsStr), " ") {
		maxCard, _ := strconv.Atoi(maxCardStr)
		res := dsu.get(maxCard)
		if res == -1 {
			break
		}
		results = append(results, res+1)
	}

	if len(results) < n {
		fmt.Fprintln(out, -1)
	} else {
		for _, res := range results {
			fmt.Fprintf(out, "%d ", res)
		}
	}
}
