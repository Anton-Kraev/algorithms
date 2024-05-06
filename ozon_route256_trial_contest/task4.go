package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type RunResult struct {
	runner, time int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(in, &n)
		timesStr, _ := in.ReadString('\n')
		timesStrArr := strings.Split(strings.TrimSpace(timesStr), " ")
		results := make([]RunResult, n)
		for j := 0; j < n; j++ {
			time, _ := strconv.Atoi(timesStrArr[j])
			results[j] = RunResult{j, time}
		}
		for _, place := range solve4(results) {
			fmt.Fprintf(out, "%d ", place)
		}
		fmt.Fprintln(out)
	}
}

func solve4(runs []RunResult) []int {
	results := make([]int, len(runs))
	sort.Slice(runs, func(i, j int) bool {
		return runs[i].time < runs[j].time
	})
	results[runs[0].runner] = 1
	for i := 1; i < len(runs); i++ {
		if runs[i].time-runs[i-1].time <= 1 {
			results[runs[i].runner] = results[runs[i-1].runner]
		} else {
			results[runs[i].runner] = i + 1
		}
	}
	return results
}
