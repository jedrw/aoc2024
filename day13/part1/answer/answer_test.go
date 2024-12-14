package answer_test

import (
	"testing"

	"day13/part1/answer"
)

func TestCompute(t *testing.T) {
	input := answer.Parse("sample.txt")
	expected := 480
	answer := answer.Compute(input)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
