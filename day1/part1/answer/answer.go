package answer

import (
	"bufio"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parse(file *os.File) ([]int, []int) {
	var l1 = []int{}
	var l2 = []int{}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return l1, l2
		}

		data := strings.Split(line, "   ")

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

	l1, l2 := parse(file)

	sort.Ints(l1)
	sort.Ints(l2)

	total := 0
	for i := range l1 {
		total += int(math.Abs(float64(l1[i] - l2[i])))
	}

	return total
}
