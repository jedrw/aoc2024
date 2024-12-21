package answer

import (
	"bufio"
	"fmt"
	"math"
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
		line := []string{}
		for _, cell := range strings.Split(scanner.Text(), "") {
			line = append(line, cell)
		}

		grid = append(grid, line)
	}

	return grid
}

type XY struct {
	X int
	Y int
}

var (
	grid       = [][]string{}
	bestScore  = math.MaxInt
	directions = map[string]XY{
		"^": {X: 0, Y: -1},
		"v": {X: 0, Y: 1},
		"<": {X: -1, Y: 0},
		">": {X: 1, Y: 0},
	}
)

func walk(pos XY, direction string, steps, turns int, seen map[string]string) {
	seenString := fmt.Sprintf("%d,%d", pos.X, pos.Y)
	newSeen := make(map[string]string, len(seen))
	_, haveSeen := seen[seenString]
	if haveSeen {
		return
	} else {
		for k, v := range seen {
			newSeen[k] = v
		}
		newSeen[seenString] = direction
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
		addTurn := 0
		if nextDirection != direction {
			addTurn = 1
		}

		switch grid[nextPos.Y][nextPos.X] {
		case "#":
			continue
		case ".":
			walk(nextPos, nextDirection, steps+1, turns+addTurn, newSeen)
		case "E":
			score := (steps + 1) + (turns * 1000)
			if score < bestScore {
				bestScore = score
			}

			return
		}
	}
}

func Compute(input [][]string) int {
	grid = input
	startPos := XY{}
	for y, row := range grid {
		for x, c := range row {
			if c == "S" {
				startPos.X = x
				startPos.Y = y
			}
		}
	}

	walk(startPos, ">", 0, 0, map[string]string{})
	return bestScore
}
