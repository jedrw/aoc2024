package answer_test

import (
	"testing"

	"day5/part2/answer"
)

func TestCompute(t *testing.T) {
	rules, updates := answer.Parse("sample.txt")
	expected := 123
	answer := answer.Compute(rules, updates)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
