package answer_test

import (
	"reflect"
	"testing"
	"unicode"

	"day9/part2/answer"
)

func TestParse(t *testing.T) {
	parsed := answer.Parse("sample.txt")
	expectedString := "00...111...2...333.44.5555.6666.777.888899"
	expected := []int{}
	for _, r := range expectedString {
		if unicode.IsDigit(r) {
			expected = append(expected, int(r-'0'))
		} else {
			expected = append(expected, -1)
		}

	}

	if !reflect.DeepEqual(parsed, expected) {
		t.Errorf("expected: %+v, got: %+v", expected, parsed)
	}
}

func TestCompute(t *testing.T) {
	blocks := answer.Parse("sample.txt")
	expected := 2858
	answer := answer.Compute(blocks)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
