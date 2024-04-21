package main

import (
	"fmt"
	"math"
)

func main() {
	var length, x1, v1, x2, v2 int64
	fmt.Scan(&length, &x1, &v1, &x2, &v2)
	fmt.Println(Solve1H(length, x1, v1, x2, v2))
}

func Solve1H(length, x1, v1, x2, v2 int64) (string, float64) {
	if math.Abs(float64(x1)) == math.Abs(float64(x2)) {
		return "YES", 0
	}
	if v1 == 0 && v2 == 0 {
		return "NO", math.Inf(0)
	}

	var t1 float64
	if v1+v2 >= 0 {
		if x1+x2 >= length {
			t1 = float64(2*length-x2-x1) / float64(v1+v2)
		} else {
			t1 = float64(length-x2-x1) / float64(v1+v2)
		}
	} else {
		if x1+x2 > length {
			t1 = float64(length-x1-x2) / float64(v1+v2)
		} else {
			t1 = float64(-x1-x2) / float64(v1+v2)
		}
	}

	var t2 float64
	if (v1-v2)*(x2-x1) < 0 {
		if x2-x1 > 0 {
			t2 = float64(x2-x1-length) / float64(v1-v2)
		} else {
			t2 = float64(length+x2-x1) / float64(v1-v2)
		}
	} else {
		t2 = float64(x2-x1) / float64(v1-v2)
	}

	if math.Min(t1, t2) > 0 {
		return "YES", math.Min(t1, t2)
	}
	return "YES", math.Max(t1, t2)
}
