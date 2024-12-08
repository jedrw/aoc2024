package answer_test

import (
	"testing"

	"day6/part2/answer"
)

func TestCompute(t *testing.T) {
	input := answer.Parse("sample.txt")
	expected := 6
	answer := answer.Compute(input)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
