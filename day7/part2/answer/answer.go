package answer

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Parse(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	tests := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(strings.ReplaceAll(scanner.Text(), ":", ""), " ")
		test := []int{}
		for _, s := range line {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			test = append(test, i)
		}

		tests = append(tests, test)
	}

	return tests
}

func generateCombinations(symbols []string, length int) []string {
	var result []string
	var current []string
	helper(symbols, length, current, &result)
	return result
}

func helper(symbols []string, length int, current []string, result *[]string) {
	if len(current) == length {
		*result = append(*result, join(current))
		return
	}

	for _, symbol := range symbols {
		helper(symbols, length, append(current, symbol), result)
	}
}

func join(parts []string) string {
	result := ""
	for _, part := range parts {
		result += part
	}
	return result
}

func Compute(input [][]int) int {
	total := 0

	for _, test := range input {
		value := test[0]
		numbers := test[1:]
		permutations := generateCombinations([]string{"*", "+", "|"}, len(numbers)-1)
		for _, perm := range permutations {
			operators := strings.Split(perm, "")
			res := numbers[0]
			for i, num := range numbers[1:] {
				switch operators[i] {
				case "*":
					res *= num
				case "+":
					res += num
				case "|":
					var err error
					res, err = strconv.Atoi(strconv.Itoa(res) + strconv.Itoa(num))
					if err != nil {
						panic(err)
					}
				}
			}
			if res == value {
				total += value
				break
			}
		}

	}

	return total
}
