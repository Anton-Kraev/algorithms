package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type Dequeue struct {
	head *Node
	tail *Node
	size int
}

func NewDequeue() *Dequeue {
	return &Dequeue{}
}

func (dq *Dequeue) IsEmpty() bool {
	return dq.size == 0
}

func (dq *Dequeue) PushFront(value int) {
	newNode := &Node{value: value}
	if dq.head == nil {
		dq.head = newNode
		dq.tail = newNode
	} else {
		newNode.next = dq.head
		dq.head.prev = newNode
		dq.head = newNode
	}
	dq.size++
}

func (dq *Dequeue) PushBack(value int) {
	newNode := &Node{value: value}
	if dq.tail == nil {
		dq.head = newNode
		dq.tail = newNode
	} else {
		newNode.prev = dq.tail
		dq.tail.next = newNode
		dq.tail = newNode
	}
	dq.size++
}

func (dq *Dequeue) PopFront() int {
	removedNode := dq.head
	dq.head = dq.head.next
	if dq.head == nil {
		dq.tail = nil
	} else {
		dq.head.prev = nil
	}
	dq.size--
	return removedNode.value
}

func (dq *Dequeue) PopBack() int {
	removedNode := dq.tail
	dq.tail = dq.tail.prev
	if dq.tail == nil {
		dq.head = nil
	} else {
		dq.tail.next = nil
	}
	dq.size--
	return removedNode.value
}

const infF = math.MaxInt

type CostFlowEdge struct {
	from, to, cap, cost, flow, reversed int
}

type MinCostMaxFlowSearch struct {
	vCnt  int
	edges [][]CostFlowEdge
	d, id []int
	p     []*CostFlowEdge
}

func InitMinCostMaxFlowSearch(vCnt int) *MinCostMaxFlowSearch {
	return &MinCostMaxFlowSearch{vCnt: vCnt, edges: make([][]CostFlowEdge, vCnt, vCnt)}
}

func (s *MinCostMaxFlowSearch) addEdge(from, to, cap, cost int) {
	r1, r2 := len(s.edges[to]), len(s.edges[from])
	s.edges[from] = append(s.edges[from], CostFlowEdge{from, to, cap, cost, 0, r1})
	s.edges[to] = append(s.edges[to], CostFlowEdge{to, from, 0, -cost, 0, r2})
}

func (s *MinCostMaxFlowSearch) Levit() int {
	minCost, maxFlow := 0, 0

	for {
		s.id = make([]int, s.vCnt, s.vCnt)
		s.p = make([]*CostFlowEdge, s.vCnt, s.vCnt)
		s.d = make([]int, s.vCnt, s.vCnt)
		for i := 0; i < s.vCnt; i++ {
			s.d[i] = infF
		}
		s.d[0] = 0
		dq := NewDequeue()
		dq.PushBack(0)

		for !dq.IsEmpty() {
			v := dq.PopFront()
			s.id[v] = 2
			for i, edge := range s.edges[v] {
				if edge.flow < edge.cap && s.d[edge.to] > s.d[edge.from]+edge.cost {
					s.d[edge.to] = s.d[edge.from] + edge.cost
					if s.id[edge.to] == 0 {
						dq.PushBack(edge.to)
					} else if s.id[edge.to] == 2 {
						dq.PushFront(edge.to)
					}
					s.id[edge.to] = 1
					s.p[edge.to] = &s.edges[v][i]
				}
			}
		}

		if s.d[s.vCnt-1] == infF {
			break
		}

		del := infF
		for v := s.vCnt - 1; v != 0; v = s.p[v].from {
			edge := s.p[v]
			if edge.cap-edge.flow < del {
				del = edge.cap - edge.flow
			}
		}
		for v := s.vCnt - 1; v != 0; v = s.p[v].from {
			edge := s.p[v]
			revEdge := &s.edges[edge.to][edge.reversed]
			edge.flow += del
			revEdge.flow -= del
			minCost += del * edge.cost
		}
		maxFlow += del
	}

	return minCost
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	var vCnt, eCnt int
	fmt.Sscanf(strings.TrimSpace(firstLine), "%d %d", &vCnt, &eCnt)

	s := InitMinCostMaxFlowSearch(vCnt)
	for i := 0; i < eCnt; i++ {
		edge, _ := reader.ReadString('\n')
		var from, to, capacity, cost int
		fmt.Sscanf(strings.TrimSpace(edge), "%d %d %d %d", &from, &to, &capacity, &cost)
		s.addEdge(from-1, to-1, capacity, cost)
	}

	fmt.Println(s.Levit())
}
