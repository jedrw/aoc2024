package main

import (
	"fmt"

	"day6/part1/answer"
)

func main() {
	input := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(input))
}
