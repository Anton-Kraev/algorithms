package main

import (
	"testing"
)

func Test1G(t *testing.T) {
	var testCases = []struct {
		myCnt, hp, enemyNew int
		expected            int
	}{
		{10, 11, 15, 4},
		{1, 2, 1, -1},
		{1, 1, 1, 1},
		{25, 200, 10, 13},
		{250, 500, 187, 4},
		{250, 500, 218, 6},
		{250, 500, 226, 8},
	}

	for _, tc := range testCases {
		result := Solve1G(tc.myCnt, tc.hp, tc.enemyNew)
		if result != tc.expected {
			t.Errorf("Solve1G(%v, %v, %v): expected %v, got %v",
				tc.myCnt, tc.hp, tc.enemyNew, tc.expected, result)
		}
		// else {
		//	fmt.Printf("--- OK:   Solve1A(%v, %v, %v, %v)\n",
		//		tc.vStart, tc.mStart, tc.vDist, tc.mDist)
		//}
	}
}
