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

	str1, _ := reader.ReadString('\n')
	var playerCnt, queriesCnt int
	fmt.Sscanf(strings.TrimSpace(str1), "%d %d", &playerCnt, &queriesCnt)

	clans := newClans(playerCnt)
	var results []string
	for i := 0; i < queriesCnt; i++ {
		query, _ := reader.ReadString('\n')
		var (
			op     string
			n1, n2 int
		)
		fmt.Sscanf(strings.TrimSpace(query), "%s %d %d", &op, &n1, &n2)
		n1--

		if op == "get" {
			results = append(results, strconv.FormatInt(int64(clans.get(n1)), 10))
		} else if op == "join" {
			clans.join(n1, n2-1)
		} else if op == "add" {
			clans.add(n1, n2)
		}
	}

	fmt.Print(strings.Join(results, "\n"))
}

type Clans struct {
	p, xp []int
}

func newClans(size int) *Clans {
	xp := make([]int, size, size)
	p := make([]int, size, size)
	for i := 0; i < size; i++ {
		p[i] = i
	}
	return &Clans{p, xp}
}

func (clans *Clans) find(n int) (int, int) {
	if clans.p[n] == n {
		return n, clans.xp[n]
	}
	var clanXp int
	clans.p[n], clanXp = clans.find(clans.p[n])
	clans.xp[n] += clanXp - clans.xp[clans.p[n]]
	return clans.p[n], clans.xp[n] + clans.xp[clans.p[n]]
}

func (clans *Clans) join(n, m int) {
	pN, _ := clans.find(n)
	pM, _ := clans.find(m)
	if pN == pM {
		return
	}
	if clans.xp[pN] < clans.xp[pM] {
		pN, pM = pM, pN
	}
	clans.xp[pN] -= clans.xp[pM]
	clans.p[pN] = pM
}

func (clans *Clans) add(n, xp int) {
	p, _ := clans.find(n)
	clans.xp[p] += xp
}

func (clans *Clans) get(n int) int {
	_, xp := clans.find(n)
	return xp
}
