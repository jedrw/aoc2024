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

func Compute(rules [][]int, updates [][]int) int {
	total := 0
	for _, update := range updates {
		correct := true
		for _, rule := range rules {
			firstI := slices.IndexFunc(update, func(page int) bool { return page == rule[0] })
			secondI := slices.IndexFunc(update, func(page int) bool { return page == rule[1] })
			if firstI != -1 && secondI != -1 {
				if firstI > secondI {
					correct = false
				}
			}

		}

		if correct {
			middle := update[int(math.Ceil(float64(len(update)/2)))]
			total += middle
		}
	}
	return total
}
