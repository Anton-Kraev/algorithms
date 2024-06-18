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
	var vCnt, eCnt, queriesCnt int
	fmt.Sscanf(strings.TrimSpace(str1), "%d %d %d", &vCnt, &eCnt, &queriesCnt)

	for i := 0; i < eCnt; i++ {
		reader.ReadString('\n')
	}

	var queries []string
	for i := 0; i < queriesCnt; i++ {
		query, _ := reader.ReadString('\n')
		queries = append(queries, strings.TrimSpace(query))
	}

	dsuConn := newDsuConn(vCnt)
	var results []bool
	for i := queriesCnt - 1; i >= 0; i-- {
		var (
			op     string
			n1, n2 int
		)
		fmt.Sscanf(queries[i], "%s %d %d", &op, &n1, &n2)
		n1--
		n2--

		if op == "ask" {
			results = append(results, dsuConn.ask(n1, n2))
		} else if op == "cut" {
			dsuConn.add(n1, n2)
		}
	}

	for i := len(results) - 1; i >= 0; i-- {
		if results[i] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

type DsuConn struct {
	p, eCnt []int
}

func newDsuConn(size int) *DsuConn {
	p := make([]int, size, size)
	for i := 0; i < size; i++ {
		p[i] = i
	}
	return &DsuConn{p, make([]int, size, size)}
}

func (dsu *DsuConn) find(n int) int {
	if dsu.p[n] == n {
		return n
	}
	dsu.p[n] = dsu.find(dsu.p[n])
	return dsu.p[n]
}

func (dsu *DsuConn) add(n, m int) {
	n, m = dsu.find(n), dsu.find(m)
	if n == m {
		return
	}
	if rand.Intn(2) == 0 {
		n, m = m, n
	}
	dsu.p[n] = m
}

func (dsu *DsuConn) ask(n, m int) bool {
	return dsu.find(n) == dsu.find(m)
}
