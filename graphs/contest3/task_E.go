package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const infFlow = math.MaxInt64

type FlowEdge struct {
	dest, cap, flow int64
}

type MaxFlowSearch struct {
	s, t, vCnt int64
	adj        [][]int64
	edges      []FlowEdge
	level, ptr []int64
}

func InitMaxFlowSearch(vCnt int64) *MaxFlowSearch {
	return &MaxFlowSearch{adj: make([][]int64, vCnt, vCnt), s: 0, t: vCnt - 1, vCnt: vCnt}
}

func (mfs *MaxFlowSearch) addEdge(from, to, capacity int64) {
	mfs.adj[from] = append(mfs.adj[from], int64(len(mfs.edges)))
	mfs.edges = append(mfs.edges, FlowEdge{to, capacity, 0})
	mfs.adj[to] = append(mfs.adj[to], int64(len(mfs.edges)))
	mfs.edges = append(mfs.edges, FlowEdge{from, 0, 0})
}

func (mfs *MaxFlowSearch) bfs() bool {
	mfs.level = make([]int64, mfs.vCnt, mfs.vCnt)
	for i := int64(0); i < mfs.vCnt; i++ {
		mfs.level[i] = infFlow
	}
	mfs.level[mfs.s] = 0

	q := make([]int64, 0, mfs.vCnt)
	q = append(q, mfs.s)

	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for i := 0; i < len(mfs.adj[v]); i++ {
			id := mfs.adj[v][i]
			u := mfs.edges[id].dest
			if mfs.level[u] == infFlow && mfs.edges[id].flow < mfs.edges[id].cap {
				mfs.level[u] = mfs.level[v] + 1
				q = append(q, u)
			}
		}
	}

	return mfs.level[mfs.t] != infFlow
}

func (mfs *MaxFlowSearch) dfs(v, toPush int64) int64 {
	if v == mfs.t || toPush == 0 {
		return toPush
	}

	for ; mfs.ptr[v] < int64(len(mfs.adj[v])); mfs.ptr[v]++ {
		id := mfs.adj[v][mfs.ptr[v]]
		u := mfs.edges[id].dest
		if mfs.level[u] == mfs.level[v]+1 {
			pushed := mfs.dfs(u, int64(math.Min(float64(toPush), float64(mfs.edges[id].cap-mfs.edges[id].flow))))
			if pushed > 0 {
				mfs.edges[id].flow += pushed
				mfs.edges[id^1].flow -= pushed
				return pushed
			}
		}
	}

	return 0
}

func (mfs *MaxFlowSearch) Dinic() int64 {
	var flow int64
	for mfs.bfs() {
		mfs.ptr = make([]int64, mfs.vCnt, mfs.vCnt)
		for f := mfs.dfs(mfs.s, infFlow); f > 0; f = mfs.dfs(mfs.s, infFlow) {
			flow += f
		}
	}
	return flow
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	var vCnt, eCnt int64
	fmt.Sscanf(strings.TrimSpace(firstLine), "%d %d", &vCnt, &eCnt)

	mfs := InitMaxFlowSearch(vCnt)
	for i := int64(0); i < eCnt; i++ {
		edge, _ := reader.ReadString('\n')
		var from, to, capacity int64
		fmt.Sscanf(strings.TrimSpace(edge), "%d %d %d", &from, &to, &capacity)
		mfs.addEdge(from-1, to-1, capacity)
	}

	maxFlow := mfs.Dinic()
	fmt.Println(maxFlow)
	for i := int64(0); i < eCnt*2; i += 2 {
		fmt.Println(mfs.edges[i].flow)
	}
}
