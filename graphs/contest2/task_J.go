package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var vCntLeft, vCntRight int
	firstLine, _ := reader.ReadString('\n')
	fmt.Sscanf(strings.TrimSpace(firstLine), "%d %d", &vCntLeft, &vCntRight)

	edgesFromLeft := make([][]int, vCntLeft, vCntLeft)
	for vLeft := 0; vLeft < vCntLeft; vLeft++ {
		line, _ := reader.ReadString('\n')
		for _, vRightStr := range strings.Split(strings.TrimSpace(line), " ")[1:] {
			vRight, _ := strconv.Atoi(vRightStr)
			edgesFromLeft[vLeft] = append(edgesFromLeft[vLeft], vRight-1)
		}
	}

	matching := make([]int, vCntLeft, vCntLeft)
	edgesFromRight := make([][]int, vCntRight, vCntRight)
	lastLine, _ := reader.ReadString('\n')
	for vLeft, vRightStr := range strings.Split(strings.TrimSpace(lastLine), " ") {
		vRight, _ := strconv.Atoi(vRightStr)
		matching[vLeft] = vRight - 1
		if vRight-1 >= 0 {
			edgesFromRight[vRight-1] = append(edgesFromRight[vRight-1], vLeft)
			for i, v := range edgesFromLeft[vLeft] {
				if v == vRight-1 {
					edgesFromLeft[vLeft][i] = edgesFromLeft[vLeft][len(edgesFromLeft[vLeft])-1]
					edgesFromLeft[vLeft] = edgesFromLeft[vLeft][:len(edgesFromLeft[vLeft])-1]
					break
				}
			}
		}
	}

	visitedLeft, visitedRight := make([]bool, vCntLeft, vCntLeft), make([]bool, vCntRight, vCntRight)
	for v := 0; v < vCntLeft; v++ {
		if matching[v] == -1 && !visitedLeft[v] {
			dfs(v, true, visitedLeft, visitedRight, edgesFromLeft, edgesFromRight)
		}
	}

	var minLeft, minRight []interface{}
	for v, visited := range visitedLeft {
		if !visited {
			minLeft = append(minLeft, v+1)
		}
	}
	for v, visited := range visitedRight {
		if visited {
			minRight = append(minRight, v+1)
		}
	}
	fmt.Println(len(minLeft) + len(minRight))
	fmt.Print(len(minLeft), " ")
	fmt.Println(minLeft...)
	fmt.Print(len(minRight), " ")
	fmt.Println(minRight...)
}

func dfs(v int, vleft bool, visitedL, visitedR []bool, edgesL, edgesR [][]int) {
	if vleft {
		visitedL[v] = true
		for _, u := range edgesL[v] {
			if !visitedR[u] {
				dfs(u, !vleft, visitedL, visitedR, edgesL, edgesR)
			}
		}
	} else {
		visitedR[v] = true
		for _, u := range edgesR[v] {
			if !visitedL[u] {
				dfs(u, !vleft, visitedL, visitedR, edgesL, edgesR)
			}
		}
	}
}
