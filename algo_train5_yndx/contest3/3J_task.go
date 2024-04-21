package main

import (
	"fmt"
	"math"
)

type Request struct {
	device, part int
}

func main() {
	var deviceCnt, partsCnt int
	fmt.Scanln(&deviceCnt, &partsCnt)
	deviceParts := make([]map[int]struct{}, 0, deviceCnt)
	partDevices := make([]map[int]struct{}, 0, partsCnt)
	values := make([][]int, 0, deviceCnt)
	stepsCount := make([]int, 0, deviceCnt)
	for i := 0; i < deviceCnt; i++ {
		deviceParts = append(deviceParts, make(map[int]struct{}))
		values = append(values, make([]int, deviceCnt, deviceCnt))
		for j := 0; j < deviceCnt; j++ {
			values[i][j] = 0
		}
		stepsCount = append(stepsCount, math.MaxInt)
	}
	for i := 0; i < partsCnt; i++ {
		partDevices = append(partDevices, make(map[int]struct{}))
		partDevices[i][0] = struct{}{}
		deviceParts[0][i] = struct{}{}
	}

	currStep := 0
	for {
		requests := make(map[int][]Request)
		lastStep := true
		for i := 0; i < deviceCnt; i++ {
			part := choosePart(i, partDevices)
			if part == -1 || len(deviceParts[i]) == partsCnt {
				stepsCount[i] = int(math.Min(float64(stepsCount[i]), float64(currStep)))
				continue
			}
			lastStep = false
			device := chooseDevice(part, partDevices, deviceParts)
			if device == -1 {
				panic("device not found")
			}

			_, hasReqs := requests[device]
			if !hasReqs {
				requests[device] = make([]Request, 0)
			}
			requests[device] = append(requests[device], Request{i, part})
		}

		if lastStep {
			break
		}

		allowed := make(map[int]Request)
		for device, req := range requests {
			allowedRequest := pickMostValuable(device, deviceParts, values, req)
			allowed[device] = allowedRequest
		}
		for device, req := range allowed {
			deviceParts[req.device][req.part] = struct{}{}
			partDevices[req.part][req.device] = struct{}{}
			values[req.device][device]++
		}

		currStep++
	}

	for _, s := range stepsCount[1:] {
		fmt.Print(s, " ")
	}
}

func choosePart(device int, partDevices []map[int]struct{}) int {
	min, idx := math.MaxInt, -1
	for part, devices := range partDevices {
		_, hasPart := devices[device]
		if !hasPart && len(devices) < min {
			min = len(devices)
			idx = part
		} else if !hasPart && len(devices) == min && part < idx {
			idx = part
		}
	}
	return idx
}

func chooseDevice(part int, partDevices, deviceParts []map[int]struct{}) int {
	min, idx := math.MaxInt, -1
	for device := range partDevices[part] {
		if len(deviceParts[device]) < min {
			min = len(deviceParts[device])
			idx = device
		} else if len(deviceParts[device]) == min && device < idx {
			idx = device
		}
	}
	return idx
}

func pickMostValuable(device int, deviceParts []map[int]struct{}, values [][]int, requests []Request) Request {
	mostValuable := requests[0]
	maxValue, minParts, idx := values[device][requests[0].device], len(deviceParts[requests[0].device]), requests[0].device
	for _, req := range requests[1:] {
		currIdx := req.device
		currValue := values[device][currIdx]
		currParts := len(deviceParts[currIdx])
		if currValue > maxValue {
			mostValuable = req
			maxValue, minParts, idx = currValue, currParts, currIdx
		} else if currValue == maxValue && currParts < minParts {
			mostValuable = req
			minParts, idx = currParts, currIdx
		} else if currValue == maxValue && currParts == minParts && currIdx < idx {
			mostValuable = req
			idx = currIdx
		}
	}
	return mostValuable
}
