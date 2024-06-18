package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str1, _ := reader.ReadString('\n')
	var vCnt, eCnt int
	fmt.Sscanf(strings.TrimSpace(str1), "%d %d", &vCnt, &eCnt)

	str2, _ := reader.ReadString('\n')
	var s, f int
	fmt.Sscanf(strings.TrimSpace(str2), "%d %d", &s, &f)

	graph := make([][]WeightedEdge, vCnt, vCnt)
	for k := 0; k < eCnt; k++ {
		edge, _ := reader.ReadString('\n')
		var v, u, w int
		fmt.Sscanf(strings.TrimSpace(edge), "%d %d %d", &v, &u, &w)
		v--
		u--
		graph[v] = append(graph[v], WeightedEdge{u, w})
		graph[u] = append(graph[u], WeightedEdge{v, w})
	}

	pathLen, path := dijkstra(s-1, f-1, vCnt, graph)
	if len(path) > 0 {
		fmt.Println(pathLen)
	} else {
		fmt.Println(-1)
	}
	for _, v := range path {
		fmt.Print(v+1, " ")
	}
}

func dijkstra(s, f, vCnt int, graph [][]WeightedEdge) (int, []int) {
	d := make([]int, vCnt, vCnt)
	p := make([]int, vCnt, vCnt)
	for i := 0; i < vCnt; i++ {
		d[i] = math.MaxInt
		p[i] = -1
	}
	d[s] = 0

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Item{value: s, priority: 0})

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		v, priority := item.value, item.priority
		if priority > d[v] {
			continue
		}
		for _, e := range graph[v] {
			u, w := e.v, e.w
			if d[u] > d[v]+w {
				d[u] = d[v] + w
				p[u] = v
				heap.Push(&pq, &Item{value: u, priority: d[u]})
			}
		}
	}

	if d[f] == math.MaxInt {
		return math.MaxInt, []int{}
	}

	path := []int{f}
	for v := p[f]; v != -1; v = p[v] {
		path = append(path, v)
	}
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}
	return d[f], path
}

type WeightedEdge struct {
	v, w int
}

type Item struct {
	value, priority, index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
