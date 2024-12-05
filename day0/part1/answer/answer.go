package answer

import (
	"bufio"
	"os"
)

func Parse(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Compute(input []string) int {

	return len(input)
}
