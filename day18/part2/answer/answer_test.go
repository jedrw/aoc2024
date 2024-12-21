package answer_test

import (
	"testing"

	"day18/part2/answer"
)

func TestCompute(t *testing.T) {
	input := answer.Parse("sample.txt")
	expected := "6,1"
	answer := answer.Compute(input, 7, 12)
	if answer != expected {
		t.Errorf("expected %s, got %s", expected, answer)
	}
}
