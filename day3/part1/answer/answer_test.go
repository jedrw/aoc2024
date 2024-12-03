package answer_test

import (
	"testing"

	"github.com/jedrw/aoc2024/day3/part1/answer"
)

func TestCompute(t *testing.T) {
	expected := 161
	answer := answer.Compute("sample.txt")

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
