package answer_test

import (
	"testing"

	"day11/part2/answer"
)

func TestCompute(t *testing.T) {
	input := answer.Parse("sample.txt")
	expected := 55312
	answer := answer.Compute(input, 25)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
