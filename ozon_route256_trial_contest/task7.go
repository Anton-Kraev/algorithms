package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	patient, window int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		var windowsCnt, patientsCnt int
		fmt.Fscanln(in, &windowsCnt, &patientsCnt)
		windowsStr, _ := in.ReadString('\n')
		windowsStrArr := strings.Split(strings.TrimSpace(windowsStr), " ")
		entries := make([]Entry, patientsCnt)
		for j := 0; j < patientsCnt; j++ {
			window, _ := strconv.Atoi(windowsStrArr[j])
			entries[j] = Entry{j, window}
		}
		changes, status := solve7(windowsCnt, entries)
		if !status {
			fmt.Fprint(out, "x")
		} else {
			for _, change := range changes {
				if change == 1 {
					fmt.Fprint(out, "+")
				} else if change == -1 {
					fmt.Fprint(out, "-")
				} else if change == 0 {
					fmt.Fprint(out, "0")
				}
			}
		}
		fmt.Fprintln(out)
	}
}

func solve7(maxWindow int, entries []Entry) ([]int, bool) {
	changes := make([]int, len(entries))
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].window < entries[j].window
	})
	if entries[0].window > 1 {
		changes[entries[0].patient] = -1
	}
	for i := 1; i < len(entries); i++ {
		windowCurr := entries[i].window
		windowPrev := entries[i-1].window + changes[entries[i-1].patient]
		if windowCurr-1 > windowPrev && windowCurr > 1 {
			changes[entries[i].patient] = -1
		} else if windowCurr == windowPrev && windowCurr < maxWindow {
			changes[entries[i].patient] = 1
		} else if windowCurr == windowPrev && windowCurr >= maxWindow || windowCurr < windowPrev {
			return []int{}, false
		}
	}
	return changes, true
}
