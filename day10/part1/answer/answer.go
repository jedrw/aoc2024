package answer

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type XY struct {
	x int
	y int
}

func Parse(filePath string) ([][]int, []XY) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	grid := [][]int{}
	trailHeads := []XY{}
	scanner := bufio.NewScanner(file)
	rowNum := 0
	for scanner.Scan() {
		row := []int{}
		for i, posString := range strings.Split(strings.Trim(scanner.Text(), " "), "") {
			pos, err := strconv.Atoi(posString)
			if err != nil {
				panic(err)
			}

			if pos == 0 {
				trailHeads = append(trailHeads, XY{x: i, y: rowNum})
			}

			row = append(row, pos)

		}

		grid = append(grid, row)
		rowNum++
	}

	return grid, trailHeads
}

var directions = []XY{
	{x: 0, y: -1}, // Up
	{x: 0, y: 1},  // Down
	{x: -1, y: 0}, // Left
	{x: 1, y: 0},  // Right
}

func inBounds(pos XY, grid [][]int) bool {
	return pos.y >= 0 &&
		pos.y < len(grid) &&
		pos.x >= 0 &&
		pos.x < len(grid[0])

}

func walkTrail(pos XY, grid [][]int, paths *[]string) {
	for _, direction := range directions {
		newPos := XY{x: pos.x + direction.x, y: pos.y + direction.y}
		if inBounds(newPos, grid) {
			height := grid[newPos.y][newPos.x]
			if height-grid[pos.y][pos.x] == 1 {
				if height == 9 {
					trailEnd := fmt.Sprintf("%d,%d", newPos.x, newPos.y)
					if !slices.Contains(*paths, trailEnd) {
						*paths = append(*paths, trailEnd)
					}
				}

				walkTrail(newPos, grid, paths)
			}
		}
		continue
	}
}

func Compute(grid [][]int, trailHeads []XY) int {
	score := 0
	for _, head := range trailHeads {
		trailsEnds := []string{}
		walkTrail(head, grid, &trailsEnds)
		score += len(trailsEnds)
	}

	return score
}
