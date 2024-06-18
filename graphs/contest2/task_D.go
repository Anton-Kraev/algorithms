package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str1, _ := reader.ReadString('\n')
	var vCnt, eCnt int
	fmt.Sscanf(strings.TrimSpace(str1), "%d %d", &vCnt, &eCnt)

	var edges []edge
	for i := 0; i < eCnt; i++ {
		query, _ := reader.ReadString('\n')
		var v1, v2, w int
		fmt.Sscanf(strings.TrimSpace(query), "%d %d %d", &v1, &v2, &w)
		edges = append(edges, edge{v1 - 1, v2 - 1, w})
	}

	fmt.Print(kruskal(vCnt, edges))
}

func kruskal(vCnt int, edges []edge) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].lt(edges[j])
	})

	mstWeight := 0
	d := makeDsu(vCnt)
	for _, e := range edges {
		if d.merge(e.v1, e.v2) {
			mstWeight += e.w
		}
	}

	return mstWeight
}

type edge struct {
	v1, v2, w int
}

func (e edge) lt(other edge) bool {
	return e.w < other.w
}

type dsu struct {
	p []int
}

func makeDsu(size int) *dsu {
	p := make([]int, size, size)
	for i := 0; i < size; i++ {
		p[i] = i
	}
	return &dsu{p}
}

func (dsu *dsu) find(n int) int {
	if dsu.p[n] == n {
		return n
	}
	dsu.p[n] = dsu.find(dsu.p[n])
	return dsu.p[n]
}

func (dsu *dsu) merge(n, m int) bool {
	n, m = dsu.find(n), dsu.find(m)
	if n == m {
		return false
	}
	if rand.Intn(2) == 0 {
		n, m = m, n
	}
	dsu.p[n] = m
	return true
}
