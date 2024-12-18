package answer_test

import (
	"testing"

	"day15/part2/answer"
)

func TestCompute(t *testing.T) {
	warehouse, moves := answer.Parse("sample.txt")
	expected := 9021
	answer := answer.Compute(warehouse, moves)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
