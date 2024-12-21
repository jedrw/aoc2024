package answer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type XY struct {
	X int
	Y int
}

func Parse(filePath string) []XY {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	bytesPositions := []XY{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xy := strings.Split(scanner.Text(), ",")
		xInt, _ := strconv.Atoi(xy[0])
		yInt, _ := strconv.Atoi(xy[1])
		bytesPositions = append(bytesPositions, XY{X: xInt, Y: yInt})
	}

	return bytesPositions
}

var (
	grid         = [][]string{}
	shortestPath int
	possible     bool
	bestToPos    = map[string]int{}
	directions   = map[string]XY{
		"^": {X: 0, Y: -1},
		"v": {X: 0, Y: 1},
		"<": {X: -1, Y: 0},
		">": {X: 1, Y: 0},
	}
)

func inBounds(pos XY) bool {
	return pos.Y >= 0 &&
		pos.Y < len(grid) &&
		pos.X >= 0 &&
		pos.X < len(grid[0])

}

func walk(pos XY, direction string, steps int) {
	posString := fmt.Sprintf("%d,%d", pos.X, pos.Y)
	bestSteps, haveSeen := bestToPos[posString]
	if haveSeen {
		if steps >= bestSteps {
			return
		}
	}

	bestToPos[posString] = steps

	if pos.Y == len(grid)-1 && pos.X == len(grid[0])-1 {
		possible = true
		return
	}

	var nextDirections []string
	switch direction {
	case "^", "v":
		nextDirections = []string{direction, "<", ">"}
	case "<", ">":
		nextDirections = []string{direction, "^", "v"}
	}

	for _, nextDirection := range nextDirections {
		nextPos := XY{X: pos.X + directions[nextDirection].X, Y: pos.Y + directions[nextDirection].Y}
		if !inBounds(nextPos) {
			continue
		}

		if grid[nextPos.Y][nextPos.X] == "." {
			walk(nextPos, nextDirection, steps+1)
		}
	}
}

func Compute(input []XY, gridDims, numIntialBytes int) string {
	shortestPath = gridDims * gridDims
	for range gridDims {
		row := []string{}
		for range gridDims {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	for _, xy := range input[:numIntialBytes] {
		grid[xy.Y][xy.X] = "#"
	}

	var lastPossible XY
	for i, xy := range input[numIntialBytes:] {
		fmt.Printf("Checking %d: %+v\n", i+numIntialBytes, xy)
		shortestPath = gridDims * gridDims
		bestToPos = map[string]int{}
		grid[xy.Y][xy.X] = "#"
		possible = false
		walk(XY{X: 0, Y: 0}, ">", 0)
		if !possible {
			lastPossible = xy
			break
		}
	}

	return fmt.Sprintf("%d,%d", lastPossible.X, lastPossible.Y)
}
