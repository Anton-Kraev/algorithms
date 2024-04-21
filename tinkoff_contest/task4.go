package main

import (
	"fmt"
)

type cell struct {
	r, c int
}

func main() {
	var (
		n         int
		direction string
	)
	fmt.Scanln(&n, &direction)

	fmt.Println(n * n / 4 * 3)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			toRotate := [4]cell{{i, j}, {j, n - i - 1}, {n - j - 1, i}, {n - i - 1, n - j - 1}}

			fmt.Println(toRotate[0].r, toRotate[0].c, toRotate[1].r, toRotate[1].c)
			if direction == "L" {
				fmt.Println(toRotate[1].r, toRotate[1].c, toRotate[3].r, toRotate[3].c)
			} else if direction == "R" {
				fmt.Println(toRotate[0].r, toRotate[0].c, toRotate[2].r, toRotate[2].c)
			}
			fmt.Println(toRotate[2].r, toRotate[2].c, toRotate[3].r, toRotate[3].c)
		}
	}
}
