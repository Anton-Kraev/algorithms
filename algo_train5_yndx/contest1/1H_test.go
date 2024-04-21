package main

import (
	"strconv"
	"testing"
)

func Test1H(t *testing.T) {
	var testCases = []struct {
		length, x1, v1, x2, v2 int64
		expected               string
	}{
		{6, 3, 1, 1, 1, "1.0000000000"},
		{12, 8, 10, 5, 20, "0.3000000000"},
		{5, 0, 0, 1, 2, "2.0000000000"},
		{10, 7, -3, 1, 4, "0.8571428571"},
		{762899414, 556082848, -539099316, 556082848, -582799403, "0.0000000000"},
		{615143346, 79387687, -80123649, 306422480, -80123649, "2.4075923389"},
		{72, 20, -38121735, 66, 288888467, "0.0000000795"},
		{55444931, 17419156, 0, 53245822, -398046024, "0.0382369025"},
		{956390104, 549514100, 7, 315097830, -7, "51569559.5714285746"},
		{5, 4, 0, 2, 0, "+Inf"},
	}

	for _, tc := range testCases {
		_, result := Solve1H(tc.length, tc.x1, tc.v1, tc.x2, tc.v2)
		resultStr := strconv.FormatFloat(result, 'f', 10, 64)
		if resultStr != tc.expected {
			t.Errorf("Solve1H(%v, %v, %v, %v, %v): expected %s, got %s",
				tc.length, tc.x1, tc.v1, tc.x2, tc.v2, tc.expected, resultStr)
		}
		// else {
		//	fmt.Printf("--- OK:   Solve1A(%v, %v, %v, %v)\n",
		//		tc.vStart, tc.mStart, tc.vDist, tc.mDist)
		//}
	}
}
