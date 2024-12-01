package answer

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Parse(file *os.File) ([]int, []int) {
	var l1 = []int{}
	var l2 = []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "   ")

		l1Int, err := strconv.Atoi(strings.TrimSpace(data[0]))
		if err != nil {
			panic(err)
		}

		l2Int, err := strconv.Atoi(strings.TrimSpace(data[1]))
		if err != nil {
			panic(err)
		}

		l1 = append(l1, l1Int)
		l2 = append(l2, l2Int)
	}

	return l1, l2
}

func Compute(inputFile string) int {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	l1, l2 := Parse(file)

	total := 0
	for _, id := range l1 {
		count := 0
		for _, x := range l2 {
			if id == x {
				count++
			}
		}

		total += (id * count)
	}

	return total
}
