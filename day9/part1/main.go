package main

import (
	"fmt"

	"day9/part1/answer"
)

func main() {
	input := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(input))
}
