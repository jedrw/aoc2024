package answer

import (
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Parse(filePath string) ([][]int, [][]int) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	lines, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	sections := strings.Split(strings.Trim(string(lines), "\n"), "\n\n")
	rulesInput := sections[0]
	updatesInput := sections[1]

	rules := [][]int{}
	for _, line := range strings.Split(rulesInput, "\n") {
		ruleStrings := strings.Split(line, "|")
		first, err := strconv.Atoi(ruleStrings[0])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(ruleStrings[1])
		if err != nil {
			panic(err)
		}

		rule := []int{first, second}
		rules = append(rules, rule)
	}

	updates := [][]int{}
	for _, line := range strings.Split(updatesInput, "\n") {
		updateStrings := strings.Split(line, ",")
		update := []int{}
		for _, pageString := range updateStrings {
			pageInt, err := strconv.Atoi(pageString)
			if err != nil {
				panic(err)
			}

			update = append(update, pageInt)
		}

		updates = append(updates, update)
	}

	return rules, updates
}

func insertInt(array []int, value int, index int) []int {
	return append(array[:index], append([]int{value}, array[index:]...)...)
}

func removeInt(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func moveInt(array []int, srcIndex int, dstIndex int) []int {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}

func alignUpdateToRule(update []int, rules [][]int) []int {
	for _, rule := range rules {
		firstI := slices.IndexFunc(update, func(page int) bool { return page == rule[0] })
		secondI := slices.IndexFunc(update, func(page int) bool { return page == rule[1] })

		if firstI != -1 && secondI != -1 {
			if firstI > secondI {
				update = moveInt(update, firstI, 0)
			}
		}

	}

	return update
}

func isCorrect(update []int, rules [][]int) bool {
	for _, rule := range rules {
		firstI := slices.IndexFunc(update, func(page int) bool { return page == rule[0] })
		secondI := slices.IndexFunc(update, func(page int) bool { return page == rule[1] })
		if firstI != -1 && secondI != -1 {
			if firstI > secondI {
				return false
			}
		}
	}

	return true
}

func Compute(rules [][]int, updates [][]int) int {
	total := 0
	incorrectUpdates := [][]int{}
	for _, update := range updates {
		if !isCorrect(update, rules) {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	for _, update := range incorrectUpdates {
		update := alignUpdateToRule(update, rules)
		for !isCorrect(update, rules) {
			update = alignUpdateToRule(update, rules)
		}

		middle := update[int(math.Ceil(float64(len(update)/2)))]
		total += middle
	}

	return total
}
