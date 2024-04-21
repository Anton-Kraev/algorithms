package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Date struct {
	day   int
	month string
}

type Month struct {
	name      string
	daysCount int
}

func Main1I() (string, string) {
	var n, year int
	var firstDay string
	fmt.Scanln(&n)
	fmt.Scanln(&year)
	var weekends []Date
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		input := strings.Split(scanner.Text(), " ")
		day, _ := strconv.Atoi(input[0])
		weekend := Date{day, input[1]}
		weekends = append(weekends, weekend)
	}
	fmt.Scan(&firstDay)
	return Solve1I(n, year, weekends, firstDay)
}

func Solve1I(n, year int, weekends []Date, firstDay string) (string, string) {
	var months = []Month{
		{"January", 31},
		{"February", 28},
		{"March", 31},
		{"April", 30},
		{"May", 31},
		{"June", 30},
		{"July", 31},
		{"August", 31},
		{"September", 30},
		{"October", 31},
		{"November", 30},
		{"December", 31},
	}
	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	var weekendsCount = map[string]int{
		"Monday":    52,
		"Tuesday":   52,
		"Wednesday": 52,
		"Thursday":  52,
		"Friday":    52,
		"Saturday":  52,
		"Sunday":    52,
	}
	weekendsCount[firstDay]++

	isLeap := year%400 == 0 || year%4 == 0 && year%100 != 0
	if isLeap {
		for i, day := range daysOfWeek {
			if day == firstDay {
				nextDay := daysOfWeek[(i+1)%7]
				weekendsCount[nextDay]++
			}
		}
	}

	for _, weekend := range weekends {
		dayNumber := 0
		for _, month := range months {
			if month.name == weekend.month {
				dayNumber += weekend.day
				break
			}
			dayNumber += month.daysCount
			if month.name == "February" && isLeap {
				dayNumber++
			}
		}
		for i, day := range daysOfWeek {
			if day == firstDay {
				dayNumber += i - 1
			}
		}
		weekendDay := daysOfWeek[dayNumber%7]
		for day := range weekendsCount {
			if day != weekendDay {
				weekendsCount[day]++
			}
		}
	}

	max, min := -1, math.MaxInt
	maxDay, minDay := "", ""
	for day, cnt := range weekendsCount {
		if cnt < min {
			min = cnt
			minDay = day
		}
		if cnt > max {
			max = cnt
			maxDay = day
		}
	}

	return maxDay, minDay
}
