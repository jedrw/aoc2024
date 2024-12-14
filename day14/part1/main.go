package main

import (
	"fmt"

	"day14/part1/answer"
)

func main() {
	input := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(input, answer.XY{X: 101, Y: 103}, 100))
}
