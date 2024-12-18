package answer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func Parse(filePath string) ([][]string, []string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	lines, err := io.ReadAll(file)
	inputSections := strings.Split(string(lines), "\n\n")

	warehouseString := inputSections[0]
	warehouse := [][]string{}
	for _, line := range strings.Split(warehouseString, "\n") {
		warehouse = append(warehouse, strings.Split(line, ""))
	}

	movesString := inputSections[1]
	scanner := bufio.NewScanner(strings.NewReader(movesString))
	moves := []string{}
	for scanner.Scan() {
		movesLine := scanner.Text()
		moves = append(moves, strings.Split(movesLine, "")...)
	}

	return warehouse, moves
}

type XY struct {
	X int
	Y int
}

var directions = map[string]XY{
	"^": {X: 0, Y: -1},
	">": {X: 1, Y: 0},
	"v": {X: 0, Y: 1},
	"<": {X: -1, Y: 0},
}

func tryMove(pos XY, move XY) (XY, error) {
	next := XY{X: pos.X + move.X, Y: pos.Y + move.Y}
	switch warehouse[next.Y][next.X] {
	case ".":
		warehouse[next.Y][next.X] = "@"
		warehouse[pos.Y][pos.X] = "."
		return next, nil

	case "#":
		return pos, nil

	case "O":
		peek := XY{X: next.X + move.X, Y: next.Y + move.Y}
		for warehouse[peek.Y][peek.X] != "#" {
			if warehouse[peek.Y][peek.X] == "." {
				warehouse[peek.Y][peek.X] = "O"
				warehouse[next.Y][next.X] = "@"
				warehouse[pos.Y][pos.X] = "."
				return next, nil
			}
			if warehouse[peek.Y][peek.X] == "O" {
				peek = XY{X: peek.X + move.X, Y: peek.Y + move.Y}
			}
		}

		return pos, nil

	default:
		return XY{}, errors.New(fmt.Sprintf("uh oh spaghettios: %s", warehouse[next.Y][next.X]))
	}
}

var warehouse [][]string

func Compute(warehouseInput [][]string, moves []string) int {
	warehouse = warehouseInput
	var pos XY
	for y, row := range warehouse {
		for x, cell := range row {
			if cell == "@" {
				pos = XY{X: x, Y: y}
			}
		}
	}

	var err error
	moveNum := 0
	for _, direction := range moves {
		moveNum++
		pos, err = tryMove(pos, directions[direction])
		if err != nil {
			fmt.Println(err)
			return 0
		}
	}

	sum := 0
	for y, row := range warehouse {
		for x, cell := range row {
			if cell == "O" {
				sum += (100 * y) + x
			}
		}
	}

	return sum
}
