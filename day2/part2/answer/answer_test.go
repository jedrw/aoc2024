package answer_test

import (
	"testing"

	"github.com/jedrw/aoc2024/day2/part2/answer"
)

func TestCompute(t *testing.T) {
	expected := 4
	answer := answer.Compute("sample.txt")

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
