package answer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Parse(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	grid := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	return grid
}

type Region struct {
	area      int
	perimeter int
}

type XY struct {
	x int
	y int
}

func inBounds(pos XY, grid [][]string) bool {
	return pos.y >= 0 &&
		pos.y < len(grid) &&
		pos.x >= 0 &&
		pos.x < len(grid[0])
}

func check(pos XY, grid [][]string) (int, int) {
	id := fmt.Sprintf("%d,%d", pos.x, pos.y)
	_, ok := seen[id]
	if ok {
		return 0, 0
	}

	area := 1
	perim := 0
	seen[id] = struct{}{}
	plotType := grid[pos.y][pos.x]

	for _, direction := range []XY{
		{x: 0, y: -1}, // Up
		{x: 0, y: 1},  // Down
		{x: -1, y: 0}, // Left
		{x: 1, y: 0},  // Right
	} {
		newPos := XY{x: pos.x + direction.x, y: pos.y + direction.y}
		if !inBounds(newPos, grid) {
			perim++
		} else if grid[newPos.y][newPos.x] != plotType {
			perim++
		} else {
			a, p := check(newPos, grid)
			area += a
			perim += p
		}
	}

	return area, perim
}

var seen = map[string]struct{}{}

func Compute(grid [][]string) int {
	regions := map[int]Region{}

	regionId := 0
	for y := range grid {
		for x := range grid[y] {
			area, perim := check(XY{y: y, x: x}, grid)

			regions[regionId] = Region{area: area, perimeter: perim}
			regionId++
		}
	}

	cost := 0
	for _, region := range regions {
		cost += region.area * region.perimeter
	}

	return cost
}
