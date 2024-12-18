package answer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	wall     = "#"
	free     = "."
	boxLeft  = "["
	boxRight = "]"
	robot    = "@"
)

func Parse(filePath string) ([][]*string, []string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	lines, err := io.ReadAll(file)
	inputSections := strings.Split(string(lines), "\n\n")

	warehouseString := inputSections[0]
	warehouse := [][]*string{}
	for _, line := range strings.Split(warehouseString, "\n") {
		warehouseY := []*string{}
		for _, cell := range line {
			switch string(cell) {
			case "#":
				warehouseY = append(warehouseY, []*string{&wall, &wall}...)
			case "O":
				warehouseY = append(warehouseY, []*string{&boxLeft, &boxRight}...)
			case ".":
				warehouseY = append(warehouseY, []*string{&free, &free}...)
			case "@":
				warehouseY = append(warehouseY, []*string{&robot, &free}...)
			}
		}

		warehouse = append(warehouse, warehouseY)
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

func canMoveBox(boxLeftPos XY, boxRightPos XY, direction string) bool {
	canMove := true
	move := directions[direction]
	switch direction {
	case "<":
		nextPos := XY{X: boxLeftPos.X + move.X, Y: boxLeftPos.Y + move.Y}
		switch *warehouse[nextPos.Y][nextPos.X] {
		case free:
			return true
		case wall:
			return false
		case boxRight:
			return canMoveBox(XY{X: nextPos.X - 1, Y: nextPos.Y}, nextPos, direction)
		}
	case ">":
		nextPos := XY{X: boxRightPos.X + move.X, Y: boxRightPos.Y + move.Y}
		switch *warehouse[nextPos.Y][nextPos.X] {
		case free:
			return true
		case wall:
			return false
		case boxRight:
			return canMoveBox(nextPos, XY{X: nextPos.X + 1, Y: nextPos.Y}, direction)
		}
	case "^", "v":
		leftCanMove := false
		leftNext := XY{X: boxLeftPos.X + move.X, Y: boxLeftPos.Y + move.Y}
		switch *warehouse[leftNext.Y][leftNext.X] {
		case free:
			leftCanMove = true
		case wall:
			leftCanMove = false
		case boxLeft:
			leftCanMove = canMoveBox(leftNext, XY{X: leftNext.X + 1, Y: leftNext.Y}, direction)
		case boxRight:
			leftCanMove = canMoveBox(XY{X: leftNext.X - 1, Y: leftNext.Y}, leftNext, direction)
		}

		rightCanMove := false
		rightNext := XY{X: boxRightPos.X + move.X, Y: boxRightPos.Y + move.Y}
		switch *warehouse[rightNext.Y][rightNext.X] {
		case free:
			rightCanMove = true
		case wall:
			rightCanMove = false
		case boxLeft:
			rightCanMove = canMoveBox(rightNext, XY{X: rightNext.X + 1, Y: rightNext.Y}, direction)
		case boxRight:
			rightCanMove = canMoveBox(XY{X: rightNext.X - 1, Y: rightNext.Y}, rightNext, direction)
		}

		for _, check := range []bool{leftCanMove, rightCanMove} {
			if check == false {
				canMove = false
			}
		}
	}

	return canMove
}

func moveBox(leftBox XY, rightBox XY, direction string) {
	move := directions[direction]
	nextPosLeft := XY{X: leftBox.X + move.X, Y: leftBox.Y + move.Y}
	switch *warehouse[nextPosLeft.Y][nextPosLeft.X] {
	case boxLeft:
		moveBox(nextPosLeft, XY{X: nextPosLeft.X + 1, Y: nextPosLeft.Y}, direction)
	case boxRight:
		moveBox(XY{X: nextPosLeft.X - 1, Y: nextPosLeft.Y}, nextPosLeft, direction)
	}

	nextPosRight := XY{X: rightBox.X + move.X, Y: rightBox.Y + move.Y}
	switch *warehouse[nextPosRight.Y][nextPosRight.X] {
	case boxLeft:
		moveBox(nextPosRight, XY{X: nextPosRight.X + 1, Y: nextPosRight.Y}, direction)
	case boxRight:
		moveBox(XY{X: nextPosRight.X - 1, Y: nextPosRight.Y}, nextPosRight, direction)
	}

	if warehouse[boxPos.Y+move.Y][boxPos.X+move.X+1] == nil ||
		*warehouse[boxPos.Y+move.Y][boxPos.X+move.X] == box {
		moveBox(XY{X: boxPos.X + move.X, Y: boxPos.Y + move.Y}, move)
		warehouse[boxPos.Y+move.Y][boxPos.X+move.X] = &box
		warehouse[boxPos.Y+move.Y][boxPos.X+move.X+1] = &box
	}
}

func tryMove(pos XY, direction string) (XY, error) {
	move := directions[direction]
	nextPos := XY{X: pos.X + move.X, Y: pos.Y + move.Y}
	nextSym := warehouse[nextPos.Y][nextPos.X]

	switch *nextSym {
	case free:
		warehouse[nextPos.Y][nextPos.X] = &robot
		warehouse[pos.Y][pos.X] = &free
		return nextPos, nil

	case wall:
		return pos, nil

	case boxLeft:
		rightBox := XY{X: nextPos.X + 1, Y: nextPos.Y}
		if canMoveBox(nextPos, rightBox, direction) {
			moveBox(nextPos, rightBox, direction)
			warehouse[nextPos.Y][nextPos.X] = &robot
			warehouse[pos.Y][pos.X] = &free

			return nextPos, nil
		}
		// peek := XY{X: next.X + move.X, Y: next.Y + move.Y}
		// for *warehouse[peek.Y][peek.X] != "#" {
		// 	if *warehouse[peek.Y][peek.X] == "." {
		// 		*warehouse[peek.Y][peek.X] = "O"
		// 		*warehouse[next.Y][next.X] = "@"
		// 		*warehouse[pos.Y][pos.X] = "."
		// 		return next, nil
		// 	}
		// 	if *warehouse[peek.Y][peek.X] == "O" {
		// 		peek = XY{X: peek.X + move.X, Y: peek.Y + move.Y}
		// 	}
		// }

		return pos, nil

	default:
		return XY{}, errors.New(fmt.Sprintf("uh oh spaghettios: %s", *warehouse[nextPos.Y][nextPos.X]))
	}
}

var warehouse [][]*string

func Compute(warehouseInput [][]*string, moves []string) int {
	warehouse = warehouseInput
	var pos XY
	for y, row := range warehouse {
		for x, cell := range row {
			if cell != nil {
				if *cell == robot {
					pos = XY{X: x, Y: y}
				}
			}
		}
	}

	var err error
	for _, direction := range moves {
		pos, err = tryMove(pos, direction)
		if err != nil {
			fmt.Println(err)
			return 0
		}
	}

	sum := 0
	for y, row := range warehouse {
		for x, cell := range row {
			if *cell == boxLeft {
				sum += (100 * y) + x
			}
		}
	}

	return sum
}
