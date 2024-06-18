package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Exchanger struct {
	from, to   int
	rate, comm float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	var (
		currenciesCnt, exchangersCnt, startCurrency int
		startAmount                                 float64
	)
	fmt.Sscanf(
		strings.TrimSpace(firstLine),
		"%d %d %d %f",
		&currenciesCnt, &exchangersCnt, &startCurrency, &startAmount,
	)

	var exchangers []Exchanger
	for i := 0; i < exchangersCnt; i++ {
		exchanger, _ := reader.ReadString('\n')
		var (
			a, b               int
			rab, cab, rba, cba float64
		)
		fmt.Sscanf(strings.TrimSpace(exchanger), "%d %d %f %f %f %f", &a, &b, &rab, &cab, &rba, &cba)
		exchangers = append(exchangers, Exchanger{a - 1, b - 1, rab, cab})
		exchangers = append(exchangers, Exchanger{b - 1, a - 1, rba, cba})
	}

	if canIncreaseAmount(currenciesCnt, startCurrency-1, startAmount, exchangers) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func canIncreaseAmount(currenciesCnt, startCurrency int, startAmount float64, exchangers []Exchanger) bool {
	d := make([]float64, currenciesCnt, currenciesCnt)
	for i := 0; i < currenciesCnt; i++ {
		d[i] = math.Inf(-1)
	}
	d[startCurrency] = startAmount

	for i := 0; i < currenciesCnt-1; i++ {
		for _, e := range exchangers {
			if d[e.from] > math.Inf(-1) && (d[e.from]-e.comm)*e.rate > math.Max(0, d[e.to]) {
				d[e.to] = (d[e.from] - e.comm) * e.rate
			}
		}
	}

	used := make([]bool, currenciesCnt, currenciesCnt)
	for _, e := range exchangers {
		if d[e.from] > math.Inf(-1) && (d[e.from]-e.comm)*e.rate > math.Max(0, d[e.to]) {
			findPath(e.to, used, exchangers)
			return used[startCurrency]
		}
	}
	return false
}

func findPath(v int, used []bool, edges []Exchanger) {
	used[v] = true
	for _, edge := range edges {
		if edge.from == v && !used[edge.to] {
			findPath(edge.to, used, edges)
		}
	}
}
