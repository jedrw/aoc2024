package main

import (
	"fmt"

	"day13/part2/answer"
)

func main() {
	input := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(input))
}
