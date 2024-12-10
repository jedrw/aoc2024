package main

import (
	"fmt"

	"day10/part1/answer"
)

func main() {
	grid, trailHeads := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(grid, trailHeads))
}
