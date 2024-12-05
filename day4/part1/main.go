package main

import (
	"fmt"

	"github.com/jedrw/aoc2024/day4/part1/answer"
)

func main() {
	input := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(input))
}
