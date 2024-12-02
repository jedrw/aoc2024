package answer_test

import (
	"testing"

	"github.com/jedrw/aoc2024/day2/part1/answer"
)

func TestCompute(t *testing.T) {
	expected := 2
	answer := answer.Compute("sample.txt")

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
