package answer

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func parse(file *os.File) string {
	line := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line += scanner.Text()
	}

	return line
}

func getUpToThreeDigits(input string) (bool, string) {
	if !unicode.IsDigit(rune(input[0])) {
		return false, ""
	}

	digits := string(input[0])
	for _, r := range input[1:] {
		if !unicode.IsDigit(r) {
			return true, digits
		}

		digits += string(r)
	}

	return true, digits
}

func Compute(inputFile string) int {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	line := parse(file)

	sum := 0
	do := true
	// Begin gobbling
	for i := range line {
		if line[i:i+7] == "don't()" {
			do = false
		} else if line[i:i+4] == "do()" {
			do = true
		}

		if line[i:i+4] == "mul(" {
			ok, firstDigits := getUpToThreeDigits(line[i+4 : i+7])
			if ok {
				a, err := strconv.Atoi(firstDigits)
				if err != nil {
					panic(err)
				}

				numFirstDigits := len(firstDigits)
				if line[i+4+numFirstDigits] == ',' {
					ok, secondDigits := getUpToThreeDigits(line[i+5+numFirstDigits : i+8+numFirstDigits])
					if ok {
						b, err := strconv.Atoi(secondDigits)
						if err != nil {
							panic(err)
						}

						numSecondDigits := len(secondDigits)
						if line[i+5+numFirstDigits+numSecondDigits] == ')' {
							if do {
								sum += (a * b)
							}
							i = i + 6 + numFirstDigits + numSecondDigits
							lengthToEnd := len(line) - i
							if lengthToEnd < 6 {
								return sum
							}
						}
					}
				}
			}
		}
	}

	return sum
}
