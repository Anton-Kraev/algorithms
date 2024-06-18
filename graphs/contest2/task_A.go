package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str1, _ := reader.ReadString('\n')
	var elementsCnt, queriesCnt int
	fmt.Sscanf(strings.TrimSpace(str1), "%d %d", &elementsCnt, &queriesCnt)

	dsu := newDsu(elementsCnt)
	var results []bool
	for i := 0; i < queriesCnt; i++ {
		query, _ := reader.ReadString('\n')
		var (
			op     string
			n1, n2 int
		)
		fmt.Sscanf(strings.TrimSpace(query), "%s %d %d", &op, &n1, &n2)
		n1--
		n2--

		if op == "get" {
			results = append(results, dsu.get(n1, n2))
		} else if op == "union" {
			dsu.unite(n1, n2)
		}
	}

	for _, res := range results {
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

type Dsu struct {
	p []int
}

func newDsu(size int) *Dsu {
	p := make([]int, size, size)
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

func (dsu *Dsu) unite(n, m int) {
	n, m = dsu.find(n), dsu.find(m)
	if n == m {
		return
	}
	if rand.Intn(2) == 0 {
		n, m = m, n
	}
	dsu.p[n] = m
}

func (dsu *Dsu) get(n, m int) bool {
	return dsu.find(n) == dsu.find(m)
}
