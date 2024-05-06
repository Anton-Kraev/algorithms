package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscanln(in, &n, &q)
	msgCnt := 0
	lastGlobal := 0
	messages := make([]int, n+1)
	for i := 0; i < q; i++ {
		var qType, user int
		fmt.Fscanln(in, &qType, &user)
		if qType == 1 {
			msgCnt++
			if user == 0 {
				lastGlobal = msgCnt
			} else {
				messages[user] = msgCnt
			}
		} else if qType == 2 {
			fmt.Fprintln(out, math.Max(float64(lastGlobal), float64(messages[user])))
		}
	}
}
