package main

import (
	"fmt"

	"day7/part1/answer"
)

func main() {
	input := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(input))
}
