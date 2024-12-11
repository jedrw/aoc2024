package answer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	stones := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, stone := range strings.Split(scanner.Text(), " ") {
			stones = append(stones, stone)
		}
	}

	return stones
}

func Compute(stones []string) int {
	for range 25 {
		numStones := len(stones)
		for i := 0; i < numStones; i++ {
			if stones[i] == "0" {
				stones[i] = "1"
			} else if len(stones[i])%2 == 0 {
				half := len(stones[i]) / 2
				first := stones[i][:half]
				second := strings.TrimLeft(stones[i][half:], "0")
				if second == "" {
					second = "0"
				}
				stones[i] = first
				stones = append(stones[:i+1], stones[i:]...)
				stones[i+1] = second
				i++
			} else {
				num, err := strconv.Atoi(stones[i])
				if err != nil {
					panic(err)
				}

				stones[i] = fmt.Sprint(num * 2024)
			}

			numStones = len(stones)
		}
	}

	return len(stones)
}
