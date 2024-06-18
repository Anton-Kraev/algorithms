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

	var results []int
	for {
		if s, _ := reader.ReadString('\n'); strings.TrimSpace(s) == "-1" {
			break
		}

		qCntStr, _ := reader.ReadString('\n')
		qCnt, _ := strconv.Atoi(strings.TrimSpace(qCntStr))

		lastGood := 0
		pdsu := newParity()
		for i := 0; i < qCnt; i++ {
			question, _ := reader.ReadString('\n')
			var (
				ok   bool
				l, r int
				t    string
			)
			fmt.Sscanf(strings.TrimSpace(question), "%d %d %s", &l, &r, &t)

			if t == "even" {
				ok = pdsu.add(l, r, false)
			} else {
				ok = pdsu.add(l, r, true)
			}
			if !ok {
				break
			}
			lastGood++
		}
		for i := 0; i < qCnt-lastGood-1; i++ {
			reader.ReadString('\n')
		}

		results = append(results, lastGood)
	}

	for _, r := range results {
		fmt.Println(r)
	}
}

func newParity() *Parity {
	return &Parity{make(map[int]bool), make(map[int]bool), make(map[int]int)}
}

type Parity struct {
	exist, odd map[int]bool
	prev       map[int]int
}

func (p *Parity) add(a int, b int, c bool) bool {
	if !p.exist[b] {
		p.exist[b] = true
		p.odd[b] = c
		p.prev[b] = a
		return true
	}
	i := p.prev[b]
	if i == a {
		return p.odd[b] == c
	} else if i < a {
		return p.add(i, a-1, c != p.odd[b])
	}
	return p.add(a, i-1, c != p.odd[b])
}
