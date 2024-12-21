package answer_test

import (
	"testing"

	"day18/part1/answer"
)

func TestCompute(t *testing.T) {
	input := answer.Parse("sample.txt")
	expected := 22
	answer := answer.Compute(input, 7, 12)
	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
