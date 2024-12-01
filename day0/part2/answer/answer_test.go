package answer_test

import (
	"testing"

	"github.com/jedrw/aoc2024/day0/part2/answer"
)

func TestCompute(t *testing.T) {
	expected := 0
	answer := answer.Compute("sample.txt")

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
