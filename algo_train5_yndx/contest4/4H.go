package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var parties []Party
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		var votes, bribe int64
		fmt.Sscanf(scanner.Text(), "%d %d", &votes, &bribe)
		parties = append(parties, Party{i + 1, votes, bribe})
	}

	sort.Slice(parties, func(i, j int) bool {
		if parties[i].votes == parties[j].votes {
			if parties[i].bribe == -1 || parties[j].bribe == -1 {
				return parties[i].bribe > parties[j].bribe
			}
			return parties[i].bribe < parties[j].bribe
		}
		return parties[i].votes < parties[j].votes
	})
	var counfOfMax int64 = 0
	for i := 0; i < n; i++ {
		if parties[i].votes == parties[n-1].votes {
			counfOfMax++
		}
	}

	suffixSums := make([]int64, n+1, n+1)
	for i := n; i > 0; i-- {
		suffixSums[i-1] = suffixSums[i] + parties[i-1].votes
	}

	revotes := make([]int64, n, n)
	for i, party := range parties {
		if party.bribe != -1 {
			minAmount := binSearchMin(0, suffixSums[0], func(amount int64) bool {
				if amount == 0 {
					return n == 1 || i == n-1 && party.votes > parties[i-1].votes
				}
				firstMore := binSearchMin(0, int64(n-1), func(id int64) bool {
					return party.votes+amount <= parties[id].votes
				})
				if party.votes+amount > parties[firstMore].votes {
					return true
				}
				return (party.votes+amount-1)*(int64(n)-firstMore) >= suffixSums[firstMore]-amount
			})
			revotes[i] = minAmount
		}
	}

	idx := 0
	minId, minRevote, minBribe := parties[0].id, revotes[0], parties[0].bribe
	for i := 1; i < n; i++ {
		if parties[i].bribe != -1 && (revotes[i]+parties[i].bribe < minRevote+minBribe || minRevote+minBribe == -1) {
			minId = parties[i].id
			minRevote = revotes[i]
			minBribe = parties[i].bribe
			idx = i
		}
	}

	firstMore := n
	for i := idx + 1; i < n; i++ {
		if parties[i].votes >= parties[idx].votes+minRevote {
			firstMore = i
			break
		}
	}
	votes := minRevote
	parties[idx].votes += votes
	if firstMore != n {
		for i := firstMore; i < n; i++ {
			changes := parties[i].votes - parties[idx].votes + 1
			if changes > 0 {
				parties[i].votes -= changes
				votes -= changes
			}
		}
	}
	for i := n - 1; votes > 0; i-- {
		if i == idx {
			i = n - 1
		}
		parties[i].votes--
		votes--
	}

	sort.Slice(parties, func(i, j int) bool {
		return parties[i].id < parties[j].id
	})

	fmt.Println(minRevote + minBribe)
	fmt.Println(minId)
	for _, party := range parties {
		fmt.Print(party.votes, " ")
	}
}

type Party struct {
	id           int
	votes, bribe int64
}

func binSearchMin(low, high int64, check func(int64) bool) int64 {
	var mid int64

	for low < high {
		mid = (low + high) / 2
		if check(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
