package answer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	stones := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, stone := range strings.Split(scanner.Text(), " ") {
			num, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			stones = append(stones, num)
		}
	}

	return stones
}

func numDigits(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

type Split struct {
	left  int
	right int
}

var iSplit = map[int]Split{
	2024: {left: 20, right: 24},
}

func SplitInt(i int) (int, int) {
	split, ok := iSplit[i]
	if ok {
		return split.left, split.right
	}

	multiplier := 1
	for range numDigits(i) / 2 {
		multiplier *= 10
	}

	left := i / multiplier
	right := i - (left * multiplier)
	iSplit[i] = Split{left: left, right: right}
	return left, right
}

var StoneMap = map[string]int{}

func solve(i, t int) int {
	res, ok := StoneMap[fmt.Sprintf("%d,%d", i, t)]
	if ok {
		return res
	}

	if t == 0 {
		return 1
	} else if i == 0 {
		res = solve(1, t-1)
	} else if numDigits(i)%2 == 0 {
		left, right := SplitInt(i)
		res = solve(left, t-1) + solve(right, t-1)
	} else {
		res = solve(i*2024, t-1)
	}

	StoneMap[fmt.Sprintf("%d,%d", i, t)] = res
	return res
}

func Compute(stones []int, blinks int) int {
	total := 0

	for _, stone := range stones {
		total += solve(stone, blinks)
	}

	return total
}
