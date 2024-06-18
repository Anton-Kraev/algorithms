package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func valence(atom rune) int {
	switch atom {
	case 'H':
		return 1
	case 'O':
		return 2
	case 'N':
		return 3
	case 'C':
		return 4
	default:
		return 0
	}
}

type Edge struct {
	to, cap int
}

type Graph struct {
	h, w, vCnt, fromSource, toSink, maxFlow int
	edges                                   [][]Edge
	visited                                 []bool
}

func NewGraph(h, w int) *Graph {
	vCnt := h*w + 2
	return &Graph{h: h, w: w, vCnt: vCnt, edges: make([][]Edge, vCnt, vCnt), visited: make([]bool, vCnt, vCnt)}
}

func (g *Graph) InitWithValences(paper [][]rune) *Graph {
	for r := 0; r < g.h; r++ {
		for c := 0; c < g.w; c++ {
			val := valence(paper[r][c])
			if c%2 == r%2 {
				g.addEdge(0, g.getIdx(r, c), val)
				g.fromSource += val
				if r+1 < g.h {
					g.addEdge(g.getIdx(r, c), g.getIdx(r+1, c), 1)
				}
				if c+1 < g.w {
					g.addEdge(g.getIdx(r, c), g.getIdx(r, c+1), 1)
				}
				if r-1 >= 0 {
					g.addEdge(g.getIdx(r, c), g.getIdx(r-1, c), 1)
				}
				if c-1 >= 0 {
					g.addEdge(g.getIdx(r, c), g.getIdx(r, c-1), 1)
				}
			} else {
				g.addEdge(g.getIdx(r, c), g.vCnt-1, val)
				g.toSink += val
			}
		}
	}
	return g
}

func (g *Graph) getIdx(r, c int) int {
	return r*g.w + c + 1
}

func (g *Graph) addEdge(from, to, capacity int) {
	g.edges[from] = append(g.edges[from], Edge{to, capacity})
}

func (g *Graph) PushFlow(v, flow int) int {
	g.visited[v] = true
	if v == g.vCnt-1 || flow == 0 {
		g.maxFlow += flow
		return flow
	}

	for i, edge := range g.edges[v] {
		if g.visited[edge.to] || edge.cap == 0 {
			continue
		}

		minCapacity := g.PushFlow(edge.to, int(math.Min(float64(flow), float64(edge.cap))))
		g.edges[v][i].cap -= minCapacity

		if minCapacity > 0 {
			backEdge := -1
			for child := 0; child < len(g.edges[edge.to]); child++ {
				if g.edges[edge.to][child].to == v {
					backEdge = child
					break
				}
			}
			if backEdge != -1 {
				g.edges[edge.to][backEdge].cap += minCapacity
			} else {
				g.edges[edge.to] = append(g.edges[edge.to], Edge{v, minCapacity})
			}
			return minCapacity
		}
	}

	return 0
}

func main() {
	var h, w int
	fmt.Scanln(&h, &w)
	var paper [][]rune
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < h && scanner.Scan(); i++ {
		paper = append(paper, []rune(scanner.Text()))
	}

	g := NewGraph(h, w).InitWithValences(paper)
	for g.PushFlow(0, 5) > 0 {
		for i := 0; i < g.vCnt; i++ {
			g.visited[i] = false
		}
	}

	if g.fromSource == g.maxFlow && g.toSink == g.maxFlow && g.maxFlow > 0 {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
