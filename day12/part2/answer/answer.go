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
	plots     []XY
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

var directions = []XY{
	{x: 0, y: -1}, // Up
	{x: 1, y: 0},  // Right
	{x: 0, y: 1},  // Down
	{x: -1, y: 0}, // Left
}

func check(pos XY, grid [][]string, region *Region) (int, int) {
	id := fmt.Sprintf("%d,%d", pos.x, pos.y)
	_, ok := seen[id]
	if ok {
		return 0, 0
	}

	area := 1
	perim := 0
	seen[id] = struct{}{}
	plotType := grid[pos.y][pos.x]

	for _, direction := range directions {
		newPos := XY{x: pos.x + direction.x, y: pos.y + direction.y}
		if !inBounds(newPos, grid) {
			perim++
		} else if grid[newPos.y][newPos.x] != plotType {
			perim++
		} else {
			a, p := check(newPos, grid, region)
			area += a
			perim += p
		}
	}

	region.plots = append(region.plots, pos)

	return area, perim
}

var seen = map[string]struct{}{}
var regions = map[int]Region{}
var visitedPerimeter = map[string]bool{}

func countStraightSides(region Region, grid [][]string) int {
	// Track the number of straight sides
	sides := 0
	currentDirection := -1

	// Perform a DFS-like traversal around the perimeter
	for _, plot := range region.plots {
		plotId := fmt.Sprintf("%d,%d", plot.x, plot.y)
		visitedPerimeter[plotId] = true

		y, x := plot.y, plot.x
		// Try moving in each direction (clockwise)
		for i := 0; i < 4; i++ {
			newY, newX := plot.y+directions[i].y, plot.x+directions[i].x
			newPlotId := fmt.Sprintf("%d,%d", newX, newY)
			if isPerimeter(newY, newX, grid[y][x], grid) && !visitedPerimeter[newPlotId] {
				if currentDirection != i {
					sides++
					currentDirection = i
				}
				y, x = newY, newX
			}
		}

		// If we are back at the starting point, break
		if y == plot.y && x == plot.x {
			break
		}
	}

	return sides
}

// isPerimeter checks if the cell (r, c) is part of the perimeter of the region
func isPerimeter(y, x int, regionChar string, grid [][]string) bool {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
		return false
	}
	if grid[y][x] != regionChar {
		return false
	}

	// Check if any neighboring cell is out of bounds or a different region
	for _, dir := range directions {
		newR, newC := y+dir.y, x+dir.x
		if newR < 0 || newR >= len(grid) || newC < 0 || newC >= len(grid[0]) {
			return true
		}
		if grid[newR][newC] != regionChar {
			return true
		}
	}
	return false
}

func Compute(grid [][]string) int {
	regionId := 0
	for y := range grid {
		for x := range grid[y] {
			id := fmt.Sprintf("%d,%d", x, y)
			_, ok := seen[id]
			if ok {
				continue
			}

			region := Region{}
			area, perim := check(XY{y: y, x: x}, grid, &region)

			region.area = area
			region.perimeter = perim
			regions[regionId] = region
			regionId++
		}
	}

	fmt.Println(regions)

	cost := 0
	for _, region := range regions {
		sides := countStraightSides(region, grid)
		fmt.Println(sides)
		cost += region.area * sides
	}

	return cost
}
