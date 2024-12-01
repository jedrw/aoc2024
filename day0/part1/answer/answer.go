package answer

import (
	"bufio"
	"os"
)

func parse(file *os.File) []string {
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Compute(inputFile string) int {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	lines := parse(file)
	return len(lines)
}
