package answer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type XY struct {
	x int
	y int
}

var directions = map[string]XY{
	"^": {x: 0, y: -1},
	">": {x: 1, y: 0},
	"v": {x: 0, y: 1},
	"<": {x: -1, y: 0},
}

func Parse(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		area = append(area, strings.Split(strings.Trim(scanner.Text(), " "), ""))
	}

	return area
}

func findStart() XY {
	for y := range len(area) {
		for x := range len(area) {
			if area[y][x] == "^" || area[y][x] == "v" ||
				area[y][x] == "<" || area[y][x] == ">" {
				return XY{x, y}
			}
		}
	}

	panic("could not find start postion")
}

func peek(pos XY, direction XY) string {
	return area[pos.y+direction.y][pos.x+direction.x]
}

func peekInBounds(pos XY, direction XY) bool {
	return pos.y+direction.y > -1 &&
		pos.y+direction.y < len(area) &&
		pos.x+direction.x > -1 &&
		pos.x+direction.x < len(area[0])

}

func isLoop(pos, direction XY) bool {
	_, ok := visited[fmt.Sprintf("%d,%d", pos.y, pos.x)]
	if ok && peek(pos, direction) == "#" {
		return true
	}

	return false
}

func checkForLoop(pos XY, symbol string, checkArea [][]string) {
	startPos := pos
	startSym := symbol
	startDirection := directions[symbol]

	fmt.Printf("pos: %+v, direction: %s", pos, symbol)

	if !peekInBounds(pos, startDirection) {
		return
	}

	if peek(pos, startDirection) == "#" {
		return
	}

	switch symbol {
	case "^":
		symbol = ">"
	case ">":
		symbol = "v"
	case "v":
		symbol = "<"
	case "<":
		symbol = "^"
	}

	direction := directions[symbol]
	if peek(pos, direction) == "#" {
		return
	}

	for peekInBounds(pos, direction) {
		if isLoop(pos, direction) {
			fmt.Printf("found loop: %+v, direction: %s\n", startPos, startSym)
			count++

			return
		}

		if peek(pos, direction) == "#" {
			switch symbol {
			case "^":
				symbol = ">"
			case ">":
				symbol = "v"
			case "v":
				symbol = "<"
			case "<":
				symbol = "^"
			}

			direction = directions[symbol]

			continue
		}

		pos.y = pos.y + direction.y
		pos.x = pos.x + direction.x
	}
}

var area = [][]string{}
var visited = map[string]string{}
var count = 0

func Compute(input [][]string) int {
	startPos := findStart()
	pos := startPos
	symbol := area[pos.y][pos.x]
	direction := directions[symbol]
	visited[fmt.Sprintf("%d,%d", pos.y, pos.x)] = symbol

	for peekInBounds(pos, direction) {
		if peek(pos, direction) == "#" {
			switch symbol {
			case "^":
				symbol = ">"
			case ">":
				symbol = "v"
			case "v":
				symbol = "<"
			case "<":
				symbol = "^"
			}

			direction = directions[symbol]
			continue
		}

		pos.y = pos.y + direction.y
		pos.x = pos.x + direction.x
		visited[fmt.Sprintf("%d,%d", pos.y, pos.x)] = symbol
	}

	// checkForLoop(XY{x: 4, y: 6}, "<")
	// checkForLoop(XY{x: 6, y: 6}, "v")
	// checkForLoop(XY{x: 6, y: 7}, ">")
	// checkForLoop(XY{x: 2, y: 8}, "<")
	// checkForLoop(XY{x: 4, y: 8}, "<")
	// checkForLoop(XY{x: 7, y: 8}, "v")

	// checkForLoop(XY{x: 4, y: 1}, "^")

	for p, s := range visited {
		if p == fmt.Sprintf("%d,%d", startPos.y, startPos.x) {
			fmt.Println("skipping first")
			continue
		}

		py, err := strconv.Atoi(strings.Split(p, ",")[0])
		if err != nil {
			panic(err)
		}

		px, err := strconv.Atoi(strings.Split(p, ",")[1])
		if err != nil {
			panic(err)
		}

		checkForLoop(XY{x: px, y: py}, s)
	}

	return count
}
