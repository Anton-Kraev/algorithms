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
	var usersCnt, actionsCnt int
	fmt.Sscanf(strings.TrimSpace(str1), "%d %d", &usersCnt, &actionsCnt)

	chats := newChats(usersCnt)
	var results []string
	for k := 0; k < actionsCnt; k++ {
		action, _ := reader.ReadString('\n')
		var act, i, j int64
		fmt.Sscanf(strings.TrimSpace(action), "%d %d %d", &act, &i, &j)

		if act == 1 {
			chats.sendMsg(i)
		} else if act == 2 {
			chats.merge(i, j)
		} else if act == 3 {
			results = append(results, strconv.FormatInt(chats.readMsg(i), 10))
		}
	}

	fmt.Print(strings.Join(results, "\n"))
}

const maxZerg = 1000003

type Chats struct {
	size, zerg           int64
	p, msgAll, msgReaded []int64
}

func newChats(size int) *Chats {
	msgAll := make([]int64, size, size)
	msgReaded := make([]int64, size, size)
	p := make([]int64, size, size)
	for i := 0; i < size; i++ {
		p[i] = int64(i)
	}
	return &Chats{int64(size), 0, p, msgAll, msgReaded}
}

func (chats *Chats) find(n int64) (int64, int64) {
	if chats.p[n] == n {
		return n, chats.msgAll[n]
	}
	var chatMsg int64
	chats.p[n], chatMsg = chats.find(chats.p[n])
	chats.msgAll[n] += chatMsg - chats.msgAll[chats.p[n]]
	return chats.p[n], chats.msgAll[n] + chats.msgAll[chats.p[n]]
}

func (chats *Chats) merge(n, m int64) {
	pN, _ := chats.find((n + chats.zerg) % chats.size)
	pM, _ := chats.find((m + chats.zerg) % chats.size)
	if pN == pM {
		return
	}
	chats.zerg = (13*chats.zerg + 11) % maxZerg
	if chats.msgAll[pN] < chats.msgAll[pM] {
		pN, pM = pM, pN
	}
	chats.msgAll[pN] -= chats.msgAll[pM]
	chats.p[pN] = pM
}

func (chats *Chats) sendMsg(n int64) {
	p, _ := chats.find((n + chats.zerg) % chats.size)
	chats.msgAll[p] += 1
	chats.zerg = (30*chats.zerg + 239) % maxZerg
}

func (chats *Chats) readMsg(n int64) int64 {
	_, msg := chats.find((n + chats.zerg) % chats.size)
	newMsg := msg - chats.msgReaded[(n+chats.zerg)%chats.size]
	chats.msgReaded[(n+chats.zerg)%chats.size] += newMsg
	chats.zerg = (100500*chats.zerg + newMsg) % maxZerg
	return newMsg
}
