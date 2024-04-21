package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var start, final int
	board := make([]rune, n*n, n*n)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		line := scanner.Text()
		for j, el := range line {
			if el == 'S' {
				start = i*n + j
			} else if el == 'F' {
				final = i*n + j
			}
			board[i*n+j] = el
		}
	}

	res := solve6(start, final, n, board)
	fmt.Println(res)
}

func solve6(start, final, n int, board []rune) int {
	distKnight, distKing := make([]int, n*n, n*n), make([]int, n*n, n*n)
	for i := 0; i < n*n; i++ {
		distKnight[i], distKing[i] = n*n, n*n
	}
	distKnight[start] = 0

	queue := Queue{[]State{{start, Knight}}}
	for queue.isNotEmpty() {
		state := queue.pick()
		v, piece := state.v, state.piece

		var steps [8]int
		if piece == Knight {
			steps = [8]int{v - 2*n - 1, v - 2*n + 1, v - n - 2, v - n + 2, v + n - 2, v + n + 2, v + 2*n - 1, v + 2*n + 1}
		} else if piece == King {
			steps = [8]int{v - n - 1, v - n, v - n + 1, v - 1, v + 1, v + n - 1, v + n, v + n + 1}
		}

		for _, u := range steps {
			if u >= 0 && u < n*n && (distKnight[u] > distKnight[v]+1 && piece == Knight || distKing[u] > distKing[v]+1 && piece == King) {
				if piece == Knight {
					distKnight[u] = distKnight[v] + 1
				} else if piece == King {
					distKing[u] = distKing[v] + 1
				}

				if board[u] == 'K' {
					piece = Knight
				} else if board[u] == 'G' {
					piece = King
				}

				queue.push(State{u, piece})
			}
		}
	}

	minDist := int(math.Min(float64(distKnight[final]), float64(distKing[final])))
	if minDist == n*n {
		return -1
	}
	return minDist
}

type Piece int

const (
	Knight Piece = iota
	King
)

type State struct {
	v     int
	piece Piece
}

type Queue struct {
	elems []State
}

func (q *Queue) pick() State {
	el := q.elems[len(q.elems)-1]
	q.elems = q.elems[:len(q.elems)-1]
	return el
}

func (q *Queue) push(el State) {
	q.elems = append(q.elems, el)
}

func (q *Queue) isNotEmpty() bool {
	return len(q.elems) > 0
}
