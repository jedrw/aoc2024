package main

import (
	"fmt"

	"day18/part1/answer"
)

func main() {
	input := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(input, 71, 1024))
}
