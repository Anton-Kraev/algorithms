package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	v1, v2, w int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		var (
			vCnt, eCnt int
			costs      []int
			edges      []Edge
		)

		fmt.Fscanln(in, &vCnt)
		costsStr, _ := in.ReadString('\n')
		for _, costStr := range strings.Split(strings.TrimSpace(costsStr), " ") {
			cost, _ := strconv.Atoi(costStr)
			costs = append(costs, cost)
		}

		fmt.Fscanln(in, &eCnt)
		for j := 0; j < eCnt; j++ {
			var v1, v2, w int
			fmt.Fscanln(in, &v1, &v2, &w)
			edges = append(edges, Edge{v1 - 1, v2 - 1, w})
		}

		fmt.Fprintln(out, kruskal(vCnt, edges, costs))
	}
}

func kruskal(vCnt int, edges []Edge, costs []int) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	mstWeight := 0
	d := makeDsu(vCnt, costs)
	for _, e := range edges {
		if d.merge(e.v1, e.v2, e.w) {
			mstWeight += e.w
		}
	}

	connected := make([]bool, vCnt)
	for v := 0; v < vCnt; v++ {
		p := d.find(v)
		if !connected[p] {
			mstWeight += d.cost[p]
			connected[p] = true
		}
	}

	return mstWeight
}

type dsu struct {
	p, cost []int
}

func makeDsu(size int, costs []int) *dsu {
	p := make([]int, size)
	for i := 0; i < size; i++ {
		p[i] = i
	}
	return &dsu{p, costs}
}

func (dsu *dsu) find(n int) int {
	if dsu.p[n] == n {
		return n
	}
	dsu.p[n] = dsu.find(dsu.p[n])
	return dsu.p[n]
}

func (dsu *dsu) merge(n, m, w int) bool {
	n, m = dsu.find(n), dsu.find(m)
	nCost, mCost := dsu.cost[n], dsu.cost[m]
	if nCost < mCost {
		n, m = m, n
		nCost, mCost = mCost, nCost
	}
	if n == m || nCost < w {
		return false
	}
	dsu.p[n] = m
	return true
}
