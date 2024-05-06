package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Truck struct {
	id, start, end, capacity int
}

type Arrival struct {
	id, time int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		var (
			arrivalsCnt, trucksCnt int
			arrivals               []Arrival
			trucks                 []Truck
		)
		fmt.Fscanln(in, &arrivalsCnt)
		arrivalsStr, _ := in.ReadString('\n')
		for j, arrivalStr := range strings.Split(strings.TrimSpace(arrivalsStr), " ") {
			arrival, _ := strconv.Atoi(arrivalStr)
			arrivals = append(arrivals, Arrival{j, arrival})
		}
		fmt.Fscanln(in, &trucksCnt)
		for j := 1; j < trucksCnt+1; j++ {
			var start, end, capacity int
			fmt.Fscanf(in, "%d %d %d", &start, &end, &capacity)
			fmt.Fscanln(in)
			trucks = append(trucks, Truck{j, start, end, capacity})
		}
		sort.Slice(arrivals, func(i, j int) bool {
			return arrivals[i].time < arrivals[j].time
		})
		sort.Slice(trucks, func(i, j int) bool {
			if trucks[i].start == trucks[j].start {
				return trucks[i].id < trucks[j].id
			}
			return trucks[i].start < trucks[j].start
		})
		truckNumbers := make([]int, arrivalsCnt)
		for _, arrival := range arrivals {
			hasGoodTruck := false
			for j, truck := range trucks {
				if truck.start <= arrival.time && truck.end >= arrival.time && truck.capacity > 0 {
					hasGoodTruck = true
					trucks[j].capacity--
					truckNumbers[arrival.id] = truck.id
					break
				}
			}
			if !hasGoodTruck {
				truckNumbers[arrival.id] = -1
			}
		}
		for _, number := range truckNumbers {
			fmt.Fprintf(out, "%d ", number)
		}
		fmt.Fprintln(out)
	}
}
