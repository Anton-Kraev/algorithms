package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type Layout string

const (
	Embedded   Layout = "embedded"
	Surrounded Layout = "surrounded"
	Floating   Layout = "floating"
)

type Offset struct {
	dx, dy int
}

type Image struct {
	layout        Layout
	width, height int
	offset        *Offset
}

type NewParagraph any

type Word struct {
	length int
}

type Placed struct {
	start, end Point
	layout     string
}

func Main1J() []Point {
	var w, h, c int
	fmt.Scan(&w, &h, &c)
	var text []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return Solve1J(w, h, c, text)
}

func Solve1J(w, h, c int, text []string) []Point {
	var placed []Placed
	var paragraphPlaced []Placed
	lineStart, lineEnd := 0, h
	for _, element := range parse(text) {
		switch element.(type) {
		case Image:
			params := element.(Image)
			switch params.layout {
			case Embedded:
				ls, le, x, y := findPos(w, h, c, lineStart, lineEnd, params.width, params.height, string(params.layout), paragraphPlaced)
				lineStart, lineEnd = ls, le
				paragraphPlaced = append(paragraphPlaced, Placed{
					Point{x, y}, Point{x + params.width, y + params.height}, string(params.layout),
				})
			case Surrounded:
				ls, le, x, y := findPos(w, h, c, lineStart, lineEnd, params.width, 0, string(params.layout), paragraphPlaced)
				lineStart, lineEnd = ls, le
				paragraphPlaced = append(paragraphPlaced, Placed{
					Point{x, y}, Point{x + params.width, y + params.height}, string(params.layout),
				})
			case Floating:
				var lastElemX, lastElemY int
				if len(paragraphPlaced) > 0 {
					lastElem := paragraphPlaced[len(paragraphPlaced)-1]
					lastElemX, lastElemY = lastElem.end.X, lastElem.start.Y
				} else {
					lastElemX, lastElemY = 0, lineStart
				}
				if params.offset.dx+params.width+lastElemX > w {
					paragraphPlaced = append(paragraphPlaced, Placed{
						Point{w - params.width, lastElemY + params.offset.dy},
						Point{w, lastElemY + params.offset.dy + params.height},
						"floating",
					})
				} else if params.offset.dx+lastElemX < 0 {
					paragraphPlaced = append(paragraphPlaced, Placed{
						Point{0, lastElemY + params.offset.dy},
						Point{params.width, lastElemY + params.offset.dy + params.height},
						"floating",
					})
				} else {
					paragraphPlaced = append(paragraphPlaced, Placed{
						Point{lastElemX + params.offset.dx, lastElemY + params.offset.dy},
						Point{lastElemX + params.offset.dx + params.width, lastElemY + params.offset.dy + params.height},
						"floating",
					})
				}
			}
		case Word:
			ls, le, x, y := findPos(w, h, c, lineStart, lineEnd, element.(Word).length*c, 0, "word", paragraphPlaced)
			lineStart, lineEnd = ls, le
			paragraphPlaced = append(paragraphPlaced, Placed{Point{x, y}, Point{x + element.(Word).length*c, y}, "word"})
		case nil:
			lowest := lineEnd
			for _, el := range paragraphPlaced {
				if el.layout == "surrounded" {
					lowest = int(math.Max(float64(lowest), float64(el.end.Y)))
				}
			}
			if len(paragraphPlaced) == 0 {
				lowest = lineStart
			}
			lineStart, lineEnd = lowest, lowest+h
			placed = append(placed, paragraphPlaced...)
			paragraphPlaced = nil
		}
	}
	placed = append(placed, paragraphPlaced...)

	var coords []Point
	for _, el := range placed {
		if el.layout == "surrounded" || el.layout == "embedded" || el.layout == "floating" {
			coords = append(coords, el.start)
		}
	}
	return coords
}

func findPos(w, h, c, lineStart, lineEnd, desiredW, desiredH int, desiredLayout string, placed []Placed) (int, int, int, int) {
	for {
		linePlaced := []Placed{
			{Point{0, 0}, Point{0, 0}, "none"},
			{Point{w, h}, Point{w, h}, "none"},
		}
		for _, el := range placed {
			if (el.start.Y == lineStart || el.end.Y > lineStart) && el.layout != "floating" {
				linePlaced = append(linePlaced, el)
			}
		}
		sort.Slice(linePlaced, func(i, j int) bool {
			if linePlaced[i].start.X != linePlaced[j].start.X {
				return linePlaced[i].start.X < linePlaced[j].start.X
			}
			return linePlaced[i].end.X < linePlaced[j].end.X
		})
		last := 0
		for _, el := range linePlaced {
			if el.layout != "surrounded" && el.layout != "none" {
				last = el.end.X
			}
		}

		for i := 0; i < len(linePlaced)-1; i++ {
			currEnd, nextStart := linePlaced[i].end.X, linePlaced[i+1].start.X
			if linePlaced[i+1].layout != "surrounded" && linePlaced[i+1].layout != "none" {
				nextStart -= c
			}
			if linePlaced[i].layout != "none" && linePlaced[i].layout != "surrounded" && desiredLayout != "surrounded" {
				currEnd += c
			}
			if nextStart-currEnd >= desiredW && currEnd >= last {
				return lineStart, int(math.Max(float64(lineEnd), float64(lineStart+desiredH))), currEnd, lineStart
			}
		}
		lineStart, lineEnd = lineEnd, lineEnd+h
	}
}

func parse(text []string) []any {
	var tokens []string
	for _, line := range text {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			tokens = append(tokens, "")
		}
		lineTokens := strings.Split(trimmed, " ")
		tokens = append(tokens, lineTokens...)
	}

	var elements []interface{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "" {
			elements = append(elements, NewParagraph(nil))
		} else if tokens[i] == "(image" {
			img := Image{}
			for i++; ; i++ {
				param := strings.Split(tokens[i], "=")
				name, value := param[0], strings.ReplaceAll(param[1], ")", "")
				switch name {
				case "layout":
					img.layout = Layout(value)
				case "width":
					intValue, _ := strconv.Atoi(value)
					img.width = intValue
				case "height":
					intValue, _ := strconv.Atoi(value)
					img.height = intValue
				case "dx":
					if img.offset == nil {
						img.offset = &Offset{}
					}
					intValue, _ := strconv.Atoi(value)
					img.offset.dx = intValue
				case "dy":
					if img.offset == nil {
						img.offset = &Offset{}
					}
					intValue, _ := strconv.Atoi(value)
					img.offset.dy = intValue
				}
				if string(tokens[i][len(tokens[i])-1]) == ")" {
					break
				}
			}
			elements = append(elements, img)
		} else {
			elements = append(elements, Word{len(tokens[i])})
		}
	}

	return elements
}
