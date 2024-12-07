package answer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var area = [][]string{}

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

func inBounds(pos XY, direction XY) bool {
	return pos.y+direction.y > -1 &&
		pos.y+direction.y < len(area) &&
		pos.x+direction.x > -1 &&
		pos.x+direction.x < len(area[0])

}

func Compute(input [][]string) int {
	pos := findStart()
	visited := map[string]struct{}{}
	visited[fmt.Sprintf("%d,%d", pos.y, pos.x)] = struct{}{}
	symbol := area[pos.y][pos.x]
	direction := directions[symbol]

	for inBounds(pos, direction) {
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
		visited[fmt.Sprintf("%d,%d", pos.y, pos.x)] = struct{}{}
	}

	return len(visited)
}
