package answer_test

import (
	"testing"

	"day10/part2/answer"
)

func TestCompute(t *testing.T) {
	grid, trailheads := answer.Parse("sample.txt")
	expected := 81
	answer := answer.Compute(grid, trailheads)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
