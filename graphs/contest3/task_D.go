package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const maxKnowledge = math.MaxInt32

type Door struct {
	from, to, know int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	var roomsCnt, doorsCnt int
	fmt.Sscanf(strings.TrimSpace(firstLine), "%d %d", &roomsCnt, &doorsCnt)

	var doors []Door
	for i := 0; i < doorsCnt; i++ {
		door, _ := reader.ReadString('\n')
		var room1, room2, knowledge int
		fmt.Sscanf(strings.TrimSpace(door), "%d %d %d", &room1, &room2, &knowledge)
		doors = append(doors, Door{room1 - 1, room2 - 1, knowledge})
	}

	max := BellmanFord(roomsCnt, doors)
	if max == -maxKnowledge {
		fmt.Println(":(")
	} else if max == maxKnowledge {
		fmt.Println(":)")
	} else {
		fmt.Println(max)
	}
}

func BellmanFord(vCnt int, edges []Door) int {
	d := make([]int64, vCnt, vCnt)
	for i := 0; i < vCnt; i++ {
		d[i] = -maxKnowledge
	}
	d[0] = 0

	for i := 0; i < vCnt-1; i++ {
		for _, e := range edges {
			if d[e.from] > -maxKnowledge && d[e.to] < d[e.from]+int64(e.know) {
				d[e.to] = d[e.from] + int64(e.know)
				if d[e.to] > maxKnowledge {
					d[e.to] = maxKnowledge
				}
			}
		}
	}

	used := make([]bool, vCnt, vCnt)
	for _, e := range edges {
		if d[e.from] > -maxKnowledge && d[e.to] < d[e.from]+int64(e.know) {
			if dfs(e.to, used, edges); used[vCnt-1] {
				return maxKnowledge
			}
			break
		}
	}

	return int(d[vCnt-1])
}

func dfs(v int, used []bool, edges []Door) {
	used[v] = true
	for _, edge := range edges {
		if edge.from == v && !used[edge.to] {
			dfs(edge.to, used, edges)
		}
	}
}
