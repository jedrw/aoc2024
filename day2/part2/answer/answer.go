package answer

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	increasing = iota
	decreasing = iota
)

func parse(file *os.File) [][]int {
	reports := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels := []int{}
		for _, level := range strings.Split(scanner.Text(), " ") {
			levelInt, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}

			levels = append(levels, levelInt)
		}

		reports = append(reports, levels)
	}

	return reports
}

func isSafeReport(report []int) bool {
	if checkReport(report) {
		return true
	}

	for i := range len(report) {
		newReport := make([]int, len(report))
		copy(newReport, report)
		newReport = append(newReport[:i], newReport[i+1:]...)
		if checkReport(newReport) {
			return true
		}
	}

	return false
}

func checkReport(report []int) bool {
	previous := report[0]
	var direction int
	if previous > report[1] {
		direction = decreasing
	} else {
		direction = increasing
	}

	for _, current := range report[1:] {
		if previous == current {
			return false
		}

		if direction == increasing {
			if previous > current {
				return false
			}
		} else {
			if previous < current {
				return false
			}
		}

		difference := int(math.Abs(float64(current - previous)))
		if difference < 1 || difference > 3 {
			return false
		}

		previous = current
	}

	return true
}

func Compute(inputFile string) int {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	reports := parse(file)
	safeReports := 0
	for _, report := range reports {
		if isSafeReport(report) {
			safeReports++
		}
	}

	return safeReports
}
