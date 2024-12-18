package main

import (
	"fmt"

	"day15/part1/answer"
)

func main() {
	warehouse, moves := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(warehouse, moves))
}
