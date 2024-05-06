package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		var (
			trucksCnt, truckCapacity, boxesCnt int
			truckLoad, boxes                   []int
		)

		fmt.Fscanln(in, &trucksCnt, &truckCapacity)
		truckLoad = make([]int, trucksCnt)

		fmt.Fscanln(in, &boxesCnt)
		boxesStr, _ := in.ReadString('\n')
		for _, boxStr := range strings.Split(strings.TrimSpace(boxesStr), " ") {
			box, _ := strconv.Atoi(boxStr)
			boxes = append(boxes, box)
		}

		sort.Slice(boxes, func(i, j int) bool {
			return boxes[i] > boxes[j]
		})

		cargoCnt := 1
		truckPtr := 0
		for _, box := range boxes {
			initTruckPtr := truckPtr
			if truckCapacity-truckLoad[truckPtr] >= (1 << box) {
				truckLoad[truckPtr] += 1 << box
				continue
			}
			if truckPtr == trucksCnt-1 {
				truckPtr = 0
			} else {
				truckPtr++
			}
			for truckCapacity-truckLoad[truckPtr] < (1<<box) && truckPtr != initTruckPtr {
				if truckPtr >= trucksCnt-1 {
					truckPtr = 0
				} else {
					truckPtr++
				}
			}
			if truckCapacity-truckLoad[truckPtr] >= (1 << box) {
				truckLoad[truckPtr] += 1 << box
			} else {
				truckPtr = 0
				truckLoad = make([]int, trucksCnt)
				truckLoad[0] += 1 << box
				cargoCnt++
			}
		}

		fmt.Fprintln(out, cargoCnt)
	}
}
