package answer

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type XY struct {
	X int
	Y int
}

type Robot struct {
	pos   XY
	delta XY
}

func Parse(filePath string) []Robot {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	robots := []Robot{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		startPos := strings.Split(strings.Split(line[0], "=")[1], ",")
		startX, _ := strconv.Atoi(startPos[0])
		startY, _ := strconv.Atoi(startPos[1])

		delta := strings.Split(strings.Split(line[1], "=")[1], ",")
		deltaX, _ := strconv.Atoi(delta[0])
		deltaY, _ := strconv.Atoi(delta[1])

		robots = append(robots, Robot{
			pos:   XY{X: startX, Y: startY},
			delta: XY{X: deltaX, Y: deltaY},
		})
	}

	return robots
}

func move(robot Robot, bounds XY, seconds int, quadrants []int) {
	if seconds == 0 {
		if robot.pos.Y < bounds.Y/2 {
			if robot.pos.X < bounds.X/2 {
				quadrants[0]++
			}
			if robot.pos.X > bounds.X/2 {
				quadrants[1]++
			}
		} else if robot.pos.Y > bounds.Y/2 {
			if robot.pos.X > bounds.X/2 {
				quadrants[2]++
			}
			if robot.pos.X < bounds.X/2 {
				quadrants[3]++
			}
		}

		return
	}

	newX := (robot.pos.X + robot.delta.X) % bounds.X
	if newX < 0 {
		newX = bounds.X + newX
	}

	newY := (robot.pos.Y + robot.delta.Y) % bounds.Y
	if newY < 0 {
		newY = bounds.Y + newY
	}

	robot.pos = XY{X: newX, Y: newY}
	move(robot, bounds, seconds-1, quadrants)
}

func Compute(robots []Robot, bounds XY, seconds int) int {
	quadrants := make([]int, 4)

	for _, robot := range robots {
		move(robot, bounds, seconds, quadrants)
	}

	score := quadrants[0]
	for _, quad := range quadrants[1:] {
		score *= quad
	}

	return score
}
