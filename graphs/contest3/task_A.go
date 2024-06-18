package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const inf = math.MaxInt

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstStr, _ := reader.ReadString('\n')
	var vCnt, eCnt, townSequenceLen int
	fmt.Sscanf(strings.TrimSpace(firstStr), "%d %d %d", &vCnt, &eCnt, &townSequenceLen)

	distances := make([][]int, vCnt, vCnt)
	flightNumbers := make([][]int, vCnt, vCnt)
	for i := 0; i < vCnt; i++ {
		distances[i] = make([]int, vCnt, vCnt)
		flightNumbers[i] = make([]int, vCnt, vCnt)
	}
	for k := 0; k < eCnt; k++ {
		edge, _ := reader.ReadString('\n')
		var v, u, w int
		fmt.Sscanf(strings.TrimSpace(edge), "%d %d %d", &v, &u, &w)
		distances[v-1][u-1] = w
		flightNumbers[v-1][u-1] = k + 1
	}
	for i := 0; i < vCnt; i++ {
		for j := 0; j < vCnt; j++ {
			if i != j && distances[i][j] == 0 {
				distances[i][j] = -inf
			}
		}
	}

	var townSequence []int
	lastStr, _ := reader.ReadString('\n')
	for _, townStr := range strings.Split(strings.TrimSpace(lastStr), " ") {
		town, _ := strconv.Atoi(townStr)
		townSequence = append(townSequence, town-1)
	}

	next := floyd(distances)
	findCycles(distances)
	var flightSequence []int
	infinitePath := false
	for i := 0; i < townSequenceLen-1; i++ {
		path := getPath(townSequence[i], townSequence[i+1], distances, next)
		if len(path) == 0 {
			fmt.Println("no path")
			infinitePath = true
			break
		} else if len(path) == 1 {
			fmt.Println("infinitely kind")
			infinitePath = true
			break
		}
		for k := 0; k < len(path)-1; k++ {
			flightSequence = append(flightSequence, flightNumbers[path[k]][path[k+1]])
		}
	}
	if !infinitePath {
		fmt.Println(len(flightSequence))
		for _, flight := range flightSequence {
			fmt.Print(flight, " ")
		}
	}
}

func floyd(d [][]int) [][]int {
	vCnt := len(d)
	next := make([][]int, vCnt, vCnt)
	for i := 0; i < vCnt; i++ {
		next[i] = make([]int, vCnt, vCnt)
		for j := 0; j < vCnt; j++ {
			next[i][j] = j
		}
	}
	for k := 0; k < vCnt; k++ {
		for i := 0; i < vCnt; i++ {
			for j := 0; j < vCnt; j++ {
				if d[i][k] > -inf && d[k][j] > -inf && d[i][k]+d[k][j] > d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
					next[i][j] = next[i][k]
				}
			}
		}
	}
	return next
}

func findCycles(d [][]int) {
	vCnt := len(d)
	for i := 0; i < vCnt; i++ {
		for j := 0; j < vCnt; j++ {
			for t := 0; t < vCnt; t++ {
				if d[i][t] > -inf && d[t][t] > 0 && d[t][j] > -inf {
					d[i][j] = inf
				}
			}
		}
	}
}

func getPath(u, v int, d, next [][]int) []int {
	if d[u][v] == -inf {
		return []int{}
	} else if d[u][v] == inf {
		return []int{0}
	}

	var path []int
	for c := u; c != v; c = next[c][v] {
		path = append(path, c)
	}
	return append(path, v)
}
