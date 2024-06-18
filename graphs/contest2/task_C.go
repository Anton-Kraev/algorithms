package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str1, _ := reader.ReadString('\n')
	var vCnt, eCnt int
	fmt.Sscanf(strings.TrimSpace(str1), "%d %d", &vCnt, &eCnt)

	maxBipartite := 0
	cdsu := newColoredDsu(vCnt)
	for i := 0; i < eCnt; i++ {
		e, _ := reader.ReadString('\n')
		var v1, v2 int
		fmt.Sscanf(strings.TrimSpace(e), "%d %d", &v1, &v2)
		isBipartite := cdsu.unite(v1-1, v2-1)
		if !isBipartite {
			break
		}
		maxBipartite++
	}

	fmt.Println(strings.Repeat("1", maxBipartite) + strings.Repeat("0", eCnt-maxBipartite))
}

type ParityDsu struct {
	parent, parity []int
}

func newColoredDsu(size int) *ParityDsu {
	p := make([]int, size, size)
	for i := 0; i < size; i++ {
		p[i] = i
	}
	return &ParityDsu{p, make([]int, size, size)}
}

func (dsu *ParityDsu) find(n int) (int, int) {
	if dsu.parent[n] != n {
		parity := dsu.parity[n]
		dsu.parent[n], dsu.parity[n] = dsu.find(dsu.parent[n])
		dsu.parity[n] ^= parity
	}
	return dsu.parent[n], dsu.parity[n]
}

func (dsu *ParityDsu) unite(n, m int) bool {
	n, pn := dsu.find(n)
	m, pm := dsu.find(m)
	if n == m {
		return pn != pm
	}
	if rand.Intn(2) == 0 {
		n, m = m, n
	}
	dsu.parent[n] = m
	dsu.parity[n] = pn ^ pm ^ 1
	return true
}
