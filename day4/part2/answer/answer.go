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

	input := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func check(input string) bool {
	for _, ms := range []string{"MS", "SM"} {
		if input == ms {
			return true
		}
	}

	return false
}

func Compute(input []string) int {
	inputWidth := len(input[0])
	inputHeight := len(input)
	total := 0

	for x := 1; x < inputHeight-1; x++ {
		for y := 1; y < inputWidth-1; y++ {
			if string(input[y][x]) == "A" {
				if !check(string(input[y-1][x-1]) + string(input[y+1][x+1])) {
					continue
				}

				if !check(string(input[y+1][x-1]) + string(input[y-1][x+1])) {
					continue
				}

				total++
			}
		}
	}

	return total
}
