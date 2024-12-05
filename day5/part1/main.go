package main

import (
	"fmt"

	"day5/part1/answer"
)

func main() {
	rules, updates := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(rules, updates))
}
