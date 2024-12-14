package answer_test

import (
	"testing"

	"day14/part2/answer"
)

func TestCompute(t *testing.T) {
	input := answer.Parse("sample.txt")
	expected := 12

	answer := answer.Compute(input, answer.XY{X: 11, Y: 7}, 100)
	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
