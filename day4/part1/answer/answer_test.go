package answer_test

import (
	"testing"

	"github.com/jedrw/aoc2024/day4/part1/answer"
)

func TestCompute(t *testing.T) {
	input := answer.Parse("sample.txt")
	expected := 18
	answer := answer.Compute(input)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
