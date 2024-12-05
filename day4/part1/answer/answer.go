package answer

import (
	"bufio"
	"fmt"
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

func check(input string) int {
	for _, xmas := range []string{"XMAS", "SAMX"} {
		if input == xmas {
			return 1
		}
	}

	return 0
}

func Compute(input []string) int {
	inputWidth := len(input[0])
	inputHeight := len(input)
	total := 0

	// check rows
	for _, row := range input {
		for i := 0; i < len(row)-3; i++ {
			total += check(row[i : i+4])
		}
	}

	// check cols
	for x := 0; x < inputWidth; x++ {
		for y := 0; y < inputHeight-3; y++ {
			col := string(input[y][x]) + string(input[y+1][x]) + string(input[y+2][x]) + string(input[y+3][x])
			total += check(col)
		}
	}

	// check diagonals
	for x := 0; x < inputWidth-3; x++ {
		for y := 0; y < inputHeight-3; y++ {
			diag := ""
			for _, i := range []int{0, 1, 2, 3} {
				diag += string(input[y+i][x+i])
			}
			total += check(diag)

			diag = ""
			for _, pair := range [][]int{{0, 3}, {1, 2}, {2, 1}, {3, 0}} {
				diag += string(input[y+pair[0]][x+pair[1]])
			}
			fmt.Println(diag)
			total += check(diag)
		}
	}

	return total
}
